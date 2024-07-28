package iterators

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// ReverseIterator адаптирует BidirectionalIterator для итерирования в обратном порядке.
type ReverseIterator[T any] struct {
	it interfaces.BidirectionalIterator[T]
}

// NewReverseIterator возвращает ReverseIterator.
func NewReverseIterator[T any](it interfaces.BidirectionalIterator[T]) *ReverseIterator[T] {
	return &ReverseIterator[T]{it: it}
}

// HasNext проверяет, есть ли следующий элемент.
func (it *ReverseIterator[T]) HasNext() bool {
	return it.it.HasPrev()
}

// Next переходит к следующему элементу.
func (it *ReverseIterator[T]) Next() {
	it.it.Prev()
}

// HasPrev проверяет, есть ли предыдущий элемент.
func (it *ReverseIterator[T]) HasPrev() bool {
	return it.it.HasNext()
}

// Prev переходит к предыдущему элементу.
func (it *ReverseIterator[T]) Prev() {
	it.it.Next()
}

// Value возвращает текущее значение итератора.
func (it *ReverseIterator[T]) Value() T {
	return it.it.Value()
}

// Ptr возвращает указатель на текущее значение итератора.
func (it *ReverseIterator[T]) Ptr() *T {
	return it.it.Ptr()
}

// Equals проверяет, равен ли данный итератор другому итератору.
func (it *ReverseIterator[T]) Equals(another interfaces.Iterator) bool {
	return it.it.Equals(another)
}

// Copy копирует итератор.
func (it *ReverseIterator[T]) Copy() copiable.Copiable {
	return NewReverseIterator(copiable.Copy[interfaces.BidirectionalIterator[T]](it.it))
}
