package tcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"

	"pow/internal/pow/hashchash"
)

type Client struct {
	c       *logConn
	h       *hashchash.Pow
	timeout int64
}

func (c *Client) Close() error {
	return c.c.r.Close()
}

func NewClient(cfg Config, h *hashchash.Pow) (*Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("resolve addr: %w", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}
	// err = conn.SetKeepAlive(true)
	// if err != nil {
	// 	return nil, fmt.Errorf("keep alive: %w", err)
	// }
	// err = conn.SetDeadline(time.Now().Add(time.Second * time.Duration(cfg.Timeout)))
	// if err != nil {
	// 	return nil, fmt.Errorf("set deadline: %w", err)
	// }
	return &Client{
		c: &logConn{r: conn},
		h: h,
	}, nil
}

func (c *Client) GetWisdom(ctx context.Context) (string, error) {
	ch, d, err := c.getChallange(ctx)
	if err != nil {
		return "", fmt.Errorf("get chalange: %w", err)
	}
	slog.Debug("got chalange", slog.String("chalange", string(ch)))

	solution, err := c.h.Solve(ctx, ch, d)
	if err != nil {
		return "", fmt.Errorf("solve: %w", err)
	}
	slog.Debug("solve chalange", slog.String("solution", string(solution)))
	wisdom, err := c.getWisdom(ctx, solution)
	if err != nil {
		return "", fmt.Errorf("get wisdom: %w", err)
	}
	return wisdom, nil
}

func (c *Client) getChallange(ctx context.Context) ([]byte, int64, error) {
	m := MessageV1{
		Type: MessageTypeGetChallengeRequest,
		Data: data{},
	}

	err := write2conn(c.c, m)
	if err != nil {
		return nil, 0, fmt.Errorf("write resp: %w", err)
	}

	err = m.Decode(c.c)
	if err != nil {
		return nil, 0, fmt.Errorf("decode: %w", err)
	}

	resp := APIChallengeResponseV1{}
	err = json.Unmarshal(m.Data.Payload, &resp)
	if err != nil {
		return nil, 0, fmt.Errorf("Unmarshal: %w", err)
	}
	return resp.Data, resp.Difficulty, nil
}

func (c *Client) getWisdom(ctx context.Context, solution []byte) (string, error) {
	m := MessageV1{
		Type: MessageTypeGetWisdomRequest,
		Data: data{
			Headers: map[string]string{
				solutionHeader: string(solution),
			},
		},
	}
	b, err := m.Encode()
	if err != nil {
		return "", fmt.Errorf("encode: %w", err)
	}
	err = write2conn(c.c, m)
	_, err = c.c.Write(b)
	if err != nil {
		return "", fmt.Errorf("write resp: %w", err)
	}

	err = m.Decode(c.c)
	if err != nil {
		return "", fmt.Errorf("decode: %w", err)
	}

	resp := APIWisdomResponseV1{}
	err = json.Unmarshal(m.Data.Payload, &resp)
	if err != nil {
		return "", fmt.Errorf("Unmarshal: %w", err)
	}
	return resp.Text, nil
}
