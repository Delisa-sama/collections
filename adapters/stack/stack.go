package stack

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Container - это интерфейс, который описывает контейнер, поддерживающий операции с задней частью.
type Container[T any] interface {
	interfaces.Container[T]
	// Back возвращает последний элемент в контейнере.
	Back() T
	// PushBack добавляет элемент в конец контейнера.
	PushBack(T)
	// PopBack удаляет последний элемент из контейнера.
	PopBack()
}

// Stack представляет собой стек, работающий на основе контейнера C.
type Stack[T any, C Container[T]] struct {
	c C
}

func NewStack[T any, C Container[T]](cc func(...T) C, items ...T) *Stack[T, C] {
	return &Stack[T, C]{
		c: cc(items...),
	}
}

// Push добавляет элемент в стек.
func (s *Stack[T, Container]) Push(v T) {
	s.c.PushBack(v)
}

// Pop удаляет последний элемент из стека.
func (s *Stack[T, Container]) Pop() {
	s.c.PopBack()
}

// Top возвращает последний элемент стека.
func (s *Stack[T, Container]) Top() T {
	return s.c.Back()
}

// Size возвращает количество элементов в стеке.
func (s *Stack[T, Container]) Size() uint {
	return s.c.Size()
}

// IsEmpty проверяет, пуст ли стек.
func (s *Stack[T, Container]) IsEmpty() bool {
	return s.c.IsEmpty()
}
