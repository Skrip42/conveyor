package module

import (
	"context"

	bufferfactory "github.com/Skrip42/conveyor/internal/buffer_factory"
	"github.com/Skrip42/conveyor/internal/item"
	"github.com/Skrip42/conveyor/internal/worker"
)

type batchModule[I, O any] struct {
	base          worker.Worker[I]
	bufferfactory bufferfactory.BufferFactory[I]
	deligate      Deligate[[]I, []O]
}

func (b *batchModule[I, O]) Run(ctx context.Context) <-chan item.Item[O] {
	output := make(chan item.Item[O])
	runCtx, cancel := context.WithCancel(ctx)
	input := b.base.Run(runCtx)

	push := func(pushedItem item.Item[O]) {
		select {
		case output <- pushedItem:
		case <-ctx.Done():
		}
	}

	inputErrorCh := make(chan error)

	bufferInput := make(chan I)
	bufferOutput, bufferFlush := b.bufferfactory(runCtx, bufferInput)

	go func() {
		defer close(bufferInput)
		for {
			select {
			case inputItem, ok := <-input:
				if !ok {
					bufferFlush(runCtx)
					return
				}
				if inputItem.Err != nil {
					bufferFlush(runCtx)
					select {
					case inputErrorCh <- inputItem.Err:
					case <-runCtx.Done():
					}
					return
				}
				select {
				case bufferInput <- inputItem.Payload:
				case <-runCtx.Done():
					return
				}
			case <-runCtx.Done():
			}
		}
	}()

	go func() {
		defer close(output)
		for {
			select {
			case bufferOutputs, ok := <-bufferOutput:
				if !ok {
					return
				}

				results, _, err := b.deligate.Eval(runCtx, bufferOutputs)
				if err != nil {
					push(item.Err[O](err))
					cancel()
					return
				}
				for _, result := range results {
					push(item.Create(result))
				}

			case err := <-inputErrorCh:
				push(item.Err[O](err))
				return
			case <-runCtx.Done():
				return
			}
		}
	}()

	return output
}

func NewBatchModule[I, O any](
	baseWorker worker.Worker[I],
	bufferfactory bufferfactory.BufferFactory[I],
	deligate Deligate[[]I, []O],
) worker.Worker[O] {
	return &batchModule[I, O]{
		base:          baseWorker,
		bufferfactory: bufferfactory,
		deligate:      deligate,
	}
}
