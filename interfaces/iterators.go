package interfaces

import (
	"github.com/Delisa-sama/collections/copiable"
)

// Iterator - это интерфейс для перебора коллекции.
type Iterator interface {
	copiable.Copiable
	// Equals проверяет, равны ли два итератора.
	Equals(iterator Iterator) bool
	// HasNext проверяет, есть ли еще элементы для перебора.
	HasNext() bool
	// Next переходит к следующему элементу.
	Next()
}

// ValueIterator - это интерфейс для итератора, который возвращает значение.
type ValueIterator[T any] interface {
	Iterator
	// Value возвращает текущее значение итератора.
	Value() T
}

// PointerIterator - это интерфейс для итератора, который возвращает указатель на значение.
type PointerIterator[T any] interface {
	Iterator
	// Ptr возвращает указатель на текущее значение итератора.
	Ptr() *T
}

// ForwardIterator - это интерфейс для прямого итератора, который объединяет ValueIterator и PointerIterator.
type ForwardIterator[T any] interface {
	ValueIterator[T]
	PointerIterator[T]
}

// BidirectionalIterator - это интерфейс для двунаправленного итератора,
// который объединяет ForwardIterator и добавляет методы для обратного перебора.
type BidirectionalIterator[T any] interface {
	ForwardIterator[T]
	// HasPrev проверяет, есть ли предыдущие элементы для перебора.
	HasPrev() bool
	// Prev переходит к предыдущему элементу.
	Prev()
}

// RandomAccessIterator - это интерфейс для итератора с произвольным доступом,
// который объединяет BidirectionalIterator и добавляет метод для доступа по индексу.
type RandomAccessIterator[T any] interface {
	BidirectionalIterator[T]
	// At проверяет, доступен ли элемент по заданному индексу и возвращает его.
	At(index uint) (*T, bool)
	// Shift сдвигает итератор, если offset > 0, то вперед, если < 0, то назад.
	Shift(offset int)
	// Index возвращает текущий индекс итератора.
	Index() uint
}
