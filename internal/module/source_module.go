package module

import (
	"context"
	"io"

	"github.com/Skrip42/conveyor/deligate"
	"github.com/Skrip42/conveyor/internal/item"
)

type source[V any] struct {
	deligate deligate.SourceDeligate[V]
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
		err := s.deligate.Eval(ctx, push)
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

func NewSourceModule[V any](deligate deligate.SourceDeligate[V]) *source[V] {
	return &source[V]{deligate: deligate}
}
