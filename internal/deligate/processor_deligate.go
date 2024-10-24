package deligate

import (
	"context"

	"github.com/Skrip42/conveyor/deligate"
)

type ProcessorAdapter[I, O any] func(context.Context, I) (O, error)

type ProcessorDeligate[I, O any] struct {
	adapter ProcessorAdapter[I, O]
}

func (s *ProcessorDeligate[I, O]) Eval(
	ctx context.Context,
	data I,
) (O, bool, error) {
	result, err := s.adapter(ctx, data)
	return result, true, err
}

func NewProcessorDeligate[I, O any](adapter ProcessorAdapter[I, O]) deligate.Deligate[I, O] {
	return &ProcessorDeligate[I, O]{adapter: adapter}
}
