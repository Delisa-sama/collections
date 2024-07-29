package stack

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// container - это интерфейс, который описывает контейнер, поддерживающий операции с задней частью.
type container[T any] interface {
	interfaces.Container[T]
	// Back возвращает последний элемент в контейнере.
	Back() T
	// PushBack добавляет элемент в конец контейнера.
	PushBack(T)
	// PopBack удаляет последний элемент из контейнера.
	PopBack()
}

type containerConstructor[T any, C container[T]] func(...T) C

// Stack представляет собой стек, работающий на основе контейнера C.
type Stack[T any, C container[T]] struct {
	c C
}

// NewStack возвращает указатель на новый стек.
//
// Параметры:
// cc - функция, конструктор контейнера из произвольного количества элементов.
// items - элементы стека.
func NewStack[T any, C container[T]](cc containerConstructor[T, C], items ...T) *Stack[T, C] {
	return &Stack[T, C]{
		c: cc(items...),
	}
}

// Push добавляет элемент в стек.
func (s *Stack[T, C]) Push(v T) {
	s.c.PushBack(v)
}

// Pop удаляет последний элемент из стека.
func (s *Stack[T, C]) Pop() {
	s.c.PopBack()
}

// Top возвращает последний элемент стека.
func (s *Stack[T, C]) Top() T {
	return s.c.Back()
}

// Size возвращает количество элементов в стеке.
func (s *Stack[T, C]) Size() uint {
	return s.c.Size()
}

// IsEmpty проверяет, пуст ли стек.
func (s *Stack[T, C]) IsEmpty() bool {
	return s.c.IsEmpty()
}

// Copy возвращает копию стека.
func (s *Stack[T, C]) Copy() *Stack[T, C] {
	return &Stack[T, C]{
		c: copiable.Copy[C](s.c),
	}
}
