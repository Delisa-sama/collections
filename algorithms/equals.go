package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// EqualsC проверяет, равны ли все элементы двух диапазонов [a, b), используя пользовательский компаратор.
// Функция возвращает true, если все элементы равны, иначе false.
//
// Параметры:
// - a: итератор, указывающий на начало первого диапазона.
// - b: итератор, указывающий на начало второго диапазона.
// - cmp: пользовательский компаратор для сравнения элементов.
//
// Возвращает:
// - булево значение true, если все элементы равны, или false, если хотя бы один элемент не равен.
func EqualsC[T any](a, b interfaces.ForwardIterator[T], cmp comparator.Comparator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// Equals проверяет, равны ли все элементы двух диапазонов [a, b), используя оператор сравнения ==.
// Функция возвращает true, если все элементы равны, иначе false.
//
// Параметры:
// - a: итератор, указывающий на начало первого диапазона.
// - b: итератор, указывающий на начало второго диапазона.
//
// Возвращает:
// - булево значение true, если все элементы равны, или false, если хотя бы один элемент не равен.
func Equals[T comparable](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if a.Value() != b.Value() {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// EqualsRangesC проверяет, равны ли все элементы двух диапазонов [aBegin, aEnd) и [bBegin, bEnd),
// используя пользовательский компаратор.
// Функция возвращает true, если все элементы равны, иначе false.
//
// Параметры:
// - aBegin: итератор, указывающий на начало первого диапазона.
// - aEnd: итератор, указывающий на конец первого диапазона (не включается в проверку).
// - bBegin: итератор, указывающий на начало второго диапазона.
// - bEnd: итератор, указывающий на конец второго диапазона (не включается в проверку).
// - cmp: пользовательский компаратор для сравнения элементов.
//
// Возвращает:
// - булево значение true, если все элементы равны, или false, если хотя бы один элемент не равен.
func EqualsRangesC[T any](
	aBegin interfaces.ValueIterator[T], aEnd interfaces.Iterator,
	bBegin interfaces.ValueIterator[T], bEnd interfaces.Iterator,
	cmp comparator.Comparator[T],
) bool {
	a := aBegin
	b := bBegin
	for !a.Equals(aEnd) || !b.Equals(bEnd) {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// EqualsRanges проверяет, равны ли все элементы двух диапазонов [aBegin, aEnd) и [bBegin, bEnd),
// используя оператор сравнения ==.
// Функция возвращает true, если все элементы равны, иначе false.
//
// Параметры:
// - aBegin: итератор, указывающий на начало первого диапазона.
// - aEnd: итератор, указывающий на конец первого диапазона (не включается в проверку).
// - bBegin: итератор, указывающий на начало второго диапазона.
// - bEnd: итератор, указывающий на конец второго диапазона (не включается в проверку).
//
// Возвращает:
// - булево значение true, если все элементы равны, или false, если хотя бы один элемент не равен.
func EqualsRanges[T comparable](
	aBegin interfaces.ValueIterator[T], aEnd interfaces.Iterator,
	bBegin interfaces.ValueIterator[T], bEnd interfaces.Iterator,
) bool {
	a := aBegin
	b := bBegin
	for !a.Equals(aEnd) || !b.Equals(bEnd) {
		if a.Value() != b.Value() {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// xor возвращает true, если один из аргументов true, но не оба.
// Функция используется для проверки, что оба диапазона имеют одинаковую длину.
//
// Параметры:
// - x: первый булевый аргумент.
// - y: второй булевый аргумент.
//
// Возвращает:
// - булево значение true, если только один из аргументов равен true, или false, если оба аргумента равны.
func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
