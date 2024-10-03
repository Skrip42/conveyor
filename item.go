package conveyor

type item[V any] struct {
	payload V
	err     error
}

func cloneItemMeta[I, O any](sourceItem item[I]) item[O] {
	return item[O]{
		err: sourceItem.err,
	}
}
