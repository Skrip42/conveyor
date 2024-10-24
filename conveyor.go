package conveyor

import (
	bufferfactory "github.com/Skrip42/conveyor/internal/buffer_factory"
	"github.com/Skrip42/conveyor/internal/module"
	"github.com/Skrip42/conveyor/internal/module/source"
)

func NewSource[V any](adapter source.SourceAdapter[V]) Controller[V] {
	return &controller[V]{base: source.NewSource(adapter)}
}

func NewStorage[V any](
	base Controller[V],
	adapter module.StorageAdapter[V],
) Controller[V] {
	return &controller[V]{
		base: module.NewModule(
			base.worker(),
			module.NewStorageDeligate(adapter),
		),
	}
}

func NewBatchStorage[V any](
	base Controller[V],
	adapter module.StorageAdapter[[]V],
	bufferfactory bufferfactory.BufferFactory[V],
) Controller[V] {
	return &controller[V]{
		base: module.NewBatchModule(
			base.worker(),
			bufferfactory,
			module.NewStorageDeligate(adapter),
		),
	}
}

func NewProcessor[I, O any](
	base Controller[I],
	adapter module.ProcessorAdapter[I, O],
) Controller[O] {
	return &controller[O]{
		base: module.NewModule(
			base.worker(),
			module.NewProcessorDeligate(adapter),
		),
	}
}

func NewBatchProcessor[I, O any](
	base Controller[I],
	adapter module.ProcessorAdapter[[]I, []O],
	bufferfactory bufferfactory.BufferFactory[I],
) Controller[O] {
	return &controller[O]{
		base: module.NewBatchModule(
			base.worker(),
			bufferfactory,
			module.NewProcessorDeligate(adapter),
		),
	}
}

func NewFilter[V any](
	base Controller[V],
	adapter module.FilterAdapter[V],
) Controller[V] {
	return &controller[V]{
		base: module.NewModule(
			base.worker(),
			module.NewFilterDeligate(adapter),
		),
	}
}
