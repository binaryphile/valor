package enum

import (
	"github.com/binaryphile/valor/optional"
)

// Member is an enumerated type.
//
// It wraps a value that is only ok if it's a member of the allowed values.
type (
	Enum[T ~string] map[string]int

	Member[T ~string] struct {
		Enum[T]
		v string
	}
)

// Of creates an Enum of the given string values.
func Of[T ~string](items ...T) Enum[T] {
	members := make(Enum[T])

	for i, item := range items {
		members[string(item)] = i
	}

	return members
}

func (x Enum[T]) NewMember(item T) Member[T] {
	strItem := string(item)

	x[strItem] = len(x) - 1

	return x.newMember(strItem)
}

func (x Enum[T]) Includes(item string) bool {
	_, ok := x[item]

	return ok
}

// String returns e formatted as a string.
func (x Member[T]) String() string {
	return x.v
}

// Values returns the allowed values.
func (x Enum[T]) Values() []string {
	items := make([]string, len(x))

	for item, i := range x {
		items[i] = item
	}

	return items
}

func (x Enum[T]) Member(item string) optional.Value[Member[T]] {
	_, ok := x[item]

	return optional.Of(x.newMember(item), ok)
}

func (x Enum[T]) newMember(item string) Member[T] {
	return Member[T]{
		Enum: x,
		v:    item,
	}
}
