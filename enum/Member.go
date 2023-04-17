package enum

type (
	Member[T ~string, A any] struct {
		enum  Enum[T, A]
		place int
		value T
	}
)

func (x Member[_, _]) String() string {
	return string(x.value)
}

func (x Member[_, _]) Error() string {
	return string(x.value)
}

func (x Member[_, _]) Name() string {
	return string(x.value)
}

func (x Member[T, A]) Is(other Member[T, A]) bool {
	return x.value == other.value
}

func (x Member[T, A]) Enum() Enum[T, A] {
	return x.enum
}
