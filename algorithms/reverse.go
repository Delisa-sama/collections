package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Reverse разворачивает элементы в последовательности, определяемой итераторами
// begin и end, используя двунаправленные итераторы.
//
// Параметры:
// - begin: итератор на начало последовательности.
// - end: итератор на конец последовательности.
func Reverse[T any](begin, end interfaces.BidirectionalIterator[T]) {
	for !begin.Equals(end) {
		end.Prev()
		if begin.Equals(end) {
			break
		}
		SwapIter[T](begin, end)
		begin.Next()
	}
}

// ReverseCopy копирует элементы из диапазона [begin, end) в destBegin в обратном порядке.
//
// Параметры:
// - begin: итератор на начало исходной последовательности.
// - end: итератор на конец исходной последовательности.
// - destBegin: итератор на начало последовательности, куда будут копироваться элементы в обратном порядке.
//
// Возвращает итератор на конец целевой последовательности после копирования.
func ReverseCopy[T any](
	begin, end interfaces.BidirectionalIterator[T],
	destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T] {
	for !begin.Equals(end) {
		end.Prev()
		*destBegin.Ptr() = end.Value()
		destBegin.Next()
	}

	return destBegin
}
