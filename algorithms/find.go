package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

func FindC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	cmp comparator.Comparator[T],
) (interfaces.ValueIterator[T], bool) {
	return FindIf(begin, end, func(v T) bool {
		return cmp(v, value) == 0
	})
}

func Find[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) (interfaces.ValueIterator[T], bool) {
	return FindIf(begin, end, func(v T) bool {
		return v == value
	})
}

func FindIf[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) (interfaces.ValueIterator[T], bool) {
	for ; !begin.Equals(end); begin.Next() {
		if predicate(begin.Value()) {
			return begin, true
		}
	}
	return nil, false
}

func FindIfNot[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) (interfaces.ValueIterator[T], bool) {
	for ; !begin.Equals(end); begin.Next() {
		if !predicate(begin.Value()) {
			return begin, true
		}
	}
	return nil, false
}
