package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

func AllOf[T any](begin interfaces.ForwardIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool {
	_, found := FindIfNot(begin, end, predicate)
	return !found
}
