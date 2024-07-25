package vector

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// iterator представляет собой итератор для вектора.
type iterator[T any] struct {
	s       *[]T
	current uint
}

// newIterator создает новый итератор.
func newIterator[T any](s *[]T, index uint) *iterator[T] {
	return &iterator[T]{
		s:       s,
		current: index,
	}
}

// HasNext проверяет, есть ли следующий элемент.
func (it *iterator[T]) HasNext() bool {
	return it.indexInBounds(it.current + 1)
}

// Next переходит к следующему элементу.
func (it *iterator[T]) Next() {
	if it.HasNext() {
		it.current++
	}
}

// HasPrev проверяет, есть ли предыдущий элемент.
func (it *iterator[T]) HasPrev() bool {
	return it.indexInBounds(it.current - 1)
}

// Prev переходит к предыдущему элементу.
func (it *iterator[T]) Prev() {
	if it.HasPrev() {
		it.current--
	}
}

// Value возвращает текущее значение итератора.
func (it *iterator[T]) Value() T {
	return (*it.s)[it.current]
}

// Ptr возвращает указатель на текущее значение итератора.
func (it *iterator[T]) Ptr() *T {
	return &(*it.s)[it.current]
}

// At возвращает указатель на элемент по заданному индексу и сдвигает итератор на позицию индекса.
func (it *iterator[T]) At(index uint) (*T, bool) {
	if !it.indexInBounds(index) {
		return nil, false
	}
	it.current = index
	return &(*it.s)[index], true
}

// Equals проверяет, равен ли данный итератор другому итератору.
// Итераторы равны если указывают на один и тот же адрес.
func (it *iterator[T]) Equals(another interfaces.Iterator) bool {
	if a, ok := another.(*iterator[T]); ok {
		return &(*a.s)[a.current] == &(*it.s)[it.current]
	}
	return false
}

func (it *iterator[T]) indexInBounds(index uint) bool {
	return len(*it.s) > 0 && index <= uint(len(*it.s)-1)
}
