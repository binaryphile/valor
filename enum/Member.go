package enum

type (
	Member[T any] struct {
		Enum[T]
		v string
	}
)

func (x Member[T]) String() string {
	return x.v
}

func (x Member[T]) Name() string {
	return x.v
}

func (x Member[T]) Is(other Member[T]) bool {
	return x.v == other.v
}
