package partial

import (
	"github.com/binaryphile/valor/partial/ifs"
	"github.com/bits-and-blooms/bitset"
)

type (
	Type = string

	Value[T ifs.PartialAdapter] struct {
		v         T
		fieldMask bitset.BitSet
	}
)

func New[T ifs.PartialAdapter](value T, fieldMask bitset.BitSet) Value[T] {
	return Value[T]{
		v:         value,
		fieldMask: fieldMask,
	}
}

func (x Value[T]) Get(index ifs.FieldIndex) (_ any, ok bool) {
	if !x.fieldMask.Test(index) {
		return
	}

	return x.v.Value(index), true
}

func (x Value[T]) Set(index ifs.FieldIndex, value any) (_ Value[T], err error) {
	err = x.v.Set(index, value)
	if err != nil {
		return
	}

	x.fieldMask.Set(index)

	return x, nil
}
