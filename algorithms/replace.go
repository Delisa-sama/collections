package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Replace заменяет все вхождения значения oldValue на newValue в диапазоне от begin до end.
//
// Параметры:
// - begin: итератор на начало последовательности, где производится замена.
// - end: итератор на конец последовательности.
// - oldValue: значение, которое нужно заменить.
// - newValue: значение, на которое нужно заменить oldValue.
func Replace[T comparable](
	begin interfaces.PointerIterator[T],
	end interfaces.Iterator,
	oldValue T,
	newValue T,
) {
	for ; !begin.Equals(end); begin.Next() {
		if *begin.Ptr() == oldValue {
			*begin.Ptr() = newValue
		}
	}
}

// ReplaceIf заменяет все значения, которые удовлетворяют предикату, на newValue
// в диапазоне от begin до end.
//
// Параметры:
// - begin: итератор на начало последовательности, где производится замена.
// - end: итератор на конец последовательности.
// - predicate: предикат, определяющий, какие элементы заменять.
// - newValue: значение, на которое нужно заменить элементы, удовлетворяющие предикату.
func ReplaceIf[T any](
	begin interfaces.PointerIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
	newValue T,
) {
	for ; !begin.Equals(end); begin.Next() {
		if predicate(*begin.Ptr()) {
			*begin.Ptr() = newValue
		}
	}
}

// ReplaceCopy копирует элементы из диапазона [begin, end) в destBegin,
// заменяя все вхождения oldValue на newValue.
//
// Параметры:
// - begin: итератор на начало исходной последовательности.
// - end: итератор на конец исходной последовательности.
// - destBegin: итератор на начало последовательности, куда будут копироваться элементы.
// - oldValue: значение, которое нужно заменить при копировании.
// - newValue: значение, на которое нужно заменить oldValue при копировании.
func ReplaceCopy[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
	oldValue T,
	newValue T,
) {
	for ; !begin.Equals(end); begin.Next() {
		if begin.Value() == oldValue {
			*destBegin.Ptr() = newValue
		} else {
			*destBegin.Ptr() = begin.Value()
		}
		destBegin.Next()
	}
}

// ReplaceCopyIf копирует элементы из диапазона [begin, end) в destBegin,
// заменяя все значения, которые удовлетворяют предикату, на newValue.
//
// Параметры:
// - begin: итератор на начало исходной последовательности.
// - end: итератор на конец исходной последовательности.
// - destBegin: итератор на начало последовательности, куда будут копироваться элементы.
// - predicate: предикат, определяющий, какие элементы заменять при копировании.
// - newValue: значение, на которое нужно заменить элементы, удовлетворяющие предикату при копировании.
func ReplaceCopyIf[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
	predicate unaryPredicate[T],
	newValue T,
) {
	for ; !begin.Equals(end); begin.Next() {
		if predicate(begin.Value()) {
			*destBegin.Ptr() = newValue
		} else {
			*destBegin.Ptr() = begin.Value()
		}
		destBegin.Next()
	}
}
