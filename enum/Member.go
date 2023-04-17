package enum

import (
	"fmt"
)

type (
	Member[T fmt.Stringer, A any] struct {
		Enum[T, A]
		v     T
		name  string
		place int
	}
)

func (x Member[_, _]) String() string {
	return x.name
}

func (x Member[_, _]) Name() string {
	return x.name
}

func (x Member[_, _]) Is(other Member[_, _]) bool {
	return x.name == other.name
}
