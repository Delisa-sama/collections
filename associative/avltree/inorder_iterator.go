package avltree

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/pair"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// inOrderIterator представляет итератор для in-order обхода AVL дерева.
type inOrderIterator[K any, V any] struct {
	current *node[K, V]
	s       *stack.Stack[*node[K, V], *vector.Vector[*node[K, V]]]
}

// newInOrderIterator создаёт новый inOrderIterator, устанавливая начальное состояние.
func newInOrderIterator[K any, V any](root *node[K, V]) *inOrderIterator[K, V] {
	it := &inOrderIterator[K, V]{
		current: root,
		s:       stack.NewStack(vector.NewVector[*node[K, V]]),
	}

	it.pushLeft(root)
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *inOrderIterator[K, V]) HasNext() bool {
	return !it.s.IsEmpty() || it.current != nil
}

// Next перемещает итератор к следующему элементу.
func (it *inOrderIterator[K, V]) Next() {
	if !it.s.IsEmpty() {
		it.current = it.s.Top()
		it.s.Pop()
		it.pushLeft(it.current.Right)
	} else {
		it.current = nil
	}
}

// Value возвращает текущее значение узла.
func (it *inOrderIterator[K, V]) Value() pair.Pair[K, V] {
	return pair.NewPair(it.current.Key, it.current.Value)
}

// Equals сравнивает два итератора на равенство.
func (it *inOrderIterator[K, V]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *inOrderIterator[K, V]:
		return it.current == a.current
	case *iterators.EndIterator:
		return !it.HasNext()
	}
	panic("unknown iterator type")
}

// pushLeft добавляет в стек все левые узлы, начиная с заданного узла.
func (it *inOrderIterator[K, V]) pushLeft(n *node[K, V]) {
	for n != nil {
		it.s.Push(n)
		n = n.Left
	}
}

// Copy копирует итератор.
func (it *inOrderIterator[K, V]) Copy() copiable.Copiable {
	return &inOrderIterator[K, V]{
		current: it.current,
		s:       it.s.Copy(),
	}
}
