package avltree

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/pair"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// preOrderIterator представляет итератор для pre-order обхода AVL дерева.
type preOrderIterator[K any, V any] struct {
	current *node[K, V]
	s       *stack.Stack[*node[K, V], *vector.Vector[*node[K, V]]]
}

// newPreOrderIterator создаёт новый preOrderIterator, устанавливая начальное состояние.
func newPreOrderIterator[K any, V any](root *node[K, V]) *preOrderIterator[K, V] {
	it := &preOrderIterator[K, V]{
		current: root,
		s:       stack.NewStack(vector.NewVector[*node[K, V]]),
	}
	it.s.Push(root)
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *preOrderIterator[K, V]) HasNext() bool {
	return !it.s.IsEmpty()
}

// Next перемещает итератор к следующему элементу.
func (it *preOrderIterator[K, V]) Next() {
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
func (it *preOrderIterator[K, V]) Value() pair.Pair[K, V] {
	return pair.NewPair(it.current.Key, it.current.Value)
}

// Equals сравнивает два итератора на равенство.
func (it *preOrderIterator[K, V]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *preOrderIterator[K, V]:
		return it.current == a.current
	case *iterators.EndIterator:
		return !it.HasNext() && it.current == nil
	}
	panic("unknown iterator type")
}

// Copy копирует итератор.
func (it *preOrderIterator[K, V]) Copy() copiable.Copiable {
	return &preOrderIterator[K, V]{
		current: it.current,
		s:       it.s.Copy(),
	}
}
