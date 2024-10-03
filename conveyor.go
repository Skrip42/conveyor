package conveyor

import "context"

type Controller[V any] interface {
	Run(context.Context) error
	extend() worker[V]
}

type controller[V any] struct {
	worker worker[V]
}

func (c *controller[V]) extend() worker[V] {
	return c.worker
}

func (c *controller[V]) Run(ctx context.Context) error {
	output := c.worker.Run(ctx)

	for item := range output {
		if item.err != nil {
			return item.err
		}
	}

	return ctx.Err()
}
