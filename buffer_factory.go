package conveyor

import "context"

type bufferFactory[V any] func(context.Context, chan V) (<-chan []V, func(context.Context))
