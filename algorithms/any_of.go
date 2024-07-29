package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// AnyOf проверяет, удовлетворяет ли хотя бы один элемент в диапазоне [begin, end) предикату.
// Функция возвращает true, если хотя бы один элемент удовлетворяет предикату, иначе false.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в поиск).
// - predicate: унарный предикат, который применяется к каждому элементу.
//
// Возвращает:
// - булево значение true, если хотя бы один элемент удовлетворяет предикату,
// или false, если ни один элемент не удовлетворяет.
func AnyOf[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool {
	_, found := FindIf(begin, end, predicate)
	return found
}
