package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// NoneOf проверяет, удовлетворяет ли ни один элемент в диапазоне [begin, end) предикату.
// Функция возвращает true, если ни один элемент не удовлетворяет предикату, иначе false.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в проверку).
// - predicate: унарный предикат, который применяется к каждому элементу.
//
// Возвращает:
// - булево значение true, если ни один элемент не удовлетворяет предикату, или false, если хотя бы один элемент удовлетворяет.
func NoneOf[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool {
	_, found := FindIf(begin, end, predicate)
	return !found
}
