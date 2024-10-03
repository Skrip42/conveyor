package source

import (
	"context"
	"io"

	"github.com/Skrip42/conveyor/internal/item"
)

type SourceAdapter[V any] func(context.Context, func(V)) error

type source[V any] struct {
	adapter SourceAdapter[V]
}

func (s *source[V]) Run(ctx context.Context) <-chan item.Item[V] {
	output := make(chan item.Item[V])

	push := func(payload V) {
		select {
		case output <- item.Create(payload):
		case <-ctx.Done():
		}
	}

	go func() {
		defer close(output)
		err := s.adapter(ctx, push)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case output <- item.Err[V](err):
			case <-ctx.Done():
			}
		}
	}()

	return output
}

func NewSource[V any](adapter SourceAdapter[V]) *source[V] {
	return &source[V]{adapter: adapter}
}
