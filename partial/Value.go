package partial

import (
	"github.com/binaryphile/valor/partial/ifs"
	"github.com/bits-and-blooms/bitset"
)

type (
	Value[T ifs.PartialAdapter] struct {
		fieldMask *bitset.BitSet
		value     T
	}
)

func Of[T ifs.PartialAdapter](value T, fieldMask *bitset.BitSet) Value[T] {
	return Value[T]{
		fieldMask: fieldMask,
		value:     value,
	}
}

func (x Value[T]) Get(index FieldIndex) (_ any, _ bool) {
	if !x.fieldMask.Test(index) {
		return
	}

	return x.value.Get(index)
}

func (x Value[T]) Set(index FieldIndex, value any) (_ Value[T], err error) {
	err = x.value.Set(index, value)
	if err != nil {
		return
	}

	x.fieldMask = x.fieldMask.Clone()

	x.fieldMask.Set(index)

	return x, nil
}
