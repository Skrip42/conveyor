package conveyor

import (
	bufferfactory "github.com/Skrip42/conveyor/buffer_factory"
	"github.com/Skrip42/conveyor/internal/deligate"
	"golang.org/x/exp/constraints"
)

func NewStorage[V any](
	base Controller[V],
	adapter deligate.StorageAdapter[V],
) Controller[V] {
	return NewModule(
		base,
		deligate.NewStorageDeligate(adapter),
	)
}

func NewBatchStorage[V any](
	base Controller[V],
	adapter deligate.StorageAdapter[[]V],
	bufferfactory bufferfactory.BufferFactory[V],
) Controller[V] {
	return NewBatchModule(
		base,
		deligate.NewStorageDeligate(adapter),
		bufferfactory,
	)
}

func NewProcessor[I, O any](
	base Controller[I],
	adapter deligate.ProcessorAdapter[I, O],
) Controller[O] {
	return NewModule(
		base,
		deligate.NewProcessorDeligate(adapter),
	)
}

func NewBatchProcessor[I, O any](
	base Controller[I],
	adapter deligate.ProcessorAdapter[[]I, []O],
	bufferfactory bufferfactory.BufferFactory[I],
) Controller[O] {
	return NewBatchModule(
		base,
		deligate.NewProcessorDeligate(adapter),
		bufferfactory,
	)
}

func NewFilter[V any](
	base Controller[V],
	adapter deligate.FilterAdapter[V],
) Controller[V] {
	return NewModule(
		base,
		deligate.NewFilterDeligate(adapter),
	)
}

func NewSimpleSource[V any](
	adapter deligate.SimpleSourceAdapter[V],
) Controller[V] {
	return NewSource(
		deligate.NewSimplaSourceDeligate(adapter),
	)
}

func NewLoopSource[V any](
	adapter deligate.LoopSourceAdapter[V],
) Controller[V] {
	return NewSource(
		deligate.NewLoopSourceDeligate(adapter),
	)
}

func NewByVersionSource[V any, C constraints.Ordered](
	adapter deligate.ByVersionSourceAdapter[V, C],
	initer deligate.VersionIniter[C],
	exstractor deligate.ExstractVersion[V, C],
) Controller[V] {
	return NewSource(
		deligate.NewByVersionSourceDeligate(adapter, initer, exstractor),
	)
}
