package module

import (
	"context"
)

type Deligate[I, O any] interface {
	Eval(context.Context, I) (O, bool, error)
}
