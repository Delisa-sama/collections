package vector

import (
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
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

// NewVectorFromSlice создает новый Vector на основе переданного слайса, без копирования элементов.
func NewVectorFromSlice[T any](items []T) *Vector[T] {
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
func (l *Vector[T]) End() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// RBegin возвращает перевернутый итератор на последний элемент вектора.
func (l *Vector[T]) RBegin() interfaces.BidirectionalIterator[T] {
	return iterators.NewReverseIterator[T](newIterator(&l.s, uint(len(l.s)-1)))
}

// REnd возвращает перевернутый итератор на первый элемент вектора.
func (l *Vector[T]) REnd() interfaces.Iterator {
	return iterators.NewReverseIterator[T](l.Begin())
}

// At возвращает итератор на элемент вектора по переданному индексу.
func (l *Vector[T]) At(index uint) interfaces.RandomAccessIterator[T] {
	return newIterator(&l.s, index)
}

// PushBack добавляет новый элемент в конец вектора.
func (l *Vector[T]) PushBack(value T) {
	l.s = append(l.s, value)
}

// Back возвращает последний элемент в векторе.
func (l *Vector[T]) Back() T {
	return l.s[len(l.s)-1]
}

// PopBack удаляет последний элемент из вектора.
func (l *Vector[T]) PopBack() {
	if l.IsEmpty() {
		return
	}
	l.s = l.s[:len(l.s)-1]
}
