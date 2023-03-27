// Copyright 2022 phelmkamp. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package optional

import (
	"encoding/json"
	"reflect"
)

// Value either contains a value (ok) or nothing (not ok).
type Value[T any] struct {
	v  T
	ok bool
}

// Of creates a Value of v if ok is true.
// This aids interoperability with return values
// that follow the "comma ok" idiom.
func Of[T any](v T, ok bool) (zero Value[T]) {
	return ifThenElse(ok, Value[T]{v: v, ok: ok}, zero)
}

// OfOk creates an ok Value of v.
func OfOk[T any](v T) Value[T] {
	return Value[T]{v: v, ok: true}
}

// OfPointer creates a Value of the pointer p.
// Returns a not-ok Value if p is nil.
func OfPointer[T any](p *T) Value[*T] {
	return Of(p, p != nil)
}

// OfPointee creates a Value of the target of pointer p.
// Returns a not-ok Value if p is nil.
func OfPointee[T any](p *T) (zero Value[T]) {
	return ifThenElseDo(p == nil, zero, func() Value[T] {
		return OfOk(*p)
	})
}

// OfAssert performs the type assertion x.(T) and creates a Value of the result.
// Returns a not-ok Value if the type assertion fails.
func OfAssert[T, Tx any](x Tx) Value[T] {
	v, ok := any(x).(T)
	return Of(v, ok)
}

// OfIndex performs the map index m[k] and creates a Value of the result.
// Returns a not-ok Value if k is not present in the map.
func OfIndex[K comparable, V any, M ~map[K]V](m M, k K) Value[V] {
	v, ok := m[k]
	return Of(v, ok)
}

// OfReceive performs a blocking receive on ch and creates a Value of the result.
// Returns a not-ok Value if ch is nil or closed.
func OfReceive[T any](ch <-chan T) Value[T] {
	if ch == nil {
		return OfNotOk[T]()
	}
	v, ok := <-ch
	return Of(v, ok)
}

// OfNotOk creates a Value that is not ok.
// This aids in comparisons, enabling the use of Value in switch statements.
func OfNotOk[T any]() (zero Value[T]) {
	return
}

// OfNonZero creates an ok Value if nonzero otherwise not ok.
func OfNonZero[T comparable](value T) (zero Value[T]) {
	return Of(value, value != zero.v)
}

// OfFunc creates an ok Value if value is non-nil otherwise not ok.
// It's more efficient to use a non-reflection-based method if you need performance.
// E.g. Of(myFunc, myFunc != nil)
func OfFunc[T any](value T) Value[T] {
	// 🤮
	return Of(value, !reflect.ValueOf(&value).Elem().IsZero())
}

// IsOk returns whether v contains a value.
func (val Value[T]) IsOk() bool {
	return val.ok
}

// Ok sets dst to the underlying value if ok.
// Returns true if ok, false if not ok.
func (val Value[T]) Ok(dst *T) bool {
	if val.ok {
		*dst = val.v
	}

	return val.ok
}

// MustOk is like Ok but panics if not ok.
// This simplifies access to the underlying value
// in cases where it's known that val is ok.
func (val Value[T]) MustOk() T {
	return ifThenElseDo(val.ok, val.v, func() T {
		panic("Value.MustOk(): not ok")
	})
}

// Or returns the underlying value if ok, or def if not ok.
func (val Value[T]) Or(def T) T {
	return ifThenElse(val.ok, val.v, def)
}

// SelfOr returns the optional value if ok, or def if not ok.
func (val Value[T]) SelfOr(def Value[T]) Value[T] {
	return ifThenElse(val.ok, val, def)
}

// OrZero returns the underlying value if ok, or the zero value if not ok.
func (val Value[T]) OrZero() T {
	return val.v
}

// OrTake returns the underlying value if ok, or the result of f if not ok.
func (val Value[T]) OrTake(f func() T) T {
	return ifThenElseDo(val.ok, val.v, f)
}

// SelfOrTake returns the optional value if ok, or the result of f if not ok.
func (val Value[T]) SelfOrTake(f func() Value[T]) Value[T] {
	return ifThenElseDo(val.ok, val, f)
}

// OfOk creates an ok Value of the underlying value.
// This aids in comparisons, enabling the use of val in switch statements.
func (val Value[T]) OfOk() Value[T] {
	return OfOk(val.v)
}

// Do calls f with the underlying value if ok.
// Does nothing if not ok.
func (val Value[T]) Do(f func(T)) Value[T] {
	if val.ok {
		f(val.v)
	}

	return val
}

// Filter returns val if f returns true for the underlying value.
// Otherwise returns a not-ok Value.
func (val Value[T]) Filter(f func(T) bool) (zero Value[T]) {
	return ifThenElse(val.ok && f(val.v), val, zero)
}

// MarshalJSON encodes val as JSON.
// Marshals the underlying value if ok, the literal null if not ok.
func (val Value[T]) MarshalJSON() ([]byte, error) {
	if !val.ok {
		return []byte("null"), nil
	}

	return json.Marshal(val.v)
}

// UnmarshalJSON decodes data into val.
// val will be ok if the underlying value was unmarshaled successfully.
// Does nothing if data is the literal null.
func (val *Value[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		// by convention, null is no-op
		return nil
	}
	// unmarshal into temp first in case of error
	var temp T
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	val.v, val.ok = temp, true
	return nil
}

// Unpack returns the underlying value and whether it is ok.
// This aids in assigning to variables or function arguments.
func (val Value[T]) Unpack() (T, bool) {
	return val.v, val.ok
}

// Map returns a Value of the result of f on the underlying value.
// Returns a not-ok Value if val is not ok.
func Map[T, T2 any](f func(T) T2, val Value[T]) (zero Value[T2]) {
	return ifThenElseDo(!val.ok, zero, func() Value[T2] {
		return OfOk(f(val.v))
	})
}

// FlatMap returns the result of f on the underlying value.
// Returns a not-ok Value if val is not ok.
func FlatMap[T, T2 any](f func(T) Value[T2], val Value[T]) (zero Value[T2]) {
	return ifThenElseDo(!val.ok, zero, func() Value[T2] {
		return f(val.v)
	})
}

// Contains returns whether the underlying value equals v.
// Returns false if val is not ok.
func Contains[T comparable](v T, val Value[T]) bool {
	return val.ok && val.v == v
}

// ZipWith calls f with the underlying values of val and val2 and returns a Value of the result.
// Returns a not-ok Value if either val or val2 is not ok.
func ZipWith[T, T2, T3 any](val Value[T], val2 Value[T2], f func(T, T2) T3) (zero Value[T3]) {
	return ifThenElseDo(!(val.ok && val2.ok), zero, func() Value[T3] {
		return OfOk(f(val.v, val2.v))
	})
}

// UnzipWith calls f with the underlying value of val and returns Values of the result.
// Does nothing and returns not-ok Values if val is not ok.
func UnzipWith[T, T2, T3 any](val Value[T], f func(T) (T2, T3)) (val2 Value[T2], val3 Value[T3]) {
	if val.ok {
		val2.v, val3.v = f(val.v)
		val2.ok, val3.ok = true, true
	}
	return
}

// Flatten returns the underlying Value of val.
// Returns a not-ok Value if val is not ok.
func Flatten[T any](val Value[Value[T]]) Value[T] {
	return val.v
}

func IsOk[T any](val Value[T]) bool {
	return val.IsOk()
}

func MustOk[T any](val Value[T]) T {
	return val.MustOk()
}

func FilterOk[T any](opts []Value[T]) []Value[T] {
	result := make([]Value[T], 0, len(opts))

	for _, option := range opts {
		if option.IsOk() {
			result = append(result, option)
		}
	}

	return result
}

func ToValues[T any](opts []Value[T]) []T {
	result := make([]T, len(opts))

	for i, option := range opts {
		result[i] = option.MustOk()
	}

	return result
}

func FilterOkValues[T any](opts []Value[T]) []T {
	result := make([]T, 0, len(opts))

	for _, opt := range opts {
		if opt, ok := opt.Unpack(); ok {
			result = append(result, opt)
		}
	}

	return result
}

func Do[T any](f func(T), opts []Value[T]) {
	for _, opt := range opts {
		_ = opt.Do(f)
	}
}

func DoWith[T any](item T, opts []Value[func(T)]) {
	for _, opt := range opts {
		_ = opt.Do(func(f func(T)) {
			f(item)
		})
	}
}

func ifThenElseDo[T any](condition bool, first T, f func() T) T {
	if condition {
		return first
	}

	return f()
}

func ifThenElse[T any](condition bool, first, second T) T {
	if condition {
		return first
	}

	return second
}
