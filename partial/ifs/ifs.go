package ifs

type (
	PartialAdapter interface {
		Get(FieldIndex) (any, bool)
		Set(FieldIndex, any) error
	}
)
