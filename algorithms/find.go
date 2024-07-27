package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// FindC выполняет поиск элемента в диапазоне [begin, end) с использованием
// пользовательского компаратора. Функция возвращает итератор на найденный элемент и
// булево значение, указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона поиска.
// - end: итератор, указывающий на конец диапазона поиска (не включается в поиск).
// - value: значение, которое необходимо найти.
// - cmp: пользовательский компаратор для сравнения элементов.
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func FindC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
	cmp comparator.Comparator[T],
) (interfaces.ValueIterator[T], bool) {
	return FindIf(begin, end, func(v T) bool {
		return cmp(v, value) == 0
	})
}

// Find выполняет поиск элемента в диапазоне [begin, end) с использованием
// оператора сравнения ==. Функция возвращает итератор на найденный элемент и
// булево значение, указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона поиска.
// - end: итератор, указывающий на конец диапазона поиска (не включается в поиск).
// - value: значение, которое необходимо найти.
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func Find[T comparable](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	value T,
) (interfaces.ValueIterator[T], bool) {
	return FindIf(begin, end, func(v T) bool {
		return v == value
	})
}

// FindIf выполняет поиск элемента в диапазоне [begin, end), для которого предикат
// возвращает true. Функция возвращает итератор на найденный элемент и булево значение,
// указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона поиска.
// - end: итератор, указывающий на конец диапазона поиска (не включается в поиск).
// - predicate: унарный предикат, который применяется к каждому элементу диапазона.
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func FindIf[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) (interfaces.ValueIterator[T], bool) {
	for ; !begin.Equals(end); begin.Next() {
		if predicate(begin.Value()) {
			return begin, true
		}
	}
	return nil, false
}

// FindIfNot выполняет поиск элемента в диапазоне [begin, end), для которого предикат
// возвращает false. Функция возвращает итератор на найденный элемент и булево значение,
// указывающее на успех поиска.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона поиска.
// - end: итератор, указывающий на конец диапазона поиска (не включается в поиск).
// - predicate: унарный предикат, который применяется к каждому элементу диапазона.
//
// Возвращает:
// - итератор на найденный элемент, если элемент найден, или nil, если элемент не найден.
// - булево значение true, если элемент найден, или false, если элемент не найден.
func FindIfNot[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) (interfaces.ValueIterator[T], bool) {
	for ; !begin.Equals(end); begin.Next() {
		if !predicate(begin.Value()) {
			return begin, true
		}
	}
	return nil, false
}
