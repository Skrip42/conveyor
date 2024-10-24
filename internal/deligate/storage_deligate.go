package deligate

import (
	"context"

	"github.com/Skrip42/conveyor/deligate"
)

type StorageAdapter[V any] func(context.Context, V) error

type StorageDeligate[V any] struct {
	adapter StorageAdapter[V]
}

func (s *StorageDeligate[V]) Eval(
	ctx context.Context,
	data V,
) (V, bool, error) {
	err := s.adapter(ctx, data)
	return data, true, err
}

func NewStorageDeligate[V any](adapter StorageAdapter[V]) deligate.Deligate[V, V] {
	return &StorageDeligate[V]{adapter: adapter}
}
