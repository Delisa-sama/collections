package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

func CountC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	cmp comparator.Comparator[T],
) uint {
	return CountIf(begin, end, func(v T) bool {
		return cmp(v, value) == 0
	})
}

func Count[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) uint {
	return CountIf(begin, end, func(v T) bool {
		return v == value
	})
}

func CountIf[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) uint {
	var count uint
	for ; !begin.Equals(end); begin.Next() {
		if predicate(begin.Value()) {
			count++
		}
	}
	return count
}
