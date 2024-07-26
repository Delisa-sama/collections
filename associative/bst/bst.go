package bst

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// node представляет узел в бинарном поисковом дереве (BST).
type node[T any] struct {
	Value  T
	Parent *node[T]
	Left   *node[T]
	Right  *node[T]
}

func (n *node[T]) next() *node[T] {
	if n.Right != nil {
		return n.Right.min()
	}
	parent := n.Parent
	current := n
	for parent != nil && current == parent.Right {
		current = parent
		parent = parent.Parent
	}
	return parent
}

func (n *node[T]) prev() *node[T] {
	if n.Left != nil {
		return n.Left.max()
	}
	parent := n.Parent
	current := n
	for parent != nil && current == parent.Left {
		current = parent
		parent = parent.Parent
	}
	return parent
}

func (n *node[T]) min() *node[T] {
	if n.Left == nil {
		return n
	}
	return n.Left.min()
}

func (n *node[T]) max() *node[T] {
	if n.Right == nil {
		return n
	}
	return n.Right.max()
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

// IsEmpty проверяет что дерево пустое.
func (t *BST[T]) IsEmpty() bool {
	return t.size == 0
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
		if t.comp(z.Value, x.Value) > 0 {
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
	n := t.find(t.root, k)
	if n == nil {
		return nil
	}
	return &n.Value
}

func (t *BST[T]) find(x *node[T], k T) *node[T] {
	if x == nil {
		return nil
	}

	c := t.comp(k, x.Value)
	switch {
	case c == 0:
		return x
	case c < 0:
		return t.find(x.Left, k)
	default:
		return t.find(x.Right, k)
	}
}

// Delete выполняет рекурсивный поиск элемента, и удаляет его если найден.
// Возвращает true в случае успешного удаления.
func (t *BST[T]) Delete(k T) bool {
	v := t.find(t.root, k)
	if v == nil {
		return false
	}
	t.delete(v)
	return true
}

func (t *BST[T]) delete(v *node[T]) {
	p := v.Parent
	if v.Left == nil && v.Right == nil {
		if p.Left == v {
			p.Left = nil
		}
		if p.Right == v {
			p.Right = nil
		}
	} else if v.Left == nil || v.Right == nil {
		if v.Left == nil {
			if p.Left == v {
				p.Left = v.Right
			} else {
				p.Right = v.Right
			}
			v.Right.Parent = p
		} else {
			if p.Left == v {
				p.Left = v.Left
			} else {
				p.Right = v.Left
			}
			v.Left.Parent = p
		}
	} else {
		successor := v.next()
		v.Value = successor.Value
		if successor.Parent.Left == successor {
			successor.Parent.Left = successor.Right
			if successor.Right != nil {
				successor.Right.Parent = successor.Parent
			}
		} else {
			successor.Parent.Right = successor.Right
			if successor.Right != nil {
				successor.Right.Parent = successor.Parent
			}
		}
	}
}

// Min возвращает минимальный элемент в дереве.
// Возможна паника при поиске в пустом дереве.
func (t *BST[T]) Min() T {
	if t.root == nil || t.IsEmpty() {
		panic("searching min in empty tree")
	}
	n := t.root.min()
	return n.Value
}

// Max возвращает максимальный элемент в дереве.
// Возможна паника при поиске в пустом дереве.
func (t *BST[T]) Max() T {
	if t.root == nil || t.IsEmpty() {
		panic("searching max in empty tree")
	}
	n := t.root.max()
	return n.Value
}

// InOrderBegin возвращает итератор для in-order обхода с начала.
func (t *BST[T]) InOrderBegin() interfaces.ForwardIterator[T] {
	return newInOrderIterator(t.root)
}

// InOrderEnd возвращает конечный итератор для in-order обхода.
func (t *BST[T]) InOrderEnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// PreOrderBegin возвращает итератор для pre-order обхода с начала.
func (t *BST[T]) PreOrderBegin() interfaces.ForwardIterator[T] {
	return newPreOrderIterator(t.root)
}

// PreOrderEnd возвращает конечный итератор для pre-order обхода.
func (t *BST[T]) PreOrderEnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// PostOrderBegin возвращает итератор для post-order обхода с начала.
func (t *BST[T]) PostOrderBegin() interfaces.ForwardIterator[T] {
	return newPostOrderIterator(t.root)
}

// PostOrderEnd возвращает конечный итератор для post-order обхода.
func (t *BST[T]) PostOrderEnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}
