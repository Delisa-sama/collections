package bst

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// preOrderIterator представляет итератор для pre-order обхода BST.
type preOrderIterator[T any] struct {
	current *node[T]
	s       *stack.Stack[*node[T], *vector.Vector[*node[T]]]
}

// newPreOrderIterator создаёт новый preOrderIterator, устанавливая начальное состояние.
func newPreOrderIterator[T any](root *node[T]) *preOrderIterator[T] {
	it := &preOrderIterator[T]{
		current: root,
		s:       stack.NewStack(vector.NewVector[*node[T]]),
	}
	it.s.Push(root)
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *preOrderIterator[T]) HasNext() bool {
	return !it.s.IsEmpty()
}

// Next перемещает итератор к следующему элементу.
func (it *preOrderIterator[T]) Next() {
	if it.s.IsEmpty() {
		it.current = nil
		return
	}
	it.current = it.s.Top()
	it.s.Pop()
	if it.current.Right != nil {
		it.s.Push(it.current.Right)
	}
	if it.current.Left != nil {
		it.s.Push(it.current.Left)
	}
}

// Value возвращает текущее значение узла.
func (it *preOrderIterator[T]) Value() T {
	return it.current.Value
}

// Ptr возвращает указатель на текущее значение узла.
func (it *preOrderIterator[T]) Ptr() *T {
	return &it.current.Value
}

// Equals сравнивает два итератора на равенство.
func (it *preOrderIterator[T]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *preOrderIterator[T]:
		return it.current == a.current
	case *iterators.EndIterator:
		return !it.HasNext() && it.current == nil
	}
	panic("unknown iterator type")
}

// Copy копирует итератор.
func (it *preOrderIterator[T]) Copy() interfaces.Iterator {
	return &preOrderIterator[T]{
		current: it.current,
		s:       it.s.Copy(),
	}
}
