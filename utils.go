package containers

import (
	"fmt"

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

func Print[T any](a interfaces.ForwardIterator[T]) {
	for ; a.HasNext(); a.Next() {
		fmt.Println(a.Value())
	}
}
