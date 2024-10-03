package item

type Item[V any] struct {
	Payload V
	Err     error
}

func Create[V any](payload V) Item[V] {
	return Item[V]{
		Payload: payload,
	}
}

func Err[V any](err error) Item[V] {
	return Item[V]{
		Err: err,
	}
}

func Convert[I, O any](source Item[I]) Item[O] {
	return Item[O]{
		Err: source.Err,
	}
}

func ConvertWithPayload[I, O any](source Item[I], payload O) Item[O] {
	return Item[O]{
		Err:     source.Err,
		Payload: payload,
	}
}
