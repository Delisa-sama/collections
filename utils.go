package containers

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// EqualsByIterators проверяет равенство двух контейнеров путем сравнения их итераторов.
func EqualsByIterators[T any](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T], cmp comparator.Comparator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !((a.HasNext() || b.HasNext()) && !(a.HasNext() && b.HasNext())) // XOR(a.HasNext(), b.HasNext())
}

type ForEachFunc[T any] func(T)

func ForEach[T any](begin interfaces.ForwardIterator[T], end interfaces.Iterator, f ForEachFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Value())
	}
}
