package forwardlist

import (
	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// ForwardList представляет собой односвязный список.
type ForwardList[T any] struct {
	top  *node[T]
	end  *node[T]
	size uint
}

// NewForwardList создает новый ForwardList и заполняет его переданными элементами.
func NewForwardList[T any](items ...T) *ForwardList[T] {
	l := &ForwardList[T]{}
	for i := range items {
		l.PushBack(items[i])
	}
	return l
}

// node представляет собой узел односвязного списка.
type node[T any] struct {
	Value *T
	Next  *node[T]
}

// Size возвращает количество элементов в списке.
func (l *ForwardList[T]) Size() uint {
	return l.size
}

// IsEmpty проверяет что список пустой.
func (l *ForwardList[T]) IsEmpty() bool {
	return l.size == 0
}

// Begin возвращает итератор на первый элемент списка.
func (l *ForwardList[T]) Begin() interfaces.ForwardIterator[T] {
	return newIterator(l.top)
}

// End возвращает итератор на последний элемент списка.
func (l *ForwardList[T]) End() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// PushBack добавляет новый элемент в конец списка.
func (l *ForwardList[T]) PushBack(value T) {
	newNode := &node[T]{
		Value: &value,
		Next:  nil,
	}
	if l.top == nil {
		l.top = newNode
	}
	if l.end != nil {
		l.end.Next = newNode
	}
	l.end = newNode
	l.size++
}

// Back возвращает последний элемент списка.
func (l *ForwardList[T]) Back() T {
	return *l.end.Value
}

// PopBack удаляет последний элемент из списка.
func (l *ForwardList[T]) PopBack() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.top = nil
		l.end = nil
		l.size = 0
		return
	}

	current := l.top
	for current.Next != l.end {
		current = current.Next
	}
	current.Next = nil
	l.end = current
	l.size--
}

// Copy копирует список.
func (l *ForwardList[T]) Copy() copiable.Copiable {
	copyList := NewForwardList[T]()
	algorithms.ForEach(l.Begin(), l.End(), func(value T) {
		copyList.PushBack(value)
	})
	return copyList
}
