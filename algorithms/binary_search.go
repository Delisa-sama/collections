package algorithms

import (
	"cmp"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// BinarySearch выполняет двоичный поиск элемента в отсортированной последовательности.
// Использует интерфейс итератора для обхода последовательности и стандартный оператор сравнения.
// Параметры:
// - begin: итератор на начало последовательности.
// - end: итератор на конец последовательности.
// - value: значение, которое необходимо найти.
// Возвращает: true, если элемент найден, иначе false.
func BinarySearch[T cmp.Ordered](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	value T,
) bool {
	return BinarySearchC(begin, end, value, comparator.DefaultLess[T]())
}

// BinarySearchC выполняет двоичный поиск элемента в отсортированной последовательности
// с использованием настраиваемого компаратора.
// Параметры:
// - begin: итератор на начало последовательности.
// - end: итератор на конец последовательности.
// - value: значение, которое необходимо найти.
// - less: функция сравнения (компаратор), которая определяет порядок элементов.
// Возвращает: true, если элемент найден, иначе false.
func BinarySearchC[T any](
	begin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	value T,
	less comparator.Less[T],
) bool {
	first := LowerBoundC(begin, end, value, less)
	return !first.Equals(end) && !less(value, first.Value())
}
