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

// FromLess возвращает компаратор, который сравнивает значения при помощи less.
func FromLess[T any](less Less[T]) Comparator[T] {
	return func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if Greater(a, b, less) {
			return 1
		}
		return 0
	}
}

// Less - это функция проверяющая что x < y.
type Less[T any] func(x, y T) bool

// DefaultLess возвращает функцию сравнения для упорядочиваемых типов.
func DefaultLess[T cmp.Ordered]() Less[T] {
	return func(x, y T) bool {
		return x < y
	}
}

// Greater проверяет что x > y используя less.
func Greater[T any](x, y T, less Less[T]) bool {
	return less(y, x)
}

// Equal проверяет что x == y используя less.
func Equal[T any](x, y T, less Less[T]) bool {
	return !less(x, y) && !less(y, x)
}
