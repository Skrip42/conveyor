package storage

import (
	"context"

	"github.com/Skrip42/conveyor/internal/item"
	"github.com/Skrip42/conveyor/internal/worker"
)

type StorageAdapter[V any] func(context.Context, V) error

type storage[V any] struct {
	base    worker.Worker[V]
	adapter StorageAdapter[V]
}

func (s *storage[V]) Run(ctx context.Context) <-chan item.Item[V] {
	output := make(chan item.Item[V])
	runCtx, cancel := context.WithCancel(ctx)
	input := s.base.Run(runCtx)

	push := func(pushedItem item.Item[V]) {
		select {
		case output <- pushedItem:
		case <-ctx.Done():
		}
	}

	go func() {
		defer close(output)
		for {
			select {
			case inputItem, ok := <-input:
				if !ok {
					return
				}
				if inputItem.Err != nil {
					push(inputItem)
					return
				}

				err := s.adapter(runCtx, inputItem.Payload)

				if err != nil {
					inputItem.Err = err
					push(inputItem)
					cancel()
					return
				}
				push(inputItem)
			case <-ctx.Done():
				return
			}
		}
	}()

	return output
}

func NewStorage[V any](base worker.Worker[V], adapter StorageAdapter[V]) *storage[V] {
	return &storage[V]{
		base:    base,
		adapter: adapter,
	}
}
