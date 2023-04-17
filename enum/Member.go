package enum

type (
	Member[T ~string, A any] struct {
		Enum[T, A]
		place int
		v     T
	}
)

func (x Member[_, _]) String() string {
	return string(x.v)
}

func (x Member[_, _]) Name() string {
	return string(x.v)
}

func (x Member[T, A]) Is(other Member[T, A]) bool {
	return x.v == other.v
}
