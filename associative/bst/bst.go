package bst

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// node представляет узел в бинарном поисковом дереве (BST).
type node[T any] struct {
	Value  T
	Parent *node[T]
	Left   *node[T]
	Right  *node[T]
}

// BST представляет бинарное поисковое дерево.
type BST[T any] struct {
	root *node[T]
	size uint
	comp comparator.Comparator[T]
}

// NewBST создает новое BST с заданным компаратором.
func NewBST[T any](comp comparator.Comparator[T], values ...T) *BST[T] {
	bst := &BST[T]{comp: comp}
	for i := range values {
		bst.Insert(values[i])
	}
	return bst
}

// Size возвращает размер BST.
func (t *BST[T]) Size() uint {
	return t.size
}

// Insert вставляет значение в BST.
func (t *BST[T]) Insert(value T) {
	if t.root == nil {
		t.root = &node[T]{
			Value:  value,
			Parent: nil,
			Left:   nil,
			Right:  nil,
		}
	} else {
		t.insert(t.root, &node[T]{
			Value:  value,
			Parent: nil,
			Left:   nil,
			Right:  nil,
		})
	}
	t.size++
}

func (t *BST[T]) insert(x *node[T], z *node[T]) {
	for x != nil {
		if t.comp(z.Value, x.Value) >= 0 {
			if x.Right != nil {
				x = x.Right
			} else {
				z.Parent = x
				x.Right = z
				break
			}
		} else {
			if x.Left != nil {
				x = x.Left
			} else {
				z.Parent = x
				x.Left = z
				break
			}
		}
	}
}

// Find возвращает указатель на элемент дерева, если элемент не найден, возвращает nil.
func (t *BST[T]) Find(k T) *T {
	return t.find(t.root, k)
}

func (t *BST[T]) find(x *node[T], k T) *T {
	if x == nil {
		return nil
	}

	c := t.comp(k, x.Value)
	switch {
	case c == 0:
		return &x.Value
	case c < 0:
		return t.find(x.Left, k)
	default:
		return t.find(x.Right, k)
	}
}

// InOrderIterator возвращает итератор для in-order обхода.
func (t *BST[T]) InOrderIterator() interfaces.ForwardIterator[T] {
	return newInOrderIterator(t.root)
}
