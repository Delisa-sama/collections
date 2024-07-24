package forwardlist

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// iterator представляет собой итератор для односвязного списка.
type iterator[T any] struct {
	current *node[T]
}

// newIterator создает новый итератор.
func newIterator[T any](n *node[T]) *iterator[T] {
	return &iterator[T]{
		current: n,
	}
}

// HasNext проверяет, есть ли следующий элемент.
func (it *iterator[T]) HasNext() bool {
	return it.current != nil && it.current.Next != nil
}

// Next переходит к следующему элементу.
func (it *iterator[T]) Next() {
	it.current = it.current.Next
}

// Value возвращает текущее значение итератора.
func (it *iterator[T]) Value() T {
	return *it.current.Value
}

// Ptr возвращает указатель на текущее значение итератора.
func (it *iterator[T]) Ptr() *T {
	return it.current.Value
}

// Equals проверяет, равен ли данный итератор другому итератору.
func (it *iterator[T]) Equals(another interfaces.Iterator) bool {
	if a, ok := another.(*iterator[T]); ok {
		return a.current == it.current
	}
	return false
}
