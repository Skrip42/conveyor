package conveyor

import (
	"context"

	"github.com/Skrip42/conveyor/internal/worker"
)

type Controller[V any] interface {
	Run(context.Context) error
	extend() worker.Worker[V]
}

type controller[V any] struct {
	worker worker.Worker[V]
}

func (c *controller[V]) extend() worker.Worker[V] {
	return c.worker
}

func (c *controller[V]) Run(ctx context.Context) error {
	output := c.worker.Run(ctx)

	for item := range output {
		if item.Err != nil {
			return item.Err
		}
	}

	return ctx.Err()
}
