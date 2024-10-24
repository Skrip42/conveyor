package conveyor

import (
	bufferfactory "github.com/Skrip42/conveyor/buffer_factory"
	"github.com/Skrip42/conveyor/deligate"
	"github.com/Skrip42/conveyor/internal/module"
)

func NewSource[V any](
	deligate deligate.SourceDeligate[V],
) Controller[V] {
	return &controller[V]{
		base: module.NewSourceModule(deligate),
	}
}

func NewModule[I, O any](
	base Controller[I],
	deligate deligate.Deligate[I, O],
) Controller[O] {
	return &controller[O]{
		base: module.NewModule(
			base.worker(),
			deligate,
		),
	}
}

func NewBatchModule[I, O any](
	base Controller[I],
	deligate deligate.Deligate[[]I, []O],
	bufferfactory bufferfactory.BufferFactory[I],
) Controller[O] {
	return &controller[O]{
		base: module.NewBatchModule(
			base.worker(),
			deligate,
			bufferfactory,
		),
	}
}
