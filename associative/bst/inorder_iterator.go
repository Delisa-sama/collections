package bst

import (
	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/sequence/vector"
)

// InOrderIterator представляет итератор для in-order обхода BST.
type InOrderIterator[T any] struct {
	current *node[T]
	s       *stack.Stack[*node[T], *vector.Vector[*node[T]]]
}

// newInOrderIterator создаёт новый InOrderIterator, устанавливая начальное состояние.
func newInOrderIterator[T any](root *node[T]) *InOrderIterator[T] {
	it := &InOrderIterator[T]{
		current: root,
		s:       stack.NewStack(vector.NewVector[*node[T]]),
	}

	it.pushLeft(root)
	it.Next()
	return it
}

// HasNext проверяет, есть ли ещё элементы для обхода.
func (it *InOrderIterator[T]) HasNext() bool {
	return !it.s.IsEmpty() || it.current != nil
}

// Next перемещает итератор к следующему элементу.
func (it *InOrderIterator[T]) Next() {
	if !it.s.IsEmpty() {
		it.current = it.s.Top()
		it.s.Pop()
		it.pushLeft(it.current.Right)
	} else {
		it.current = nil
	}
}

// Value возвращает текущее значение узла.
func (it *InOrderIterator[T]) Value() T {
	return it.current.Value
}

// Ptr возвращает указатель на текущее значение узла.
func (it *InOrderIterator[T]) Ptr() *T {
	return &it.current.Value
}

// Equals сравнивает два итератора на равенство.
func (it *InOrderIterator[T]) Equals(another interfaces.Iterator) bool {
	switch a := another.(type) {
	case *InOrderIterator[T]:
		return it.current == a.current
	case *inorderEndIterator:
		return !it.HasNext()
	}
	panic("unknown iterator type")
}

// pushLeft добавляет в стек все левые узлы, начиная с заданного узла.
func (it *InOrderIterator[T]) pushLeft(n *node[T]) {
	for n != nil {
		it.s.Push(n)
		n = n.Left
	}
}

type inorderEndIterator struct{}

func newInOrderEndIterator() interfaces.Iterator {
	return &inorderEndIterator{}
}

// Equals проверяет, равны ли два итератора.
func (it *inorderEndIterator) Equals(another interfaces.Iterator) bool {
	switch another.(type) {
	case *inorderEndIterator:
		return true
	default:
		return another.Equals(it)
	}
}

// HasNext проверяет, есть ли еще элементы для перебора.
func (it *inorderEndIterator) HasNext() bool {
	return false
}

// Next переходит к следующему элементу.
func (it *inorderEndIterator) Next() {
	return
}
