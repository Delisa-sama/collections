package comparator

import (
	"cmp"
)

// Comparator - это функция, сравнивающая два элемента типа T и возвращающая результат сравнения.
type Comparator[T any] func(a, b T) int

// DefaultComparator возвращает функцию сравнения для упорядочиваемых типов.
func DefaultComparator[T cmp.Ordered]() func(x, y T) int {
	return cmp.Compare[T]
}
