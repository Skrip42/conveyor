package deligate

import (
	"context"

	"github.com/Skrip42/conveyor/deligate"
)

type FilterAdapter[V any] func(context.Context, V) (bool, error)

type FilterDeligate[V any] struct {
	adapter FilterAdapter[V]
}

func (s *FilterDeligate[V]) Eval(
	ctx context.Context,
	data V,
) (V, bool, error) {
	isFiltered, err := s.adapter(ctx, data)
	return data, isFiltered, err
}

func NewFilterDeligate[V any](adapter FilterAdapter[V]) deligate.Deligate[V, V] {
	return &FilterDeligate[V]{adapter: adapter}
}
