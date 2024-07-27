package list

import (
	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// List представляет собой двусвязный список.
type List[T any] struct {
	top  *node[T]
	end  *node[T]
	size uint
}

// NewList создает новый List и заполняет его переданными элементами.
func NewList[T any](items ...T) *List[T] {
	l := &List[T]{}
	for i := range items {
		l.PushBack(items[i])
	}
	return l
}

// node представляет собой узел двусвязного списка.
type node[T any] struct {
	Value *T
	Next  *node[T]
	Prev  *node[T]
}

// Size возвращает количество элементов в списке.
func (l *List[T]) Size() uint {
	return l.size
}

// IsEmpty проверяет что список пустой.
func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

// Begin возвращает итератор на первый элемент списка.
func (l *List[T]) Begin() interfaces.BidirectionalIterator[T] {
	return newIterator(l.top)
}

// End возвращает итератор на последний элемент списка.
func (l *List[T]) End() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// RBegin возвращает итератор на последний элемент списка.
func (l *List[T]) RBegin() interfaces.BidirectionalIterator[T] {
	return iterators.NewReverseIterator[T](newIterator(l.end))
}

// REnd возвращает итератор на конец списка.
func (l *List[T]) REnd() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// PushBack добавляет новый элемент в конец списка.
func (l *List[T]) PushBack(value T) {
	newNode := &node[T]{
		Value: &value,
		Next:  nil,
		Prev:  nil,
	}
	if l.top == nil {
		l.top = newNode
	}
	if l.end != nil {
		newNode.Prev = l.end
		l.end.Next = newNode
	}
	l.end = newNode
	l.size++
}

// Back возвращает последний элемент списка.
func (l *List[T]) Back() T {
	return *l.end.Value
}

// PopBack удаляет последний элемент из списка.
func (l *List[T]) PopBack() {
	if l.IsEmpty() {
		return
	}
	l.end = l.end.Prev
	l.size--
}

// Copy копирует список.
func (l *List[T]) Copy() interfaces.Container[T] {
	copyList := NewList[T]()
	algorithms.ForEach(l.Begin(), l.End(), func(value T) {
		copyList.PushBack(value)
	})
	return copyList
}
