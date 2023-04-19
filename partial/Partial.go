package partial

import (
	"encoding/json"
	"errors"
	"github.com/bits-and-blooms/bitset"
	"github.com/tidwall/gjson"
)

type (
	Partial[T any] struct {
		FieldMask  *bitset.BitSet
		FieldNames []string
		Value      T
	}
)

func NewPartial[T any](value T, fieldMask *bitset.BitSet, fieldNames []string) Partial[T] {
	return Partial[T]{
		FieldMask:  fieldMask,
		FieldNames: fieldNames,
		Value:      value,
	}
}

func (x Partial[T]) MarshalJSON() (_ []byte, err error) {
	paths := x.activePaths()

	byteJSON, err := json.Marshal(x.Value)
	if err != nil {
		return
	}

	results := gjson.GetManyBytes(byteJSON, paths...)
	if len(results) != len(paths) {
		return nil, errors.New("parsing error")
	}

	patch := make([]byte, 0)

	for i, result := range results {
		op := map[string]any{
			"op":    "replace",
			"path":  paths[i],
			"value": result.Value(),
		}

		byteJSON, err = json.Marshal(op)
		if err != nil {
			return
		}

		patch = append(patch, byteJSON...)
	}

	return patch, nil
}

func (x Partial[T]) activePaths() []string {
	paths := make([]string, 0)

	for i, fieldName := range x.FieldNames {
		if x.FieldMask.Test(uint(i)) {
			paths = append(paths, "/"+fieldName)
		}
	}

	return paths
}
