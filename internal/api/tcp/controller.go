package tcp

import (
	"context"
	"fmt"
	"log/slog"

	"pow/internal/errors"
	localErr "pow/internal/errors"
)

type ProcessorFunc func(context.Context, []byte) ([]byte, error)

type Processor struct {
	// Protected means that is required to solve challenge
	Protected    bool
	TypeRequest  MessageTypeV1
	TypeResponse MessageTypeV1
	Processor    ProcessorFunc
}

type verifier interface {
	Verify(solution []byte) (bool, error)
}
type Controller struct {
	m map[MessageTypeV1]*Processor
	v verifier
}

func (r *Controller) get(m MessageV1) (*Processor, error) {
	slog.Debug("getting resolver", "message", m)

	a, ok := r.m[m.Type]
	if !ok {
		return nil, fmt.Errorf("processor: %w", errors.ErrNotFound)
	}

	// that's kinda middleware but really dumb
	if a.Protected {
		solution := m.Data.Headers[solutionHeader]
		good, err := r.v.Verify([]byte(solution))
		if err != nil || !good {
			slog.Debug("no solution")
			return nil, fmt.Errorf("solution is required: request: %w", localErr.ErrInvalid)
		}
	}

	return a, nil
}

func NewController(actors []*Processor, v verifier) *Controller {
	r := &Controller{
		m: make(map[MessageTypeV1]*Processor, len(actors)),
		v: v,
	}

	for _, a := range actors {
		r.m[a.TypeRequest] = a
	}

	return r
}
