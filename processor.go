package conveyor

import "context"

type ProcessorAdater[I, O any] func(context.Context, I) (O, error)

type processor[I, O any] struct {
	worker  worker[I]
	process ProcessorAdater[I, O]
}

func (s processor[I, O]) Run(ctx context.Context) <-chan item[O] {
	output := make(chan item[O])

	runCtx, cancel := context.WithCancel(ctx)

	input := s.worker.Run(runCtx)

	go func() {
		defer close(output)
		for {
			select {
			case inputItem, ok := <-input:
				if !ok {
					return
				}
				if inputItem.err != nil {
					select {
					case output <- cloneItemMeta[I, O](inputItem):
						cancel()
						return
					case <-ctx.Done():
						return
					}
				}
				outputItem := cloneItemMeta[I, O](inputItem)

				outputItem.payload, outputItem.err = s.process(runCtx, inputItem.payload)
				select {
				case output <- outputItem:
					if outputItem.err != nil {
						cancel()
						return
					}
				case <-ctx.Done():
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	return output
}

func NewProcessor[I, O any](base Controller[I], process ProcessorAdater[I, O]) Controller[O] {
	worker := processor[I, O]{
		worker:  base.extend(),
		process: process,
	}

	return &controller[O]{worker: worker}
}
