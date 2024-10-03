package conveyor

import "context"

type worker[V any] interface {
	Run(context.Context) <-chan item[V]
}
