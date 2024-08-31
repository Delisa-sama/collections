package queue

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// container - это интерфейс, описывающий контейнер, поддерживающий операции
// над элементами в конце и начале контейнера.
// Контейнер должен реализовать методы получения элементов с начала и конца,
// добавления элементов в конец и удаления с начала.
type container[T any] interface {
	interfaces.Container[T]
	// Back возвращает последний элемент в контейнере.
	Back() T
	// Front возвращает первый элемент в контейнере.
	Front() T
	// PushBack добавляет элемент в конец контейнера.
	PushBack(T)
	// PopFront удаляет первый элемент из контейнера.
	PopFront()
}

// containerConstructor представляет собой функцию-конструктор,
// которая создает экземпляр контейнера с элементами.
// Параметры:
// - items: начальные элементы, которые будут добавлены в контейнер.
// Возвращает: созданный контейнер C, содержащий элементы.
type containerConstructor[T any, C container[T]] func(...T) C

// Queue представляет собой общую очередь (FIFO - First In, First Out),
// которая использует контейнер C для хранения элементов.
// Тип T - это тип элементов, хранящихся в очереди.
// Тип C - это тип контейнера, реализующий интерфейс container, используемый для управления элементами.
type Queue[T any, C container[T]] struct {
	c C // Контейнер, реализующий функциональность очереди.
}

// NewQueue создает новую очередь с использованием указанного конструктора контейнера и начальных элементов.
// Параметры:
// - cc: функция-конструктор, которая создает контейнер.
// - items: начальные элементы, которые будут добавлены в очередь.
// Возвращает: указатель на созданную очередь.
func NewQueue[T any, C container[T]](cc containerConstructor[T, C], items ...T) *Queue[T, C] {
	return &Queue[T, C]{
		c: cc(items...),
	}
}

// Front возвращает первый элемент в очереди (без его удаления).
// Возвращает: первый элемент из очереди.
func (q *Queue[T, C]) Front() T {
	return q.c.Front()
}

// Back возвращает последний элемент в очереди (без его удаления).
// Возвращает: последний элемент из очереди.
func (q *Queue[T, C]) Back() T {
	return q.c.Back()
}

// PushBack добавляет элемент в конец очереди.
// Параметры:
// - value: элемент, который будет добавлен в конец очереди.
func (q *Queue[T, C]) PushBack(value T) {
	q.c.PushBack(value)
}

// PopFront удаляет первый элемент из очереди.
func (q *Queue[T, C]) PopFront() {
	q.c.PopFront()
}

// Size возвращает количество элементов в очереди.
// Возвращает: количество элементов в очереди.
func (q *Queue[T, C]) Size() uint {
	return q.c.Size()
}

// IsEmpty проверяет, пуста ли очередь.
// Возвращает: true, если очередь пуста, иначе false.
func (q *Queue[T, C]) IsEmpty() bool {
	return q.c.IsEmpty()
}

// Copy создает и возвращает копию текущей очереди.
// Возвращает: копию очереди с теми же элементами.
func (q *Queue[T, C]) Copy() copiable.Copiable {
	return &Queue[T, C]{
		c: copiable.Copy[C](q.c),
	}
}