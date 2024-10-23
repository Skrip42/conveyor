package module

import (
	"context"

	"github.com/Skrip42/conveyor/internal/item"
	"github.com/Skrip42/conveyor/internal/worker"
)

type module[I, O any] struct {
	base     worker.Worker[I]
	deligate Deligate[I, O]
}

func (b *module[I, O]) Run(ctx context.Context) <-chan item.Item[O] {
	output := make(chan item.Item[O])
	runCtx, cancel := context.WithCancel(ctx)
	input := b.base.Run(runCtx)

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
					push(item.Err[O](inputItem.Err))
					return
				}

				outputItem, isActual, err := b.deligate.Eval(runCtx, inputItem.Payload)
				if err != nil {
					push(item.Err[O](err))
					cancel()
					return
				}
				if isActual {
					push(item.Create(outputItem))
				}

			case <-runCtx.Done():
				return
			}
		}
	}()

	return output
}

func NewModule[I, O any](baseWorker worker.Worker[I], deligate Deligate[I, O]) worker.Worker[O] {
	return &module[I, O]{
		base:     baseWorker,
		deligate: deligate,
	}
}
