package processor

import (
	"context"

	"github.com/Skrip42/conveyor/internal/item"
	"github.com/Skrip42/conveyor/internal/worker"
)

type ProcessorAdater[I, O any] func(context.Context, I) (O, error)

type processor[I, O any] struct {
	base    worker.Worker[I]
	adapter ProcessorAdater[I, O]
}

func (s *processor[I, O]) Run(ctx context.Context) <-chan item.Item[O] {
	output := make(chan item.Item[O])
	runCtx, cancel := context.WithCancel(ctx)
	input := s.base.Run(runCtx)

	push := func(pushedItem item.Item[O]) {
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
					push(item.Convert[I, O](inputItem))
					return
				}
				outputItem := item.Convert[I, O](inputItem)

				outputItem.Payload, outputItem.Err = s.adapter(runCtx, inputItem.Payload)

				push(outputItem)
				if outputItem.Err != nil {
					cancel()
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return output
}

func NewProcessor[I, O any](base worker.Worker[I], adapter ProcessorAdater[I, O]) *processor[I, O] {
	return &processor[I, O]{
		base:    base,
		adapter: adapter,
	}
}
