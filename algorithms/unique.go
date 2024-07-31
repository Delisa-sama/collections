package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Unique удаляет дублирующиеся последовательные элементы в диапазоне [begin, end), оставляя только первые
// вхождения каждого элемента. Элементы считаются дубликатами, если они равны друг другу (==).
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
//
// Возвращает:
// - итератор, указывающий на конец уникальной последовательности элементов.
func Unique[T comparable](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
) interfaces.ForwardIterator[T] {
	return UniqueIf(begin, end, func(a T, b T) bool {
		return a == b
	})
}

// UniqueC удаляет дублирующиеся последовательные элементы в диапазоне [begin, end), оставляя только первые
// вхождения каждого элемента. Элементы считаются дубликатами, если они равны согласно пользовательскому компаратору.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
// - cmp: пользовательский компаратор, который определяет, равны ли два элемента.
//
// Возвращает:
// - итератор, указывающий на конец уникальной последовательности элементов.
func UniqueC[T any](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	cmp comparator.Comparator[T],
) interfaces.ForwardIterator[T] {
	return UniqueIf(begin, end, func(a T, b T) bool {
		return cmp(a, b) == 0
	})
}

// UniqueIf удаляет дублирующиеся последовательные элементы в диапазоне [begin, end), оставляя только первые
// вхождения каждого элемента. Элементы считаются дубликатами, если они удовлетворяют условию заданному предикатом.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
// - predicate: бинарный предикат, который определяет, являются ли два элемента равными.
//
// Возвращает:
// - итератор, указывающий на конец уникальной последовательности элементов.
func UniqueIf[T any](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	predicate binaryPredicate[T],
) interfaces.ForwardIterator[T] {
	if begin.Equals(end) {
		return begin
	}

	result := copiable.Copy[interfaces.ForwardIterator[T]](begin)

	begin.Next()
	for !begin.Equals(end) {
		if !predicate(result.Value(), begin.Value()) {
			result.Next()
			if !result.Equals(begin) {
				*result.Ptr() = *begin.Ptr()
			}
		}
		begin.Next()
	}

	result.Next()
	return result
}
