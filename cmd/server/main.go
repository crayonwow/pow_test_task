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
	"pow/internal/protection"
	"pow/internal/wisdom"
)

type Config struct {
	HashCash  hashchash.Confing `json:"hash_cash"`
	Protector protection.Config `json:"protector"`
	TCPServer tcp.Config        `json:"tcp_server"`
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
	cfgPath := os.Getenv("POW_SERVER_CONFIG_PATH")
	if cfgPath == "" {
		slog.Error("empty config path")
		os.Exit(1)
	}
	isDebug := os.Getenv("POW_SERVER_DEBUG")
	if isDebug == "1" {
		slog.SetLogLoggerLevel(slog.LevelDebug.Level())
	}
	slog.Debug(
		"found config file",
		slog.String("path", cfgPath),
	)

	cfg, err := newConfigFromFile(cfgPath)
	if err != nil {
		slog.Error("new config", slog.String("error", err.Error()))
		os.Exit(1)
	}
	slog.Debug(
		"loaded config file",
		"cfg", cfg,
	)
	h := hashchash.New(hashchash.NewInMemmoryStorage(), cfg.HashCash)
	ws := wisdom.NewService(wisdom.NewInMemmoryRepository())
	p := protection.NewProtector(cfg.Protector, h)
	processors := []*tcp.Processor{
		{
			TypeRequest:  tcp.MessageTypeGetChallengeRequest,
			TypeResponse: tcp.MessageTypeGetChallengeResponse,
			Processor:    p.ProcessorGenerateChallenge,
			Protected:    false,
		},
		{
			TypeRequest:  tcp.MessageTypeGetWisdomRequest,
			TypeResponse: tcp.MessageTypeGetWisdomResponse,
			Processor:    ws.ProcessorGet,
			Protected:    true,
		},
	}

	c := tcp.NewController(processors, p)
	s := tcp.NewServer(cfg.TCPServer, c)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	err = s.Start(ctx)
	if err != nil {
		slog.Error("start", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("gracefully stopped")
}
