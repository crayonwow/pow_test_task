package tcp

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"

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

func handleRequest(ctx context.Context, conn net.Conn, _resolver *Controller) {
	defer conn.Close()
	m := MessageV1{}
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)

	for {
		m.clear()
		err := m.Decode(r)
		if err != nil {
			slog.Error(
				"decode",
				slog.String("error", err.Error()),
			)
			return
		}

		a, err := _resolver.get(m)
		if err != nil {
			if errors.Is(err, localErr.ErrNotFound) {
				continue
			}

			slog.Error(
				"resolver get",
				slog.Uint64("type", uint64(m.Type)),
				slog.String("error", err.Error()),
			)
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
			continue
		}
		m = MessageV1{
			Type: a.TypeResponse,
			Data: data{
				Headers: nil,
				Payload: res,
			},
		}
		err = write2conn(w, m)
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
		w.Flush()

		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}

func (s *Server) Start(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("accept: %w", err)
		}

		go handleRequest(ctx, conn, s.resolver)
		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
}

type logConn struct {
	r net.Conn
}

func (lo *logConn) Write(p []byte) (n int, err error) {
	n, err = lo.Write(p)
	slog.Debug("read",
		slog.String("data", string(p)),
		slog.String("err", err.Error()),
		slog.Int("n", n),
	)
	return
}

func (lo *logConn) Read(p []byte) (n int, err error) {
	n, err = lo.r.Read(p)
	slog.Debug("read",
		slog.String("data", string(p)),
		slog.String("err", err.Error()),
		slog.Int("n", n),
	)
	return
}
