package protection

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"

	"github.com/google/uuid"
	"golang.org/x/time/rate"

	"pow/internal/api/tcp"
)

type verifier interface {
	Verify(solution []byte, difficulty int64) (bool, error)
}

type Protector struct {
	limiter  *rate.Limiter
	verifier verifier

	defaultDifficulty int64
	currentDifficulty int64
	maxDifficulty     int64
}

func NewProtector(cfg Config, v verifier) *Protector {
	return &Protector{
		limiter:           rate.NewLimiter(rate.Limit(cfg.Limit), 0),
		verifier:          v,
		defaultDifficulty: cfg.DefaultDifficulty,
		currentDifficulty: cfg.DefaultDifficulty,
		maxDifficulty:     cfg.MaxDifficulty,
	}
}

func (p *Protector) getToken() []byte {
	return []byte(uuid.New().String())
}

func (p *Protector) incDiff() {
	atomic.AddInt64(&p.currentDifficulty, 1)
}

func (p *Protector) setDefaultDiff() {
	atomic.StoreInt64(&p.currentDifficulty, p.defaultDifficulty)
}

func (p *Protector) limit() {
	if p.limiter.Allow() {
		atomic.StoreInt64(&p.currentDifficulty, p.defaultDifficulty)
		return
	}

	if atomic.LoadInt64(&p.currentDifficulty) >= p.maxDifficulty {
		return
	}

	atomic.AddInt64(&p.currentDifficulty, 1)
}

func (p *Protector) Verify(solution []byte, difficulty int64) (bool, error) {
	p.limit()

	result, err := p.verifier.Verify(solution, difficulty)
	if err != nil {
		return false, fmt.Errorf("verify: %w", err)
	}
	return result, nil
}

// Challenge generates new challenge
func (p *Protector) Challenge(ctx context.Context) (Challenge, error) {
	return Challenge{
		Data:       []byte(uuid.New().String()),
		Difficulty: atomic.LoadInt64(&p.currentDifficulty),
	}, nil
}

func (p *Protector) ProcessorGenerateChallenge(
	ctx context.Context,
	_ []byte,
) ([]byte, error) {
	c, err := p.Challenge(ctx)
	if err != nil {
		return nil, fmt.Errorf("challenge: %w", err)
	}
	apiC := tcp.APIChallengeResponseV1{
		Data:       c.Data,
		Difficulty: c.Difficulty,
	}
	b, err := json.Marshal(apiC)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	return b, nil
}
