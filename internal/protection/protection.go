package protection

import (
	"fmt"
	"sync/atomic"

	"github.com/google/uuid"
	"golang.org/x/time/rate"
)

type Verifier interface {
	Verify(solution []byte, difficulty int64) (bool, error)
}

type Protector struct {
	limiter  *rate.Limiter
	verifier Verifier

	defaultDifficulty int64
	currentDifficulty int64
	maxDifficulty     int64
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
