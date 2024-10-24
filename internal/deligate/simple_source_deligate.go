package deligate

import (
	"context"

	"github.com/Skrip42/conveyor/deligate"
)

type SimpleSourceAdapter[V any] func(context.Context) ([]V, error)

type SimpleSourceDeligate[V any] struct {
	adapter SimpleSourceAdapter[V]
}

func (s *SimpleSourceDeligate[V]) Eval(
	ctx context.Context,
	push func(V),
) error {
	data, err := s.adapter(ctx)
	if err != nil {
		return nil
	}
	for _, item := range data {
		push(item)
	}
	return nil
}

func NewSimplaSourceDeligate[V any](
	adapter SimpleSourceAdapter[V],
) deligate.SourceDeligate[V] {
	return &SimpleSourceDeligate[V]{adapter: adapter}
}
