package algorithms

import (
	"cmp"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// LowerBound находит первый элемент, который не меньше чем значение value
// в отсортированном диапазоне [begin, end).
//
// Возвращает итератор на первый элемент, не меньший чем value.
// Если все элементы меньше value, возвращает end.
//
// T - тип элементов, поддерживающий сравнение (cmp.Ordered).
func LowerBound[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) interfaces.ValueIterator[T] {
	return LowerBoundC(begin, end, value, comparator.DefaultLess[T]())
}

const binarySearchDiv = 2

// LowerBoundC находит первый элемент, который не меньше чем значение value
// в отсортированном диапазоне [begin, end), используя пользовательский компаратор.
//
// cmp - компаратор для сравнения элементов.
//
// Возвращает итератор на первый элемент, не меньший чем value.
// Если все элементы меньше value, возвращает end.
func LowerBoundC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	less comparator.Less[T],
) interfaces.ValueIterator[T] {
	var it interfaces.ValueIterator[T]
	var step uint
	count := Distance[T](
		copiable.Copy[interfaces.ValueIterator[T]](begin),
		copiable.Copy[interfaces.Iterator](end),
	)

	for count > 0 {
		it = copiable.Copy[interfaces.ValueIterator[T]](begin)
		step = count / binarySearchDiv
		Advance[T](it, int(step))

		if less(it.Value(), value) {
			it.Next()
			begin = copiable.Copy[interfaces.ValueIterator[T]](it)
			count -= step + 1
		} else {
			count = step
		}
	}

	return begin
}

// UpperBound находит первый элемент, который больше чем значение value
// в отсортированном диапазоне [begin, end).
//
// Возвращает итератор на первый элемент, больший чем value.
// Если все элементы не больше value, возвращает end.
//
// T - тип элементов, поддерживающий сравнение (cmp.Ordered).
func UpperBound[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) interfaces.ValueIterator[T] {
	return UpperBoundC(begin, end, value, comparator.DefaultLess[T]())
}

// UpperBoundC находит первый элемент, который больше чем значение value
// в отсортированном диапазоне [begin, end), используя пользовательский компаратор.
//
// cmp - компаратор для сравнения элементов.
//
// Возвращает итератор на первый элемент, больший чем value.
// Если все элементы не больше value, возвращает end.
func UpperBoundC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	less comparator.Less[T],
) interfaces.ValueIterator[T] {
	var it interfaces.ValueIterator[T]
	var step uint
	count := Distance[T](
		copiable.Copy[interfaces.ValueIterator[T]](begin),
		copiable.Copy[interfaces.Iterator](end),
	)

	for count > 0 {
		it = copiable.Copy[interfaces.ValueIterator[T]](begin)
		step = count / binarySearchDiv
		Advance[T](it, int(step))

		if !less(value, it.Value()) {
			it.Next()
			begin = it
			count -= step + 1
		} else {
			count = step
		}
	}

	return begin
}
