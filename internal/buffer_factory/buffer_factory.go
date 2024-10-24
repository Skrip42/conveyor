package bufferfactory

import (
	"context"
)

type BufferFactory[V any] func(
	context.Context,
	chan V,
) (
	<-chan []V,
	func(context.Context),
)
