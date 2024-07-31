package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Remove удаляет все элементы, равные заданному значению `value`, из диапазона [begin, end).
// Элементы, которые не равны `value`, сохраняются в начале диапазона, а оставшиеся элементы не изменяются.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
// - value: значение, которое необходимо удалить из диапазона.
//
// Возвращает:
// - итератор, указывающий на первый элемент, следующий за последним сохранённым элементом.
func Remove[T comparable](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	value T,
) interfaces.ForwardIterator[T] {
	return RemoveIf(begin, end, func(v T) bool {
		return v == value
	})
}

// RemoveC удаляет все элементы, равные заданному значению `value`, из диапазона [begin, end),
// используя пользовательский компаратор для сравнения элементов.
// Элементы, которые не равны `value`, сохраняются в начале диапазона, а оставшиеся элементы не изменяются.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
// - value: значение, которое необходимо удалить из диапазона.
// - cmp: компаратор, используемый для сравнения элементов.
//
// Возвращает:
// - итератор, указывающий на первый элемент, следующий за последним сохранённым элементом.
func RemoveC[T any](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	value T,
	cmp comparator.Comparator[T],
) interfaces.ForwardIterator[T] {
	return RemoveIf(begin, end, func(v T) bool {
		return cmp(v, value) == 0
	})
}

// RemoveIf удаляет все элементы из диапазона [begin, end), которые удовлетворяют заданному предикату `predicate`.
// Элементы, которые не удовлетворяют предикату, сохраняются в начале диапазона, а оставшиеся элементы не изменяются.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
// - predicate: предикат, который проверяет, нужно ли удалить элемент.
//
// Возвращает:
// - итератор, указывающий на первый элемент, следующий за последним сохранённым элементом.
func RemoveIf[T any](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	predicate unaryPredicate[T],
) interfaces.ForwardIterator[T] {
	for ; !begin.Equals(end); begin.Next() {
		if predicate(begin.Value()) {
			break
		}
	}

	beginV := copiable.Copy[interfaces.ForwardIterator[T]](begin)
	if !begin.Equals(end) {
		i := copiable.Copy[interfaces.ForwardIterator[T]](beginV)
		i.Next()
		for ; !i.Equals(end); i.Next() {
			if !predicate(i.Value()) {
				*beginV.Ptr() = *i.Ptr()
				beginV.Next()
			}
		}
	}

	return beginV
}
