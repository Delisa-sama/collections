package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// AllOf проверяет, удовлетворяют ли все элементы в диапазоне [begin, end) предикату.
// Функция возвращает true, если все элементы удовлетворяют предикату, иначе false.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в поиск).
// - predicate: унарный предикат, который применяется к каждому элементу.
//
// Возвращает:
// - булево значение true, если все элементы удовлетворяют предикату,
// или false, если хотя бы один элемент не удовлетворяет.
func AllOf[T any](begin interfaces.ForwardIterator[T], end interfaces.Iterator, predicate unaryPredicate[T]) bool {
	_, found := FindIfNot(begin, end, predicate)
	return !found
}
