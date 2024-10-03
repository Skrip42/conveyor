package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Skrip42/conveyor"
)

func main() {
	source := Source{
		items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	cnt := conveyor.NewSource(source.Produce)
	cnt2 := conveyor.NewProcessor(cnt, Processor)
	cnt2 = conveyor.NewStorage(cnt2, Storage)

	cnt2.Run(context.Background())
}

type Source struct {
	items []int
}

func (s Source) Produce(ctx context.Context, push func(int)) error {
	for _, item := range s.items {
		push(item)
	}
	return nil
}

func Processor(_ context.Context, item int) (string, error) {
	return "item_is_" + strconv.Itoa(item), nil
}

func Storage(_ context.Context, item string) error {
	fmt.Println(item)
	return nil
}
