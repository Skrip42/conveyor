package conveyor

import (
	"github.com/Skrip42/conveyor/internal/module/processor"
	"github.com/Skrip42/conveyor/internal/module/source"
	"github.com/Skrip42/conveyor/internal/module/storage"
)

func NewSource[V any](adapter source.SourceAdapter[V]) Controller[V] {
	return &controller[V]{worker: source.NewSource(adapter)}
}

func NewStorage[V any](base Controller[V], adapter storage.StorageAdapter[V]) Controller[V] {
	return &controller[V]{worker: storage.NewStorage(base.extend(), adapter)}
}

func NewProcessor[I, O any](base Controller[I], adapter processor.ProcessorAdater[I, O]) Controller[O] {
	return &controller[O]{worker: processor.NewProcessor(base.extend(), adapter)}
}
