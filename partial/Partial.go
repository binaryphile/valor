package partial

import (
	"encoding/json"
	"errors"
	"github.com/bits-and-blooms/bitset"
	"github.com/tidwall/gjson"
)

type (
	Partial[T any] struct {
		fieldMask *bitset.BitSet
		paths     []string
		value     T
	}
)

func NewPartial[T any](value T, fieldMask *bitset.BitSet, fieldNames []string) Partial[T] {
	paths := make([]string, 0, len(fieldNames))

	for _, name := range fieldNames {
		paths = append(paths, "/"+name)
	}

	return Partial[T]{
		fieldMask: fieldMask,
		paths:     paths,
		value:     value,
	}
}

func (x Partial[T]) MarshalJSON() (_ []byte, err error) {
	var results []gjson.Result

	paths := x.activePaths()

	byteJSON, err := json.Marshal(x.value)
	if err != nil {
		return
	}

	results = gjson.GetManyBytes(byteJSON, paths...)
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

	for i, path := range x.paths {
		if x.fieldMask.Test(uint(i)) {
			paths = append(paths, path)
		}
	}

	return paths
}
