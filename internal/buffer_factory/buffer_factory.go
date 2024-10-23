package bufferfactory

import (
	"context"

	"github.com/Skrip42/conveyor/internal/item"
)

type BufferFactory[V any] func(context.Context, chan V) (<-chan []V, func(context.Context))

func BufferFactoryAdapter[V any](context.Context, chan item.Item[V]) (<-chan []item.Item[V], func(context.Context)) {
}
