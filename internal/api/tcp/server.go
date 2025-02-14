package tcp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync"
	"time"

	localErr "pow/internal/errors"
)

func NewServer(cfg Config, r *Controller) *Server {
	return &Server{
		resolver: r,
		host:     cfg.Host,
		port:     cfg.Port,
	}
}

type Server struct {
	resolver   *Controller
	host, port string
	wg         sync.WaitGroup
	timeout    int64
}

func write2conn(w io.Writer, m MessageV1) error {
	b, err := m.Encode()
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	_, err = w.Write(b)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func (s *Server) handleRequest(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	defer s.wg.Done()
	m := MessageV1{}
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	c := &logConn{r: conn}

	for {
		m.clear()

		err := m.Decode(c)
		if err != nil {
			slog.Error(
				"decode",
				slog.String("error", err.Error()),
			)
			return
		}

		slog.Debug("getting resolver")
		a, err := s.resolver.get(m)
		if err != nil {
			if errors.Is(err, localErr.ErrNotFound) {
				slog.Debug("controller not found")
				return
			}

			slog.Error(
				"resolver get",
				slog.Uint64("type", uint64(m.Type)),
				slog.String("error", err.Error()),
			)
			return
		}

		res, err := a.Processor(ctx, m.Data.Payload)
		if err != nil {
			slog.Error(
				"run processor",
				slog.Uint64("type", uint64(m.Type)),
				slog.String("payload", string(m.Data.Payload)),
				slog.String("error", err.Error()),
			)
			continue
		}
		if a.TypeResponse == MessageTypeEmpty {
			slog.Debug("no response")
			continue
		}
		m = MessageV1{
			Type: a.TypeResponse,
			Data: data{
				Headers: nil,
				Payload: res,
			},
		}
		err = write2conn(c, m)
		if err != nil {
			slog.Error(
				"write response",
				slog.Uint64("type", uint64(m.Type)),
				slog.String("payload", string(m.Data.Payload)),
				slog.String("res", string(res)),
				slog.String("error", err.Error()),
			)
			continue
		}

		slog.Debug("good")
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (s *Server) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	go func() {
		select {
		case <-ctx.Done():
			cerr := listener.Close()
			if cerr != nil {
				slog.Error("close", slog.String("err", err.Error()))
			}
			return
		}
	}()

	slog.Info(
		"server started",
		"addr", fmt.Sprintf("%s:%s", s.host, s.port),
	)

LOOP:
	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				break
			}
			return fmt.Errorf("accept: %w", err)
		}
		slog.Info("got connection")

		s.wg.Add(1)
		go s.handleRequest(ctx, conn)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	slog.Debug("waiting")
	s.wg.Wait()
	return nil
}

type logConn struct {
	r net.Conn
}

func (lo *logConn) Write(p []byte) (n int, err error) {
	n, err = lo.r.Write(p)
	attrs := []any{
		slog.String("data", string(p)),
		slog.Int("n", n),
	}
	if err != nil {
		attrs = append(attrs, slog.String("err", err.Error()))
	}

	slog.Debug("write", attrs...)
	return n, err
}

func (lo *logConn) Read(p []byte) (n int, err error) {
	n, err = lo.r.Read(p)
	attrs := []any{
		slog.String("data", string(p)),
		slog.Int("n", n),
	}
	if err != nil {
		attrs = append(attrs, slog.String("err", err.Error()))
	}
	slog.Debug("read", attrs...)
	return n, err
}
