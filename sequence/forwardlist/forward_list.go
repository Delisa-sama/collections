package forwardlist

import (
	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/iterators"
)

// ForwardList представляет собой односвязный список.
type ForwardList[T any] struct {
	head *node[T]
	tail *node[T]
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
	return newIterator(l.head)
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
	if l.head == nil {
		l.head = newNode
	}
	if l.tail != nil {
		l.tail.Next = newNode
	}
	l.tail = newNode
	l.size++
}

// Back возвращает последний элемент списка.
func (l *ForwardList[T]) Back() T {
	return *l.tail.Value
}

// PopBack удаляет последний элемент из списка.
func (l *ForwardList[T]) PopBack() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}

	current := l.head
	for current.Next != l.tail {
		current = current.Next
	}
	current.Next = nil
	l.tail = current
	l.size--
}

// PushFront добавляет новый элемент в начало списка.
func (l *ForwardList[T]) PushFront(value T) {
	newNode := &node[T]{
		Value: &value,
		Next:  nil,
	}
	if l.tail == nil {
		l.tail = newNode
	}
	l.head = newNode
	l.size++
}

// Front возвращает первый элемент списка.
func (l *ForwardList[T]) Front() T {
	return *l.head.Value
}

// PopFront удаляет первый элемент из списка.
func (l *ForwardList[T]) PopFront() {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return
	}

	l.head = l.head.Next
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

// Erase удаляет элементы в диапазоне [begin, end) из списка.
func (l *ForwardList[T]) Erase(begin, end interfaces.Iterator) {
	if begin.Equals(end) {
		return
	}

	i := newIterator(l.head)
	var prev *node[T]

	// Найдем узел, соответствующий началу диапазона удаления (begin).
	for !i.Equals(begin) {
		prev = i.current
		i.Next()
	}

	// Обновляем ссылки, пропуская элементы до tail.
	for !i.Equals(end) && i.current != nil {
		i.Next()
		l.size--
	}

	if prev != nil {
		// Соединяем предыдущий узел с узлом после tail.
		prev.Next = i.current
	} else {
		// Если begin был первым элементом списка, обновляем head.
		l.head = i.current
	}

	// Если tail равен конечному итератору, обновляем tail списка.
	if i.current == nil {
		l.tail = prev
	}
}
