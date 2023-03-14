package ifs

type (
	FieldIndex = uint

	Type = string

	Getter interface {
		Get(FieldIndex) (any, bool)
	}

	Setter[T any] interface {
		Set(FieldIndex, any) (Partial[T], error)
	}

	Completer[T any] interface {
		Complete() (T, bool)
	}

	Partial[T any] interface {
		Completer[T]
		Getter
		Setter[T]
	}

	Valuer interface {
		Value(FieldIndex) any
	}

	Typer interface {
		Type(index FieldIndex) Type
	}

	PSetter interface {
		Set(FieldIndex, any) error
	}

	PartialAdapter interface {
		Valuer
		Typer
		PSetter
	}
)
