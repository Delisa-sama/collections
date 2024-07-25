package set

import (
	"github.com/elliotchance/orderedmap/v2"

	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// iterator представляет собой итератор для односвязного списка.
type iterator[K comparable] struct {
	current *orderedmap.Element[K, struct{}]
}

// newIterator создает новый итератор.
func newIterator[K comparable](n *orderedmap.Element[K, struct{}]) *iterator[K] {
	return &iterator[K]{
		current: n,
	}
}

// HasNext проверяет, есть ли следующий элемент.
func (it *iterator[K]) HasNext() bool {
	return it.current != nil && it.current.Next() != nil
}

// Next переходит к следующему элементу.
func (it *iterator[K]) Next() {
	it.current = it.current.Next()
}

// HasPrev проверяет, есть ли предыдущий элемент.
func (it *iterator[K]) HasPrev() bool {
	return it.current != nil && it.current.Prev() != nil
}

// Prev переходит к предыдущему элементу.
func (it *iterator[K]) Prev() {
	it.current = it.current.Prev()
}

// Value возвращает текущее значение итератора.
func (it *iterator[K]) Value() K {
	return it.current.Key
}

// Ptr возвращает указатель на текущее значение итератора.
func (it *iterator[K]) Ptr() *K {
	return &it.current.Key
}

// Equals проверяет, равен ли данный итератор другому итератору.
func (it *iterator[K]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *iterator[K]:
		return a.current == it.current
	case *iterators.EndIterator:
		return it.current == nil
	}
	panic("unknown iterator type")
}
