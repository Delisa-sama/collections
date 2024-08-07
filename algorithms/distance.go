package algorithms

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Distance вычисляет количество элементов в диапазоне [begin, end).
// Если итераторы поддерживают случайный доступ (RandomAccessIterator), вычисление производится за константное время.
// В противном случае используется линейный обход диапазона для подсчета количества элементов.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
//
// Возвращает:
// - количество элементов в диапазоне [begin, end) в виде значения типа uint.
func Distance[T any](begin, end interfaces.Iterator) uint {
	b, bIsRand := begin.(interfaces.RandomAccessIterator[T])
	e, eIsRand := end.(interfaces.RandomAccessIterator[T])
	if bIsRand && eIsRand {
		return e.Index() - b.Index()
	}

	var diff uint
	begin = copiable.Copy[interfaces.Iterator](begin)
	for !begin.Equals(end) {
		diff++
		begin.Next()
	}

	return diff
}
