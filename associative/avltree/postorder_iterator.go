package avltree

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/pair"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// postOrderIterator представляет итератор для pre-order обхода AVL дерева.
type postOrderIterator[K any, V any] struct {
	lastNodeVisited *node[K, V]
	current         *node[K, V]
	root            *node[K, V]
	isEnded         bool
	s               *stack.Stack[*node[K, V], *vector.Vector[*node[K, V]]]
}

// newPostOrderIterator создаёт новый postOrderIterator, устанавливая начальное состояние.
func newPostOrderIterator[K any, V any](root *node[K, V]) *postOrderIterator[K, V] {
	it := &postOrderIterator[K, V]{
		lastNodeVisited: nil,
		current:         root,
		root:            root,
		isEnded:         false,
		s:               stack.NewStack(vector.NewVector[*node[K, V]]),
	}
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *postOrderIterator[K, V]) HasNext() bool {
	return !it.isEnded || !it.s.IsEmpty() || it.current != nil
}

// Next перемещает итератор к следующему элементу.
func (it *postOrderIterator[K, V]) Next() {
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
func (it *postOrderIterator[K, V]) Value() pair.Pair[K, V] {
	return pair.NewPair(it.lastNodeVisited.Key, it.lastNodeVisited.Value)
}

// Equals сравнивает два итератора на равенство.
func (it *postOrderIterator[K, V]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *postOrderIterator[K, V]:
		return it.lastNodeVisited == a.lastNodeVisited
	case *iterators.EndIterator:
		return !it.HasNext()
	}
	panic("unknown iterator type")
}

// Copy копирует итератор.
func (it *postOrderIterator[K, V]) Copy() interfaces.Iterator {
	return &postOrderIterator[K, V]{
		lastNodeVisited: it.lastNodeVisited,
		current:         it.current,
		root:            it.root,
		isEnded:         it.isEnded,
		s:               it.s.Copy(),
	}
}
