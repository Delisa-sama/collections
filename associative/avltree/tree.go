package avltree

import (
	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
	"github.com/Delisa-sama/collections/pair"
)

// node представляет узел AVL дерева
type node[K any, V any] struct {
	Key    K
	Value  V
	Height int
	Left   *node[K, V]
	Right  *node[K, V]
}

// AVLTree представляет AVL дерево
type AVLTree[K any, V any] struct {
	root       *node[K, V]
	comparator comparator.Comparator[K]
}

// NewAVLTree создает новое AVL дерево
func NewAVLTree[K any, V any](comparator comparator.Comparator[K]) *AVLTree[K, V] {
	return &AVLTree[K, V]{comparator: comparator}
}

// height возвращает высоту узла
func (tree *AVLTree[K, V]) height(n *node[K, V]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// rightRotate выполняет правое вращение
func (tree *AVLTree[K, V]) rightRotate(y *node[K, V]) *node[K, V] {
	x := y.Left
	xRight := x.Right

	x.Right = y
	y.Left = xRight

	y.Height = max(tree.height(y.Left), tree.height(y.Right)) + 1
	x.Height = max(tree.height(x.Left), tree.height(x.Right)) + 1

	return x
}

// leftRotate выполняет левое вращение
func (tree *AVLTree[K, V]) leftRotate(x *node[K, V]) *node[K, V] {
	y := x.Right
	yLeft := y.Left

	y.Left = x
	x.Right = yLeft

	x.Height = max(tree.height(x.Left), tree.height(x.Right)) + 1
	y.Height = max(tree.height(y.Left), tree.height(y.Right)) + 1

	return y
}

// getBalance получает балансирующий фактор узла
func (tree *AVLTree[K, V]) getBalance(n *node[K, V]) int {
	if n == nil {
		return 0
	}
	return tree.height(n.Left) - tree.height(n.Right)
}

// Insert вставляет новый ключ в дерево
func (tree *AVLTree[K, V]) Insert(key K, value V) {
	tree.root = tree.insert(tree.root, key, value)
}

func (tree *AVLTree[K, V]) insert(n *node[K, V], key K, value V) *node[K, V] {
	if n == nil {
		return &node[K, V]{Key: key, Value: value, Height: 1}
	}

	if tree.comparator(key, n.Key) < 0 {
		n.Left = tree.insert(n.Left, key, value)
	} else if tree.comparator(key, n.Key) > 0 {
		n.Right = tree.insert(n.Right, key, value)
	} else {
		n.Value = value
		return n
	}

	n.Height = 1 + max(tree.height(n.Left), tree.height(n.Right))

	balance := tree.getBalance(n)

	// Левый левый случай
	if balance > 1 && tree.comparator(key, n.Left.Key) < 0 {
		return tree.rightRotate(n)
	}

	// Правый правый случай
	if balance < -1 && tree.comparator(key, n.Right.Key) > 0 {
		return tree.leftRotate(n)
	}

	// Левый правый случай
	if balance > 1 && tree.comparator(key, n.Left.Key) > 0 {
		n.Left = tree.leftRotate(n.Left)
		return tree.rightRotate(n)
	}

	// Правый левый случай
	if balance < -1 && tree.comparator(key, n.Right.Key) < 0 {
		n.Right = tree.rightRotate(n.Right)
		return tree.leftRotate(n)
	}

	return n
}

// Find ищет значение по ключу
func (tree *AVLTree[K, V]) Find(key K) (V, bool) {
	n := tree.find(tree.root, key)
	if n == nil {
		var zero V
		return zero, false
	}
	return n.Value, true
}

func (tree *AVLTree[K, V]) find(n *node[K, V], key K) *node[K, V] {
	if n == nil || tree.comparator(key, n.Key) == 0 {
		return n
	}

	if tree.comparator(key, n.Key) < 0 {
		return tree.find(n.Left, key)
	}
	return tree.find(n.Right, key)
}

// InOrderBegin возвращает итератор для in-order обхода с начала.
func (tree *AVLTree[K, V]) InOrderBegin() interfaces.ValueIterator[pair.Pair[K, V]] {
	return newInOrderIterator(tree.root)
}

// InOrderEnd возвращает конечный итератор для in-order обхода.
func (tree *AVLTree[K, V]) InOrderEnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// PreOrderBegin возвращает итератор для pre-order обхода с начала.
func (tree *AVLTree[K, V]) PreOrderBegin() interfaces.ValueIterator[pair.Pair[K, V]] {
	return newPreOrderIterator(tree.root)
}

// PreOrderEnd возвращает конечный итератор для pre-order обхода.
func (tree *AVLTree[K, V]) PreOrderEnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// PostOrderBegin возвращает итератор для post-order обхода с начала.
func (tree *AVLTree[K, V]) PostOrderBegin() interfaces.ValueIterator[pair.Pair[K, V]] {
	return newPostOrderIterator(tree.root)
}

// PostOrderEnd возвращает конечный итератор для post-order обхода.
func (tree *AVLTree[K, V]) PostOrderEnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// Copy копирует дерево.
func (tree *AVLTree[K, V]) Copy() copiable.Copiable {
	treeCopy := NewAVLTree[K, V](tree.comparator)
	algorithms.ForEach(tree.PreOrderBegin(), tree.PreOrderEnd(), func(p pair.Pair[K, V]) {
		treeCopy.Insert(p.First, p.Second)
	})
	return treeCopy
}
