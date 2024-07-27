package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Copy копирует элементы из диапазона [begin, end) в диапазон, начинающийся с destBegin.
// Функция возвращает итератор на конец диапазона назначения.
//
// Параметры:
// - begin: итератор, указывающий на начало исходного диапазона.
// - end: итератор, указывающий на конец исходного диапазона (не включается в копирование).
// - destBegin: итератор, указывающий на начало диапазона назначения.
//
// Возвращает:
// - итератор на конец диапазона назначения.
func Copy[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T] {
	for !begin.Equals(end) {
		value := begin.Value()
		*destBegin.Ptr() = value

		begin.Next()
		destBegin.Next()
	}

	return destBegin
}

// CopyIf копирует элементы из диапазона [begin, end), которые удовлетворяют предикату, в диапазон, начинающийся с destBegin.
// Функция возвращает итератор на конец диапазона назначения.
//
// Параметры:
// - begin: итератор, указывающий на начало исходного диапазона.
// - end: итератор, указывающий на конец исходного диапазона (не включается в копирование).
// - destBegin: итератор, указывающий на начало диапазона назначения.
// - predicate: унарный предикат, который применяется к каждому элементу.
//
// Возвращает:
// - итератор на конец диапазона назначения.
func CopyIf[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
	predicate unaryPredicate[T],
) interfaces.PointerIterator[T] {
	for !begin.Equals(end) {
		value := begin.Value()
		if predicate(value) {
			*destBegin.Ptr() = value
			destBegin.Next()
		}

		begin.Next()
	}

	return destBegin
}
