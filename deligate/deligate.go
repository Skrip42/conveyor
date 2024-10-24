package deligate

import (
	"context"
)

type Deligate[I, O any] interface {
	Eval(context.Context, I) (O, bool, error)
}

type SourceDeligate[V any] interface {
	Eval(context.Context, func(V)) error
}
