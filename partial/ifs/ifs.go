package ifs

import "github.com/binaryphile/valor/enum"

type (
	FieldIndex = uint

	Type enum.Enum[string]

	Getter interface {
		Get(FieldIndex) (any, bool)
	}

	Setter[T any] interface {
		Set(FieldIndex, any) (T, error)
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
