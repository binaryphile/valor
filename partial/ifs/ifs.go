package ifs

type (
	Getter interface {
		Get(FieldIndex) (any, bool)
	}

	Setter[T any] interface {
		Set(FieldIndex, any) (T, error)
	}

	Partial[T any] interface {
		Getter
		Setter[T]
	}

	PSetter interface {
		Set(FieldIndex, any) error
	}

	PartialAdapter interface {
		Getter
		PSetter
	}
)
