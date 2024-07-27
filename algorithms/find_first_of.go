package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// FindFirstOfC ищет первый элемент из диапазона [begin, end), который также содержится в диапазоне [sBegin, sEnd), используя пользовательский компаратор.
// Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало первого диапазона.
// - end: итератор, указывающий на конец первого диапазона (не включается в поиск).
// - sBegin: итератор, указывающий на начало второго диапазона.
// - sEnd: итератор, указывающий на конец второго диапазона (не включается в поиск).
// - cmp: пользовательский компаратор для сравнения элементов.
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func FindFirstOfC[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
	cmp comparator.Comparator[T],
) (interfaces.ValueIterator[T], bool) {
	return FindFirstOfIf(
		begin, end,
		sBegin, sEnd,
		func(a, b T) bool {
			return cmp(a, b) == 0
		},
	)
}

// FindFirstOf ищет первый элемент из диапазона [begin, end), который также содержится в диапазоне [sBegin, sEnd), используя оператор сравнения ==.
// Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало первого диапазона.
// - end: итератор, указывающий на конец первого диапазона (не включается в поиск).
// - sBegin: итератор, указывающий на начало второго диапазона.
// - sEnd: итератор, указывающий на конец второго диапазона (не включается в поиск).
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func FindFirstOf[T comparable](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
) (interfaces.ValueIterator[T], bool) {
	return FindFirstOfIf(
		begin, end,
		sBegin, sEnd,
		func(a, b T) bool {
			return a == b
		},
	)
}

// FindFirstOfIf ищет первый элемент из диапазона [begin, end), который также содержится в диапазоне [sBegin, sEnd), для которого предикат возвращает true.
// Функция возвращает итератор на найденный элемент и булево значение, указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало первого диапазона.
// - end: итератор, указывающий на конец первого диапазона (не включается в поиск).
// - sBegin: итератор, указывающий на начало второго диапазона.
// - sEnd: итератор, указывающий на конец второго диапазона (не включается в поиск).
// - predicate: бинарный предикат, который применяется к элементам из двух диапазонов.
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func FindFirstOfIf[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	sBegin interfaces.ValueIterator[T], sEnd interfaces.Iterator,
	predicate binaryPredicate[T],
) (interfaces.ValueIterator[T], bool) {
	for ; !begin.Equals(end); begin.Next() {
		for ; !sBegin.Equals(sEnd); sBegin.Next() {
			if predicate(begin.Value(), sBegin.Value()) {
				return begin, true
			}
		}
	}
	return nil, false
}
