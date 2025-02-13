package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"pow/internal/api/tcp"
	"pow/internal/pow/hashchash"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug.Level())
}

type Config struct {
	TCPServerHost tcp.Config `json:"tcp_server"`
}

func newConfigFromFile(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("open: %w", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return Config{}, fmt.Errorf("read all: %w", err)
	}
	cfg := Config{}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshal: %w", err)
	}
	return cfg, nil
}

func main() {
	cfgPath := os.Getenv("POW_CLIENT_CONFIG_PATH")
	if cfgPath == "" {
		slog.Error("empty config path")
		os.Exit(1)
	}

	cfg, err := newConfigFromFile(cfgPath)
	if err != nil {
		slog.Error("new config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	h := hashchash.New(hashchash.NewMockStorage(), hashchash.Confing{})
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	c, err := tcp.NewClient(cfg.TCPServerHost, h)
	if err != nil {
		slog.Error("new client", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer c.Close()
	wisdom, err := c.GetWisdom(ctx)
	if err != nil {
		slog.Error("get wisdom", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Error("success", slog.String("wisdom", wisdom))
}
