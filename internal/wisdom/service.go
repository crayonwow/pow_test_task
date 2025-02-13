package wisdom

import (
	"context"
	"encoding/json"
	"fmt"

	"pow/internal/api/tcp"
)

func NewService(r repository) *Service {
	return &Service{repo: r}
}

type repository interface {
	Get(context.Context) (string, error)
}

type Service struct {
	repo repository
}

func (s *Service) Get(ctx context.Context) (string, error) {
	return s.repo.Get(ctx)
}

func (s *Service) ProcessorGet(ctx context.Context, _ []byte) ([]byte, error) {
	_wisdom, err := s.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}
	scheme := tcp.APIWisdomResponseV1{
		Text: _wisdom,
	}

	b, err := json.Marshal(scheme)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	return b, nil
}
