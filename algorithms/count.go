package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// CountC подсчитывает количество элементов в диапазоне [begin, end), равных заданному значению, используя пользовательский компаратор.
// Функция возвращает количество найденных элементов.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в подсчет).
// - value: значение, которое необходимо подсчитать.
// - cmp: пользовательский компаратор для сравнения элементов.
//
// Возвращает:
// - количество найденных элементов.
func CountC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	cmp comparator.Comparator[T],
) uint {
	return CountIf(begin, end, func(v T) bool {
		return cmp(v, value) == 0
	})
}

// Count подсчитывает количество элементов в диапазоне [begin, end), равных заданному значению, используя оператор сравнения ==.
// Функция возвращает количество найденных элементов.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в подсчет).
// - value: значение, которое необходимо подсчитать.
//
// Возвращает:
// - количество найденных элементов.
func Count[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) uint {
	return CountIf(begin, end, func(v T) bool {
		return v == value
	})
}

// CountIf подсчитывает количество элементов в диапазоне [begin, end), удовлетворяющих предикату.
// Функция возвращает количество найденных элементов.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в подсчет).
// - predicate: унарный предикат, который применяется к каждому элементу.
//
// Возвращает:
// - количество найденных элементов.
func CountIf[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) uint {
	var count uint
	for ; !begin.Equals(end); begin.Next() {
		if predicate(begin.Value()) {
			count++
		}
	}
	return count
}
