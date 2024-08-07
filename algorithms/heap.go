package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

const (
	heapMinSize               = 2
	heapChildIndexFactor      = 2
	heapLeftChildIndexOffset  = 1
	heapRightChildIndexOffset = 2
	heapParentIndexFactor     = 2
	heapMiddleIndexFactor     = 2
)

// MakeHeap преобразует диапазон [begin, end) в кучу, используя предоставленный компаратор cmp.
//
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
func MakeHeap[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) {
	size := Distance[T](begin, end)
	if size < heapMinSize {
		return
	}
	for i := int(size/heapMiddleIndexFactor - 1); i >= 0; i-- {
		heapify[T](begin, end, cmp, uint(i))
	}
}

// SortHeap выполняет сортировку кучи на месте, упорядочивая элементы в диапазоне [begin, end).
//
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
func SortHeap[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) {
	size := Distance[T](begin, end)
	for size > 1 {
		last := AdvanceCopy[T](begin, int(size-1))
		Swap[T](begin.Ptr(), last.Ptr())
		size--
		heapify[T](begin, AdvanceCopy[T](begin, int(size)), cmp, 0)
	}
}

// PopHeap удаляет максимальный элемент из кучи и перестраивает её.
//
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
func PopHeap[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) {
	last := AdvanceCopy[T](end, -1)
	Swap[T](begin.Ptr(), last.Ptr())
	heapify[T](begin, last, cmp, 0)
}

// PushHeap добавляет элемент в конец диапазона и перестраивает кучу, чтобы сохранить её свойства.
//
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
func PushHeap[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) {
	size := Distance[T](begin, end)
	if size < heapMinSize {
		return
	}

	idx := size - 1 // последний элемент, который нужно "всплыть"

	for idx > 0 {
		parentIdx := (idx - 1) / heapParentIndexFactor // родительский элемент

		currentVal, _ := begin.At(idx)
		parentVal, _ := begin.At(parentIdx)

		// Если новый элемент больше родителя, меняем их местами
		if comparator.Greater(*currentVal, *parentVal, cmp) {
			Swap[T](currentVal, parentVal)
			idx = parentIdx
		} else {
			break
		}
	}
}

// heapify перестраивает поддерево, корнем которого является элемент с индексом i,
// таким образом, чтобы оно соответствовало свойствам кучи.
//
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
// - i: индекс корня поддерева, который нужно перестроить.
func heapify[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
	i uint,
) {
	size := Distance[T](begin, end)
	if size == 0 || i >= size {
		return
	}

	leftIdx := heapChildIndexFactor*i + heapLeftChildIndexOffset
	rightIdx := heapChildIndexFactor*i + heapRightChildIndexOffset
	largestIdx := i

	if leftIdx < size {
		left, _ := begin.At(leftIdx)
		largest, _ := begin.At(largestIdx)
		if comparator.Greater(*left, *largest, cmp) {
			largestIdx = leftIdx
		}
	}

	if rightIdx < size {
		right, _ := begin.At(rightIdx)
		largest, _ := begin.At(largestIdx)
		if comparator.Greater(*right, *largest, cmp) {
			largestIdx = rightIdx
		}
	}

	if largestIdx != i {
		iVal, _ := begin.At(i)
		largest, _ := begin.At(largestIdx)
		Swap[T](iVal, largest)
		// После обмена элементов вызываем heapify рекурсивно для нового поддерева
		heapify[T](begin, end, cmp, largestIdx)
	}
}
