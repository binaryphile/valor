package enum

import (
	"github.com/binaryphile/valor/optional"
)

type (
	Enum[T any] map[string]int
)

// Of creates an Enum of the given string values.
func Of[T any](items ...string) (Enum[T], func(string) Member[T]) {
	members := make(Enum[T])

	for i, item := range items {
		members[item] = i
	}

	return members, func(item string) Member[T] {
		members[item] = len(members)

		return Member[T]{
			Enum: members,
			v:    item,
		}
	}
}

func (x Enum[T]) Includes(name string) bool {
	_, ok := x[name]

	return ok
}

func (x Enum[T]) Names() []string {
	items := make([]string, len(x))

	for item, i := range x {
		items[i] = item
	}

	return items
}

func (x Enum[T]) Member(name string) optional.Value[Member[T]] {
	if _, ok := x[name]; ok {
		return optional.OfOk(Member[T]{
			Enum: x,
			v:    name,
		})
	}

	return optional.OfNotOk[Member[T]]()
}
