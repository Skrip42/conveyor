package worker

import (
	"context"

	"github.com/Skrip42/conveyor/internal/item"
)

type Worker[V any] interface {
	Run(context.Context) <-chan item.Item[V]
}
