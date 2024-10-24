package deligate

import (
	"context"

	"github.com/Skrip42/conveyor/deligate"
	"golang.org/x/exp/constraints"
)

type ByVersionSourceAdapter[V, C any] func(context.Context, C) ([]V, error)
type VersionIniter[C constraints.Ordered] func(context.Context) (C, error)
type ExstractVersion[V, C any] func(V) C

type ByVersionSourceDeligate[V any, C constraints.Ordered] struct {
	adapter    ByVersionSourceAdapter[V, C]
	initer     VersionIniter[C]
	exstractor ExstractVersion[V, C]
}

func (s *ByVersionSourceDeligate[V, C]) Eval(
	ctx context.Context,
	push func(V),
) error {
	version, err := s.initer(ctx)
	if err != nil {
		return err
	}
	for {
		data, err := s.adapter(ctx, version)
		if err != nil {
			return err
		}
		if len(data) == 0 {
			return nil
		}
		for _, item := range data {
			push(item)
			if newVersion := s.exstractor(item); newVersion > version {
				version = newVersion
			}
		}
	}
}

func NewByVersionSourceDeligate[V any, C constraints.Ordered](
	adapter ByVersionSourceAdapter[V, C],
	initer VersionIniter[C],
	exstractor ExstractVersion[V, C],
) deligate.SourceDeligate[V] {
	return &ByVersionSourceDeligate[V, C]{
		adapter:    adapter,
		initer:     initer,
		exstractor: exstractor,
	}
}
