package list

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// iterator представляет собой итератор для двусвязного списка.
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

// HasPrev проверяет, есть ли предыдущий элемент.
func (it *iterator[T]) HasPrev() bool {
	return it.current != nil && it.current.Prev != nil
}

// Prev переходит к предыдущему элементу.
func (it *iterator[T]) Prev() {
	it.current = it.current.Prev
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
	switch a := another.(type) {
	case *iterator[T]:
		return a.current == it.current
	case *iterators.EndIterator:
		return it.current == nil
	}
	panic("unknown iterator type")
}

// Copy копирует итератор.
func (it *iterator[T]) Copy() copiable.Copiable {
	return newIterator(it.current)
}
