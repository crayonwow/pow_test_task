package tcp

import (
	"context"
	"fmt"

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

type Controller struct {
	m map[MessageTypeV1]*Processor
}

func (r *Controller) get(m MessageV1) (*Processor, error) {
	a, ok := r.m[m.Type]
	if !ok {
		return nil, fmt.Errorf("processor: %w", errors.ErrNotFound)
	}

	// that's kinda middleware but really dumb
	if solution := m.Data.Headers[solutionHeader]; a.Protected && solution == "" {
		return nil, fmt.Errorf("solution is required: request: %w", localErr.ErrInvalid)
	}

	return a, nil
}

func NewController(actors []*Processor) *Controller {
	r := &Controller{
		m: make(map[MessageTypeV1]*Processor, len(actors)),
	}

	for _, a := range actors {
		r.m[a.TypeRequest] = a
	}

	return r
}
