package bst

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// inOrderIterator представляет итератор для in-order обхода BST.
type inOrderIterator[T any] struct {
	current *node[T]
	s       *stack.Stack[*node[T], *vector.Vector[*node[T]]]
}

// newInOrderIterator создаёт новый inOrderIterator, устанавливая начальное состояние.
func newInOrderIterator[T any](root *node[T]) *inOrderIterator[T] {
	it := &inOrderIterator[T]{
		current: root,
		s:       stack.NewStack(vector.NewVector[*node[T]]),
	}

	it.pushLeft(root)
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *inOrderIterator[T]) HasNext() bool {
	return !it.s.IsEmpty() || it.current != nil
}

// Next перемещает итератор к следующему элементу.
func (it *inOrderIterator[T]) Next() {
	if !it.s.IsEmpty() {
		it.current = it.s.Top()
		it.s.Pop()
		it.pushLeft(it.current.Right)
	} else {
		it.current = nil
	}
}

// Value возвращает текущее значение узла.
func (it *inOrderIterator[T]) Value() T {
	return it.current.Value
}

// Ptr возвращает указатель на текущее значение узла.
func (it *inOrderIterator[T]) Ptr() *T {
	return &it.current.Value
}

// Equals сравнивает два итератора на равенство.
func (it *inOrderIterator[T]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *inOrderIterator[T]:
		return it.current == a.current
	case *iterators.EndIterator:
		return !it.HasNext()
	}
	panic("unknown iterator type")
}

// pushLeft добавляет в стек все левые узлы, начиная с заданного узла.
func (it *inOrderIterator[T]) pushLeft(n *node[T]) {
	for n != nil {
		it.s.Push(n)
		n = n.Left
	}
}

// Copy копирует итератор.
func (it *inOrderIterator[T]) Copy() copiable.Copiable {
	return &inOrderIterator[T]{
		current: it.current,
		s:       it.s.Copy(),
	}
}
