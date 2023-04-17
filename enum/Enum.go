package enum

import (
	"github.com/binaryphile/valor/optional"
)

type (
	Enum[T any] struct {
		members map[string]int
	}
)

// Of creates an Enum of the given string values.
// It also returns a function for generating members.
// The function is intended for use by the enum creator and
// should be discarded or not exported in order to seal the enum.
func Of[T any](names ...string) (Enum[T], func(string) Member[T]) {
	members := make(map[string]int)

	for i, item := range names {
		members[item] = i
	}

	e := Enum[T]{
		members: members,
	}

	return e, func(name string) Member[T] {
		members[name] = len(members)

		return Member[T]{
			Enum: e,
			v:    name,
		}
	}
}

func (x Enum[T]) Includes(name string) bool {
	_, ok := x.members[name]

	return ok
}

func (x Enum[T]) Names() []string {
	items := make([]string, len(x.members))

	// TODO: make length safe in case of repeats messing with it
	for item, i := range x.members {
		items[i] = item
	}

	return items
}

func (x Enum[T]) Member(name string) optional.Value[Member[T]] {
	if _, ok := x.members[name]; ok {
		return optional.OfOk(Member[T]{
			Enum: x,
			v:    name,
		})
	}

	return optional.OfNotOk[Member[T]]()
}
