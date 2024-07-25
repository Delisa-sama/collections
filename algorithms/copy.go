package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

func Copy[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T] {
	for !begin.Equals(end) {
		value := begin.Value()
		*destBegin.Ptr() = value

		begin.Next()
		destBegin.Next()
	}

	return destBegin
}

func CopyIf[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
	predicate unaryPredicate[T],
) interfaces.PointerIterator[T] {
	for !begin.Equals(end) {
		value := begin.Value()
		if predicate(value) {
			*destBegin.Ptr() = value
			destBegin.Next()
		}

		begin.Next()
	}

	return destBegin
}
