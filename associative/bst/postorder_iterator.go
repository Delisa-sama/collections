package bst

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// postOrderIterator представляет итератор для pre-order обхода BST.
type postOrderIterator[T any] struct {
	lastNodeVisited *node[T]
	current         *node[T]
	root            *node[T]
	isEnded         bool
	s               *stack.Stack[*node[T], *vector.Vector[*node[T]]]
}

// newPostOrderIterator создаёт новый postOrderIterator, устанавливая начальное состояние.
func newPostOrderIterator[T any](root *node[T]) *postOrderIterator[T] {
	it := &postOrderIterator[T]{
		lastNodeVisited: nil,
		current:         root,
		root:            root,
		isEnded:         false,
		s:               stack.NewStack(vector.NewVector[*node[T]]),
	}
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *postOrderIterator[T]) HasNext() bool {
	return !it.isEnded || !it.s.IsEmpty() || it.current != nil
}

// Next перемещает итератор к следующему элементу.
func (it *postOrderIterator[T]) Next() {
	if it.s.IsEmpty() && it.lastNodeVisited == it.root {
		it.isEnded = true
		return
	}
	for it.HasNext() {
		if it.current != nil {
			it.s.Push(it.current)
			it.current = it.current.Left
		} else {
			topNode := it.s.Top()
			if topNode.Right != nil && it.lastNodeVisited != topNode.Right {
				it.current = topNode.Right
			} else {
				it.lastNodeVisited = topNode
				it.s.Pop()
				return // visit
			}
		}
	}
}

// Value возвращает текущее значение узла.
func (it *postOrderIterator[T]) Value() T {
	return it.lastNodeVisited.Value
}

// Ptr возвращает указатель на текущее значение узла.
func (it *postOrderIterator[T]) Ptr() *T {
	return &it.lastNodeVisited.Value
}

// Equals сравнивает два итератора на равенство.
func (it *postOrderIterator[T]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *postOrderIterator[T]:
		return it.lastNodeVisited == a.lastNodeVisited
	case *iterators.EndIterator:
		return !it.HasNext()
	}
	panic("unknown iterator type")
}

// Copy копирует итератор.
func (it *postOrderIterator[T]) Copy() copiable.Copiable {
	return &postOrderIterator[T]{
		lastNodeVisited: it.lastNodeVisited,
		current:         it.current,
		root:            it.root,
		isEnded:         it.isEnded,
		s:               it.s.Copy(),
	}
}
