package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

func FindFirstOfC[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
	cmp comparator.Comparator[T],
) (interfaces.ValueIterator[T], bool) {
	return FindFirstOfIf(
		begin, end,
		sBegin, sEnd,
		func(a, b T) bool {
			return cmp(a, b) == 0
		},
	)
}

func FindFirstOf[T comparable](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
) (interfaces.ValueIterator[T], bool) {
	return FindFirstOfIf(
		begin, end,
		sBegin, sEnd,
		func(a, b T) bool {
			return a == b
		},
	)
}

func FindFirstOfIf[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
	predicate binaryPredicate[T],
) (interfaces.ValueIterator[T], bool) {
	for ; !begin.Equals(end); begin.Next() {
		for ; !sBegin.Equals(sEnd); sBegin.Next() {
			if predicate(begin.Value(), sBegin.Value()) {
				return begin, true
			}
		}
	}
	return nil, false
}
