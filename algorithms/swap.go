package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Swap производит обмен значениями двух указателей.
//
// Параметры:
// - a: указатель на первое значение.
// - b: указатель на второе значение.
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// SwapIter производит обмен значениями, на которые указывают два итератора.
//
// Параметры:
// - a: итератор, реализующий интерфейс PointerIterator, указывающий на первую переменную.
// - b: итератор, реализующий интерфейс PointerIterator, указывающий на вторую переменную.
func SwapIter[T any](a, b interfaces.PointerIterator[T]) {
	Swap(a.Ptr(), b.Ptr())
}

// SwapRanges производит обмен значениями между двумя диапазонами элементов.
//
// Параметры:
// - aBegin: итератор, указывающий на начало первого диапазона.
// - aEnd: итератор, указывающий на конец первого диапазона (невключительно).
// - bBegin: итератор, указывающий на начало второго диапазона.
//
// Функция последовательно обменивает значения, на которые указывают итераторы aBegin и bBegin,
// затем сдвигает оба итератора на следующий элемент и повторяет процесс до тех пор,
// пока aBegin не станет равным aEnd.
func SwapRanges[T any](
	aBegin interfaces.PointerIterator[T],
	aEnd interfaces.Iterator,
	bBegin interfaces.PointerIterator[T],
) {
	for !aBegin.Equals(aEnd) {
		SwapIter(aBegin, bBegin)

		aBegin.Next()
		bBegin.Next()
	}
}
