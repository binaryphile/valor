package enum

import (
	"github.com/binaryphile/valor/optional"
)

type (
	Enum[T ~string, A any] struct {
		members map[string]Member[T, A]
	}
)

// Of creates an Enum of the given string values.
// It also returns a function for generating members.
// The function is intended for use by the enum creator and
// should be discarded or not exported in order to seal the enum.
func Of[T ~string, A any](items ...T) (Enum[T, A], func(T) Member[T, A]) {
	members := make(map[string]Member[T, A])

	enum := Enum[T, A]{
		members: members,
	}

	for i, item := range items {
		name := string(item)

		members[name] = Member[T, A]{
			enum:  enum,
			place: i,
			value: item,
		}
	}

	return enum, func(item T) Member[T, A] {
		name := string(item)

		member := Member[T, A]{
			enum:  enum,
			place: len(members),
			value: item,
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
		names[member.place] = string(member.value)
	}

	return names
}

func (x Enum[T, A]) Member(name string) optional.Value[Member[T, A]] {
	return optional.OfIndex(x.members, name)
}
