package hashchash

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/umahmood/hashcash"

	"pow/internal/pow"
)

// var defaultPow = New(newInMemmoryStorage(), Confing{
// 	Expiration: 3600,
// })

func New(storage Storage, config Confing) *Pow {
	return &Pow{
		storage:    storage,
		expiration: config.Expiration,
		difficulty: config.Difficulty,
	}
}

// minimize use interfaces from dependencies
// it might be painful later
type Storage interface {
	Add(string) error
	Spent(string) bool
}

type Pow struct {
	storage Storage

	// expiration is read only
	expiration int64

	difficulty int64
}

// we consider that all data is valid
func validator(_ string) bool {
	return true
}

func (p *Pow) Solve(ctx context.Context, challenge []byte) ([]byte, error) {
	hc, err := p.new(challenge)
	if err != nil {
		return nil, fmt.Errorf("new: %w", err)
	}

	var solution string
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		solution, err = hc.Compute()
		if err != nil {
			if errors.Is(err, hashcash.ErrSolutionFail) {
				continue
			}
			return nil, fmt.Errorf("failed to compute hashcash: %w", err)
		}
		break
	}

	return []byte(solution), nil
}

func (p *Pow) Verify(solution []byte) (bool, error) {
	hc, err := p.new(solution)
	if err != nil {
		return false, fmt.Errorf("new: %w", err)
	}
	valid, err := hc.Verify(string(solution))
	if err != nil {
		return false, pow.ErrUnexpected
	}

	return valid, nil
}

func (p *Pow) IncDifficulty() {
	atomic.AddInt64(&p.difficulty, 1)
}

func (p *Pow) SetDifficulty(difficulty int64) {
	atomic.StoreInt64(&p.difficulty, difficulty)
}

func (p *Pow) new(data []byte) (*hashcash.Hashcash, error) {
	r := &hashcash.Resource{
		Data:          string(data),
		ValidatorFunc: validator,
	}
	c := &hashcash.Config{
		Bits:    int(atomic.LoadInt64(&p.difficulty)),
		Expired: time.Now().Add(time.Second * time.Duration(p.expiration)),
		Future:  time.Now().Add(time.Hour * 48),
		Storage: p.storage,
	}
	hc, err := hashcash.New(r, c)
	if err != nil {
		if errors.Is(err, hashcash.ErrResourceFail) {
			return nil, pow.ErrInvalid
		}
		return nil, fmt.Errorf("failed to create hashcash: %w", err)

	}
	return hc, nil
}
