package wisdom

import "context"

type repository interface {
	Get(context.Context) (string, error)
}

type Service struct {
	repo repository
}

func (s *Service) Get(ctx context.Context) (string, error) {
	return s.repo.Get(ctx)
}
