package vector

import (
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
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
	it.current++
}

// HasPrev проверяет, есть ли предыдущий элемент.
func (it *iterator[T]) HasPrev() bool {
	return it.indexInBounds(it.current - 1)
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

// At возвращает указатель на элемент по заданному индексу и сдвигает итератор на позицию индекса.
func (it *iterator[T]) At(index uint) (*T, bool) {
	if !it.indexInBounds(index) {
		return nil, false
	}
	it.current = index
	return &(*it.s)[index], true
}

// Shift смещает итератор на заданное количество элементов.
// Если смещение положительное - смещает вперед, если отрицательное - назад.
func (it *iterator[T]) Shift(offset int) {
	newIndex := it.current
	if offset < 0 {
		newIndex -= uint(0 - offset)
	} else {
		newIndex += uint(offset)
	}
	if it.indexInBounds(newIndex) {
		it.current = newIndex
	}
}

// Equals проверяет, равен ли данный итератор другому итератору.
// Итераторы равны если указывают на один и тот же адрес.
func (it *iterator[T]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *iterator[T]:
		return &(*a.s)[a.current] == &(*it.s)[it.current]
	case *iterators.EndIterator:
		return !it.indexInBounds(it.current)
	}
	panic("unknown iterator type")
}

func (it *iterator[T]) indexInBounds(index uint) bool {
	return len(*it.s) > 0 && index <= uint(len(*it.s)-1)
}
