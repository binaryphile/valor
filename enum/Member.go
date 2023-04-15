package enum

type (
	Member[T ~string] struct {
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
