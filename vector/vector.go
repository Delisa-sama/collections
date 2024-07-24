package vector

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Vector представляет собой динамический массив (вектор).
type Vector[T any] struct {
	s []T
}

// NewVector создает новый Vector и заполняет его переданными элементами.
func NewVector[T any](items ...T) *Vector[T] {
	l := &Vector[T]{
		s: items,
	}
	return l
}

// Size возвращает количество элементов в векторе.
func (l *Vector[T]) Size() uint {
	return uint(len(l.s))
}

// IsEmpty проверяет что вектор пустой.
func (l *Vector[T]) IsEmpty() bool {
	return len(l.s) == 0
}

// Begin возвращает итератор на первый элемент вектора.
func (l *Vector[T]) Begin() interfaces.RandomAccessIterator[T] {
	return newIterator(&l.s, 0)
}

// End возвращает итератор на последний элемент вектора.
func (l *Vector[T]) End() interfaces.RandomAccessIterator[T] {
	return newIterator(&l.s, len(l.s)-1)
}

// Append добавляет новый элемент в конец вектора.
func (l *Vector[T]) Append(value T) {
	l.s = append(l.s, value)
}
