package protection

import (
	"context"

	"golang.org/x/time/rate"
	"github.com/google/uuid"
)

type Verifier interface {
	Solve(ctx context.Context, challenge []byte) ([]byte, error)
}

type Protector struct {
	limiter  *rate.Limiter
	verifier Verifier

	defaultDifficulty int64
	currentDifficulty int64
}

func (p *Protector) getToken() ([]byte, error) {
  uuid.New().String()
}


