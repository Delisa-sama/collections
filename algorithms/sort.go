package algorithms

import (
	"cmp"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Sort выполняет сортировку диапазона [begin, end) с использованием алгоритма pdqsort.
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
func Sort[T cmp.Ordered](
	begin interfaces.RandomAccessIterator[T],
	end interfaces.RandomAccessIterator[T],
) {
	pdqsort(begin, end, comparator.DefaultLess[T]())
}

// SortC выполняет сортировку диапазона [begin, end) с использованием алгоритма pdqsort и переданного компаратора.
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
func SortC[T any](
	begin interfaces.RandomAccessIterator[T],
	end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) {
	pdqsort(begin, end, cmp)
}

// pdqsort выполняет сортировку диапазона [begin, end) с использованием гибридного алгоритма pdqsort.
// Этот алгоритм сочетает в себе quicksort, insertion sort и heapsort для достижения высокой производительности.
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
func pdqsort[T any](
	begin interfaces.RandomAccessIterator[T],
	end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) {
	if begin.Equals(end) {
		return
	}

	size := Distance[T](
		copiable.Copy[interfaces.RandomAccessIterator[T]](begin),
		copiable.Copy[interfaces.RandomAccessIterator[T]](end),
	)
	pdqsortLoop(begin, end, cmp, log2(size), true)
}

const (
	// insertionSortThreshold — Порог размера массива, при котором используется сортировка вставками.
	insertionSortThreshold = 24

	// nintherThreshold — Порог размера массива, при котором используется псевдомедиана
	// из девяти для выбора опорного элемента.
	nintherThreshold = 128

	// partialInsertionSortLimit — Максимальное количество перемещений элементов при частичной сортировке вставками.
	partialInsertionSortLimit = 8
)

// pdqsortLoop — основной цикл pdqsort, который управляет процессом сортировки
// и переключает стратегии в зависимости от ситуации.
// Параметры:
// - begin: итератор на начало диапазона.
// - end: итератор на конец диапазона (не включая).
// - cmp: компаратор, определяющий порядок элементов.
// - badAllowed: количество допустимых "плохих" разбиений, после которого переключается на heapsort.
// - leftMost: флаг, указывающий на то, является ли текущий диапазон самым левым в исходном массиве.
//
//nolint:mnd,cyclop // алгоритм реализован с оглядкой на https://github.com/orlp/pdqsort
func pdqsortLoop[T any](
	begin interfaces.RandomAccessIterator[T],
	end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
	badAllowed uint,
	leftMost bool,
) {
	for {
		size := Distance[T](
			copiable.Copy[interfaces.RandomAccessIterator[T]](begin),
			copiable.Copy[interfaces.RandomAccessIterator[T]](end),
		)

		// Для небольших массивов используется сортировка вставками.
		if size < insertionSortThreshold {
			insertionSort(begin, end, cmp)
			return
		}

		// Выбор опорного элемента как медианы трёх или псевдомедианы из девяти элементов.
		s2 := size / 2
		if size > nintherThreshold {
			sort3(
				begin,
				AdvanceCopy[T](begin, int(s2)),
				AdvanceCopy[T](end, -1),
				cmp,
			)
			sort3(
				AdvanceCopy[T](begin, 1),
				AdvanceCopy[T](begin, int(s2-1)),
				AdvanceCopy[T](end, -2),
				cmp,
			)
			sort3(
				AdvanceCopy[T](begin, 2),
				AdvanceCopy[T](begin, int(s2+1)),
				AdvanceCopy[T](end, -3),
				cmp,
			)
			sort3(
				AdvanceCopy[T](begin, int(s2-1)),
				AdvanceCopy[T](begin, int(s2)),
				AdvanceCopy[T](begin, int(s2+1)),
				cmp,
			)
			SwapIter[T](
				begin,
				AdvanceCopy[T](begin, int(s2)),
			)
		} else {
			sort3(
				AdvanceCopy[T](begin, int(s2)),
				begin,
				AdvanceCopy[T](end, -1),
				cmp,
			)
		}

		// Если предыдущий сегмент уже был отсортирован, используем другую стратегию.
		if !leftMost && !cmp(*AdvanceCopy[T](begin, -1).Ptr(), *begin.Ptr()) {
			begin = AdvanceCopy[T, interfaces.RandomAccessIterator[T]](partitionLeft(begin, end, cmp), 1)
			continue
		}

		// Разбиение массива с учетом опорного элемента.
		pivotPos, alreadyPartitioned := partitionRight[T](begin, end, cmp)

		lSize := Distance[T](begin, pivotPos)
		rSize := Distance[T](AdvanceCopy[T](pivotPos, 1), end)

		highlyUnbalanced := lSize < size/8 || rSize < size/8

		// При сильной несбалансированности, переход на сортировку кучей.
		if highlyUnbalanced {
			badAllowed--
			if badAllowed == 0 {
				MakeHeap(begin, end, cmp)
				SortHeap(begin, end, cmp)
				return
			}

			if lSize >= insertionSortThreshold {
				SwapIter[T](begin, AdvanceCopy[T](begin, int(lSize/4)))
				SwapIter[T](AdvanceCopy[T](pivotPos, -1), AdvanceCopy[T](pivotPos, -int(lSize/4)))

				if lSize > nintherThreshold {
					SwapIter[T](AdvanceCopy[T](begin, 1), AdvanceCopy[T](begin, int(lSize/4+1)))
					SwapIter[T](AdvanceCopy[T](begin, 2), AdvanceCopy[T](begin, int(lSize/4+2)))
					SwapIter[T](AdvanceCopy[T](pivotPos, -2), AdvanceCopy[T](pivotPos, -int(lSize/4+1)))
					SwapIter[T](AdvanceCopy[T](pivotPos, -3), AdvanceCopy[T](pivotPos, -int(lSize/4+2)))
				}
			}

			if rSize >= insertionSortThreshold {
				SwapIter[T](AdvanceCopy[T](pivotPos, 1), AdvanceCopy[T](pivotPos, int(1+rSize/4)))
				SwapIter[T](AdvanceCopy[T](end, -1), AdvanceCopy[T](end, -int(rSize/4)))

				if rSize > nintherThreshold {
					SwapIter[T](AdvanceCopy[T](pivotPos, 2), AdvanceCopy[T](pivotPos, int(2+rSize/4)))
					SwapIter[T](AdvanceCopy[T](pivotPos, 3), AdvanceCopy[T](pivotPos, int(3+rSize/4)))
					SwapIter[T](AdvanceCopy[T](end, -2), AdvanceCopy[T](end, -int(1+rSize/4)))
					SwapIter[T](AdvanceCopy[T](end, -3), AdvanceCopy[T](end, -int(2+rSize/4)))
				}
			}
		} else if alreadyPartitioned &&
			partialInsertionSort(begin, pivotPos, cmp) &&
			partialInsertionSort(AdvanceCopy[T](pivotPos, 1), end, cmp) {
			// Если разбиение сбалансировано, попытка сортировки вставками.
			return
		}

		// Рекурсивная сортировка левой части и хвостовая рекурсия для правой части.
		pdqsortLoop[T](begin, pivotPos, cmp, badAllowed, leftMost)
		begin = AdvanceCopy[T](pivotPos, 1)
		leftMost = false
	}
}

// insertionSort выполняет сортировку диапазона [begin, end) с использованием сортировки вставками.
// Предполагается, что *(begin - 1) — элемент, который меньше или равен любому элементу в [begin, end).
func insertionSort[T any](begin, end interfaces.RandomAccessIterator[T], cmp comparator.Less[T]) {
	if begin.Equals(end) {
		return
	}

	for cur := AdvanceCopy[T](begin, 1); !cur.Equals(end); cur.Next() {
		sift := copiable.Copy[interfaces.RandomAccessIterator[T]](cur)
		sift1 := AdvanceCopy[T](cur, -1)

		if cmp(*sift.Ptr(), *sift1.Ptr()) {
			tmp := *sift.Ptr()

			for {
				*sift.Ptr() = *sift1.Ptr()
				sift.Prev()

				if sift.Equals(begin) {
					break
				}

				sift1.Prev()
				if !(cmp(tmp, sift1.Value())) {
					break
				}
			}

			*sift.Ptr() = tmp
		}
	}
}

// partialInsertionSort пытается выполнить сортировку вставками на диапазоне [begin, end).
// Возвращает false, если было перемещено более partialInsertionSortLimit элементов
// и сортировка была прервана.
// В противном случае, сортировка завершается успешно и возвращается true.
func partialInsertionSort[T any](begin, end interfaces.RandomAccessIterator[T], cmp comparator.Less[T]) bool {
	if begin.Equals(end) {
		return true
	}

	var limit int
	for cur := AdvanceCopy[T](begin, 1); !cur.Equals(end); cur.Next() {
		sift := copiable.Copy[interfaces.RandomAccessIterator[T]](cur)
		sift1 := AdvanceCopy[T](cur, -1)

		// Сначала сравниваем, чтобы избежать лишних перемещений для уже правильно расположенных элементов.
		if cmp(*sift.Ptr(), *sift1.Ptr()) {
			tmp := *sift.Ptr()

			for {
				*sift.Ptr() = *sift1.Ptr()
				if !sift.HasPrev() {
					break
				}
				sift.Prev()

				if sift.Equals(begin) && sift1.HasPrev() {
					sift1.Prev()
					if !cmp(tmp, sift1.Value()) {
						break
					}
				}
			}

			*sift.Ptr() = tmp
			limit += int(Distance[T](sift, cur))
		}

		if limit > partialInsertionSortLimit {
			return false
		}
	}

	return true
}

// partitionLeft выполняет разбиение массива вокруг опорного элемента, помещая равные элементы слева от опорного.
// Возвращает итератор на новую позицию опорного элемента.
//
//nolint:cyclop // алгоритм реализован с оглядкой на https://github.com/orlp/pdqsort
func partitionLeft[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) interfaces.RandomAccessIterator[T] {
	pivot := begin.Value()
	first := copiable.Copy[interfaces.RandomAccessIterator[T]](begin)
	last := copiable.Copy[interfaces.RandomAccessIterator[T]](end)

	for {
		last.Prev()
		if !cmp(pivot, last.Value()) {
			break
		}
	}

	if AdvanceCopy[T](last, 1).Equals(end) {
		for first.Index() < last.Index() {
			first.Next()
			if !(cmp(pivot, first.Value())) {
				break
			}
		}
	} else {
		for {
			first.Next()
			if !(cmp(pivot, first.Value())) {
				break
			}
		}
	}

	for first.Index() < last.Index() {
		SwapIter[T](first, last)
		for {
			last.Prev()
			if !cmp(pivot, last.Value()) {
				break
			}
		}
		for {
			first.Next()
			if cmp(pivot, first.Value()) {
				break
			}
		}
	}

	pivotPos := copiable.Copy[interfaces.RandomAccessIterator[T]](last)
	*begin.Ptr() = pivotPos.Value()
	*pivotPos.Ptr() = pivot

	return pivotPos
}

// partitionRight выполняет разбиение массива [begin, end) вокруг опорного элемента.
// Элементы, равные опорному, помещаются в правую часть массива.
// Возвращает позицию опорного элемента после разбиения и флаг, указывающий, было ли разбиение уже выполнено ранее.
//
//nolint:cyclop // алгоритм реализован с оглядкой на https://github.com/orlp/pdqsort
func partitionRight[T any](
	begin, end interfaces.RandomAccessIterator[T],
	cmp comparator.Less[T],
) (interfaces.RandomAccessIterator[T], bool) {
	pivot := begin.Value()
	first := copiable.Copy[interfaces.RandomAccessIterator[T]](begin)
	last := copiable.Copy[interfaces.RandomAccessIterator[T]](end)

	for {
		first.Next()
		if !cmp(*first.Ptr(), pivot) {
			break
		}
	}

	// Поиск первого элемента, строго меньшего, чем опорный.
	if AdvanceCopy[T](first, -1).Equals(begin) {
		for first.Index() < last.Index() {
			last.Prev()
			if cmp(last.Value(), pivot) {
				break
			}
		}
	} else {
		for {
			last.Prev()
			if cmp(last.Value(), pivot) {
				break
			}
		}
	}

	alreadyPartitioned := first.Index() >= last.Index()

	// Продолжаем обмен пар элементов, которые находятся на неправильных сторонах от опорного.
	for first.Index() < last.Index() {
		SwapIter[T](first, last)

		for {
			first.Next()
			if !cmp(first.Value(), pivot) {
				break
			}
		}
		for {
			last.Prev()
			if cmp(last.Value(), pivot) {
				break
			}
		}
	}

	// Помещаем опорный элемент на его окончательное место.
	pivotPos := AdvanceCopy[T](first, -1)
	*begin.Ptr() = pivotPos.Value()
	*pivotPos.Ptr() = pivot

	return pivotPos, alreadyPartitioned
}

// sort3 выполняет сортировку трех элементов.
// Используется для улучшения выбора опорного элемента.
func sort3[T any](a, b, c interfaces.RandomAccessIterator[T], cmp comparator.Less[T]) {
	sort2(a, b, cmp)
	sort2(b, c, cmp)
	sort2(a, b, cmp)
}

// sort2 выполняет сортировку двух элементов, если они находятся в неправильном порядке.
func sort2[T any](a, b interfaces.RandomAccessIterator[T], cmp comparator.Less[T]) {
	if cmp(*b.Ptr(), *a.Ptr()) {
		SwapIter[T](a, b)
	}
}

// log2 вычисляет логарифм по основанию 2 от числа n.
func log2(n uint) uint {
	var log uint
	for {
		n >>= 1
		if n == 0 {
			break
		}
		log++
	}
	return log
}
