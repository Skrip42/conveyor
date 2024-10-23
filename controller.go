package conveyor

import (
	"context"

	"github.com/Skrip42/conveyor/internal/worker"
)

type Controller[V any] interface {
	Run(context.Context) error
	worker() worker.Worker[V]
}

type controller[V any] struct {
	base worker.Worker[V]
}

func (c *controller[V]) worker() worker.Worker[V] {
	return c.base
}

func (c *controller[V]) Run(ctx context.Context) error {
	output := c.base.Run(ctx)

	for item := range output {
		if item.Err != nil {
			return item.Err
		}
	}

	return ctx.Err()
}
