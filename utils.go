package containers

import (
	"cmp"

	"github.com/Delisa-sama/collections/interfaces"
)

// Comparator - это функция, сравнивающая два элемента типа T и возвращающая результат сравнения.
type Comparator[T any] func(a, b T) int

// DefaultComparator возвращает функцию сравнения для упорядочиваемых типов.
func DefaultComparator[T cmp.Ordered]() func(x, y T) int {
	return cmp.Compare[T]
}

// EqualsByIterators проверяет равенство двух контейнеров путем сравнения их итераторов.
func EqualsByIterators[T any](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T], cmp Comparator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !((a.HasNext() || b.HasNext()) && !(a.HasNext() && b.HasNext())) // XOR(a.HasNext(), b.HasNext())
}
