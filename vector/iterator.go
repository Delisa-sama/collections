package vector

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// iterator представляет собой итератор для вектора.
type iterator[T any] struct {
	s       *[]T
	current int
}

// newIterator создает новый итератор.
func newIterator[T any](s *[]T, index int) *iterator[T] {
	return &iterator[T]{
		s:       s,
		current: index,
	}
}

// HasNext проверяет, есть ли следующий элемент.
func (it *iterator[T]) HasNext() bool {
	return it.current+1 <= len(*it.s)-1
}

// Next переходит к следующему элементу.
func (it *iterator[T]) Next() {
	it.current++
}

// HasPrev проверяет, есть ли предыдущий элемент.
func (it *iterator[T]) HasPrev() bool {
	return it.current-1 >= 0
}

// Prev переходит к предыдущему элементу.
func (it *iterator[T]) Prev() {
	it.current--
}

// Value возвращает текущее значение итератора.
func (it *iterator[T]) Value() T {
	return (*it.s)[it.current]
}

// Ptr возвращает указатель на текущее значение итератора.
func (it *iterator[T]) Ptr() *T {
	return &(*it.s)[it.current]
}

// At проверяет, доступен ли элемент по заданному индексу.
func (it *iterator[T]) At(index uint) bool {
	return len(*it.s)-1 >= 0 && index <= uint(len(*it.s)-1)
}

// Equals проверяет, равен ли данный итератор другому итератору.
func (it *iterator[T]) Equals(another interfaces.Iterator) bool {
	if a, ok := another.(*iterator[T]); ok {
		return &(*a.s)[a.current] == &(*it.s)[it.current]
	}
	return false
}
