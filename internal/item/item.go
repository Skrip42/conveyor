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
