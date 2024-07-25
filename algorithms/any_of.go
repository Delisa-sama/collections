package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

func AnyOf[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool {
	_, found := FindIf(begin, end, predicate)
	return found
}
