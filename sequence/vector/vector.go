package vector

import (
	"github.com/Delisa-sama/collections/copiable"
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
func (l *Vector[T]) End() interfaces.RandomAccessIterator[T] {
	return newIterator(&l.s, uint(len(l.s)))
}

// RBegin возвращает перевернутый итератор на последний элемент вектора.
func (l *Vector[T]) RBegin() interfaces.BidirectionalIterator[T] {
	return iterators.NewReverseIterator[T](newIterator(&l.s, uint(len(l.s)-1)))
}

// REnd возвращает перевернутый итератор на первый элемент вектора.
func (l *Vector[T]) REnd() interfaces.BidirectionalIterator[T] {
	return iterators.NewReverseIterator[T](l.Begin())
}

// At возвращает элемент вектора по переданному индексу.
func (l *Vector[T]) At(index uint) T {
	return l.s[index]
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

// Copy копирует вектор.
func (l *Vector[T]) Copy() copiable.Copiable {
	sliceCopy := make([]T, len(l.s))
	copy(sliceCopy, l.s)
	return NewVectorFromSlice(sliceCopy)
}

// Erase удаляет элементы в диапазоне [begin, end) из вектора.
func (l *Vector[T]) Erase(begin, end interfaces.Iterator) {
	b, bOk := begin.(*iterator[T])
	e, eOk := end.(*iterator[T])
	if !bOk || !eOk {
		panic("unknown iterator type")
	}

	l.s = append(l.s[:b.Index()], l.s[e.Index():]...)
}

// RemoveRange удаляет элементы в диапазоне [from, to).
func (l *Vector[T]) RemoveRange(from, to uint) {
	l.s = append(l.s[:from], l.s[to:]...)
}

// Remove удаляет элемент по индексу.
func (l *Vector[T]) Remove(index uint) {
	l.s = append(l.s[:index], l.s[index+1:]...)
}
