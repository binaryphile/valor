package enum

import (
	"fmt"
	"github.com/binaryphile/valor/optional"
)

type (
	Enum[T fmt.Stringer, A any] struct {
		members map[string]Member[T, A]
	}
)

// Of creates an Enum of the given string values.
// It also returns a function for generating members.
// The function is intended for use by the enum creator and
// should be discarded or not exported in order to seal the enum.
func Of[T fmt.Stringer, A any](items ...T) (Enum[T, A], func(T) Member[T, A]) {
	members := make(map[string]Member[T, A])

	for i, item := range items {
		name := item.String()

		members[name] = Member[T, A]{
			v:     item,
			place: i,
			name:  name,
		}
	}

	e := Enum[T, A]{
		members: members,
	}

	return e, func(item T) Member[T, A] {
		name := item.String()

		member := Member[T, A]{
			v:     item,
			Enum:  e,
			name:  name,
			place: len(members),
		}

		members[name] = member

		return member
	}
}

func (x Enum[_, _]) Includes(name string) bool {
	_, ok := x.members[name]

	return ok
}

func (x Enum[_, _]) Names() []string {
	names := make([]string, len(x.members))

	// TODO: make length safe in case of repeats messing with it
	for _, member := range x.members {
		names[member.place] = member.name
	}

	return names
}

func (x Enum[T, A]) Member(name string) optional.Value[Member[T, A]] {
	return optional.OfIndex(x.members, name)
}
