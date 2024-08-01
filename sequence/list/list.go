package list

import (
	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// List представляет собой двусвязный список.
type List[T any] struct {
	head *node[T]
	tail *node[T]
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
	return newIterator(l.head)
}

// End возвращает итератор на последний элемент списка.
func (l *List[T]) End() interfaces.Iterator {
	return iterators.NewEndIterator()
}

// RBegin возвращает итератор на последний элемент списка.
func (l *List[T]) RBegin() interfaces.BidirectionalIterator[T] {
	return iterators.NewReverseIterator[T](newIterator(l.tail))
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
	if l.head == nil {
		l.head = newNode
	}
	if l.tail != nil {
		newNode.Prev = l.tail
		l.tail.Next = newNode
	}
	l.tail = newNode
	l.size++
}

// Back возвращает последний элемент списка.
func (l *List[T]) Back() T {
	return *l.tail.Value
}

// PopBack удаляет последний элемент из списка.
func (l *List[T]) PopBack() {
	if l.IsEmpty() {
		return
	}
	l.tail = l.tail.Prev
	l.size--
}

// Copy копирует список.
func (l *List[T]) Copy() copiable.Copiable {
	copyList := NewList[T]()
	algorithms.ForEach(l.Begin(), l.End(), func(value T) {
		copyList.PushBack(value)
	})
	return copyList
}

// Erase удаляет элементы в диапазоне [begin, end) из списка.
func (l *List[T]) Erase(begin, end interfaces.Iterator) {
	if begin.Equals(end) {
		return
	}

	b, bOk := begin.(*iterator[T])
	e, eOk := end.(*iterator[T])
	if !bOk || !eOk {
		panic("unknown iterator type")
	}

	if b.current == l.head {
		// Удаление с начала списка.
		l.head = e.current
		if l.head != nil {
			l.head.Prev = nil
		}
	} else {
		b.current.Prev.Next = e.current
	}

	if e.current == nil {
		// Удаление до конца списка.
		l.tail = b.current.Prev
	} else {
		e.current.Prev = b.current.Prev
	}

	// Корректировка размера списка.
	removedSize := uint(0)
	for !b.Equals(end) {
		removedSize++
		b.Next()
	}
	l.size -= removedSize
}
