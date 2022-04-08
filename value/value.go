package value

// Value either contains a value (ok) or nothing (not ok).
type Value[T any] struct {
	v  T
	ok bool
}

// Of creates a Value of v if ok is true.
// This aids interoperability with return values
// that follow the "comma ok" idiom.
func Of[T any](v T, ok bool) Value[T] {
	if ok {
		return OfOk(v)
	}
	return Value[T]{}
}

// OfOk creates an ok Value of v.
func OfOk[T any](v T) Value[T] {
	return Value[T]{v: v, ok: true}
}

// OfNotOk creates a Value that is not ok.
// This aids in comparisons, enabling the use of Value in switch statements.
func OfNotOk[T any]() Value[T] {
	return Value[T]{}
}

// IsOk returns whether v contains a value.
func (val Value[T]) IsOk() bool {
	return val.ok
}

// Ok sets dst to the underlying value if ok.
// Returns true if ok, false if not ok.
func (val Value[T]) Ok(dst *T) bool {
	if !val.IsOk() {
		return false
	}
	*dst = val.v
	return true
}

// MustOk is like Ok but panics if not ok.
// This simplifies access to the underlying value
// in cases where it's known that val is ok.
func (val Value[T]) MustOk() T {
	if !val.IsOk() {
		panic("value.MustOk(): not ok")
	}
	return val.v
}

// Or returns the underlying value if ok, or def if not ok.
func (val Value[T]) Or(def T) T {
	if val.IsOk() {
		return val.v
	}
	return def
}

// OrZero returns the underlying value if ok, or the zero value if not ok.
func (val Value[T]) OrZero() T {
	return val.v
}

// OrElse returns the underlying value if ok, or the result of f if not ok.
func (val Value[T]) OrElse(f func() T) T {
	if val.IsOk() {
		return val.v
	}
	return f()
}

// OfOk creates an ok Value of the underlying value.
// This aids in comparisons, enabling the use of val in switch statements.
func (val Value[T]) OfOk() Value[T] {
	return OfOk(val.v)
}

// Do calls f with the underlying value if ok.
// Does nothing if not ok.
func (val Value[T]) Do(f func(T)) Value[T] {
	if val.IsOk() {
		f(val.v)
	}
	return val
}

// Filter returns val if f returns true for the underlying value.
// Otherwise returns a not ok Value.
func (val Value[T]) Filter(f func(T) bool) Value[T] {
	if val.IsOk() && f(val.v) {
		return val
	}
	return Value[T]{}
}

// Take returns the current Value and sets val as not ok.
func (val *Value[T]) Take() Value[T] {
	var val2 Value[T]
	*val, val2 = val2, *val
	return val2
}

// Map returns a Value of the result of f on the underlying value.
// Returns a not ok Value if val is not ok.
func Map[T, T2 any](val Value[T], f func(T) T2) Value[T2] {
	if !val.IsOk() {
		return Value[T2]{}
	}
	return OfOk(f(val.v))
}

// FlatMap returns the result of f on the underlying value.
// Returns a not ok Value if val is not ok.
func FlatMap[T, T2 any](val Value[T], f func(T) Value[T2]) Value[T2] {
	if !val.IsOk() {
		return Value[T2]{}
	}
	return f(val.v)
}

// Contains returns whether the underlying value equals v.
// Returns false if val is not ok.
func Contains[T comparable](val Value[T], v T) bool {
	return val.IsOk() && val.v == v
}

// ZipWith calls f with the underlying values of val and val2 and returns a Value of the result.
// Returns a not ok Value if either val or val2 is not ok.
func ZipWith[T, T2, T3 any](val Value[T], val2 Value[T2], f func(T, T2) T3) Value[T3] {
	if !val.IsOk() || !val2.IsOk() {
		return Value[T3]{}
	}
	return OfOk(f(val.v, val2.v))
}

// UnzipWith calls f with the underlying value of val and returns Values of the result.
// Does nothing and returns not ok Values if val is not ok.
func UnzipWith[T, T2, T3 any](val Value[T], f func(T) (T2, T3)) (val2 Value[T2], val3 Value[T3]) {
	if val.IsOk() {
		val2.v, val3.v = f(val.v)
	}
	return
}

// Flatten returns the underlying Value of val.
// Returns a not ok Value if val is not ok.
func Flatten[T any](val Value[Value[T]]) Value[T] {
	if !val.IsOk() {
		return OfNotOk[T]()
	}
	return val.v
}
