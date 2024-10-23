package module

import (
	"context"
)

type Deligate[I, O any] interface {
	Eval(context.Context, I) (O, bool, error)
}

//
// type BatchDeligate[I, O any] interface {
// 	Eval(context.Context, []item.Item[I]) []item.Item[O]
// }
