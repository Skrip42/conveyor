package conveyor

import (
	"context"
	"io"
)

type SourceAdapter[V any] func(context.Context, func(V)) error

type source[V any] struct {
	items SourceAdapter[V]
}

func (s source[V]) Run(ctx context.Context) <-chan item[V] {
	output := make(chan item[V])

	push := func(payload V) {
		select {
		case output <- item[V]{payload: payload}:
		case <-ctx.Done():
		}
	}

	go func() {
		defer close(output)
		err := s.items(ctx, push)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case output <- item[V]{err: err}:
			case <-ctx.Done():
			}
		}
	}()

	return output
}

func NewSource[V any](items SourceAdapter[V]) Controller[V] {
	worker := source[V]{items: items}

	return &controller[V]{worker: worker}
}
