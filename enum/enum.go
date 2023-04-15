package enum

import (
	"encoding"
	"errors"
	"fmt"
	"github.com/binaryphile/valor/tuple/two"
)

// Enum is an enumerated type.
//
// It wraps an optional.Value that is only ok if it's a member of the allowed values.
type Enum[T comparable] struct {
	v       T
	members map[T]metadata // carry allowed values for validation
}

type metadata struct {
	i    int
	name string
}

// ComparableText is a constraint that permits
// comparable types that can be marshaled to text.
type ComparableText interface {
	comparable
	encoding.TextMarshaler
}

// Of creates an Enum of the given name-value pairs.
func Of[T comparable](pairs ...two.Tuple[string, T]) Enum[T] {
	e := Enum[T]{members: make(map[T]metadata)}
	for i, m := range pairs {
		e.members[m.V2] = metadata{i: i, name: m.V}
	}
	return e
}

// OfString creates an Enum of the given string values.
func OfString[S ~string](vals ...S) Enum[S] {
	e := Enum[S]{members: make(map[S]metadata)}
	for i, v := range vals {
		e.members[v] = metadata{i: i, name: string(v)}
	}
	return e
}

// OfText creates an Enum of the given text values.
// Panics if MarshalText returns an error for one of the values.
func OfText[T ComparableText](vals ...T) Enum[T] {
	e := Enum[T]{members: make(map[T]metadata)}
	for i, v := range vals {
		text, err := v.MarshalText()
		if err != nil {
			panic(err)
		}
		e.members[v] = metadata{i: i, name: string(text)}
	}
	return e
}

// ValueOf returns an Enum that wraps v if v is a member of the allowed values.
// Returns not ok otherwise.
func (e Enum[T]) ValueOf(v T) (_ Enum[T], ok bool) {
	if _, ok = e.members[v]; !ok {
		return
	}
	e.v = v
	return e, true
}

// MustOf returns an Enum that wraps v if v is a member of the allowed values.
// Panics otherwise.
func (e Enum[T]) MustOf(v T) Enum[T] {
	if _, ok := e.members[v]; !ok {
		panic("value not in enum")
	}
	e.v = v
	return e
}

// String returns e formatted as a string.
func (e Enum[T]) String() string {
	return fmt.Sprint(e.members[e.v])
}

// Value returns the stored value.
func (e Enum[T]) Value() T {
	return e.v
}

// Values returns the allowed values.
func (e Enum[T]) Values() []T {
	s := make([]T, len(e.members))
	for v, m := range e.members {
		s[m.i] = v
	}
	return s
}

// Names returns the allowed names.
func (e Enum[T]) Names() []string {
	s := make([]string, len(e.members))
	for _, m := range e.members {
		s[m.i] = m.name
	}
	return s
}

// MarshalText returns the name of the current member.
// Returns nil if e is not ok.
func (e Enum[T]) MarshalText() (text []byte, err error) {
	return []byte(e.members[e.v].name), nil
}

// UnmarshalText sets e to wrap the member with the given name.
func (e *Enum[T]) UnmarshalText(text []byte) (err error) {
	s := string(text)
	for v, m := range e.members {
		if m.name == s {
			e.v = v
			return
		}
	}
	return errors.New("value not in enum")
}
