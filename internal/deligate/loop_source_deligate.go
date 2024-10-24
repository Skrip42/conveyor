package deligate

import (
	"context"

	"github.com/Skrip42/conveyor/deligate"
)

type LoopSourceAdapter[V any] func(context.Context) ([]V, error)

type LoopSourceDeligate[V any] struct {
	adapter LoopSourceAdapter[V]
}

func (s *LoopSourceDeligate[V]) Eval(
	ctx context.Context,
	push func(V),
) error {
	for {
		data, err := s.adapter(ctx)
		if err != nil {
			return err
		}
		if len(data) == 0 {
			return nil
		}
		for _, item := range data {
			push(item)
		}
	}
}

func NewLoopSourceDeligate[V any](
	adapter LoopSourceAdapter[V],
) deligate.SourceDeligate[V] {
	return &LoopSourceDeligate[V]{adapter: adapter}
}
