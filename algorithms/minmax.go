package algorithms

import (
	"cmp"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Min находит минимальный элемент в диапазоне [begin, end) используя естественный порядок элементов (cmp.Ordered).
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// Возвращает итератор на минимальный элемент. Если диапазон пуст, возвращает begin.
func Min[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
) interfaces.ValueIterator[T] {
	return MinC(begin, end, comparator.DefaultComparator[T]())
}

// MinC находит минимальный элемент в диапазоне [begin, end) используя пользовательский компаратор.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// - cmp: компаратор, определяющий порядок элементов.
//
// Возвращает итератор на минимальный элемент. Если диапазон пуст, возвращает begin.
func MinC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	cmp comparator.Comparator[T],
) interfaces.ValueIterator[T] {
	if begin.Equals(end) {
		return begin
	}

	smallest := copiable.Copy[interfaces.ValueIterator[T]](begin)
	begin.Next()
	for ; !begin.Equals(end); begin.Next() {
		if cmp(begin.Value(), smallest.Value()) < 0 {
			smallest = copiable.Copy[interfaces.ValueIterator[T]](begin)
		}
	}

	return smallest
}

// Max находит максимальный элемент в диапазоне [begin, end) используя естественный порядок элементов (cmp.Ordered).
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// Возвращает итератор на максимальный элемент. Если диапазон пуст, возвращает begin.
func Max[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
) interfaces.ValueIterator[T] {
	return MaxC(begin, end, comparator.DefaultComparator[T]())
}

// MaxC находит максимальный элемент в диапазоне [begin, end) используя пользовательский компаратор.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// - cmp: компаратор, определяющий порядок элементов.
//
// Возвращает итератор на максимальный элемент. Если диапазон пуст, возвращает begin.
func MaxC[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	cmp comparator.Comparator[T],
) interfaces.ValueIterator[T] {
	if begin.Equals(end) {
		return begin
	}

	largest := copiable.Copy[interfaces.ValueIterator[T]](begin)
	begin.Next()
	for ; !begin.Equals(end); begin.Next() {
		if cmp(largest.Value(), begin.Value()) < 0 {
			largest = copiable.Copy[interfaces.ValueIterator[T]](begin)
		}
	}

	return largest
}

// MinMax находит одновременно минимальный и максимальный элементы в диапазоне [begin, end)
// используя естественный порядок элементов (cmp.Ordered).
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// Возвращает два итератора: первый на минимальный элемент, второй на максимальный.
// Если диапазон пуст, оба итератора указывают на begin.
//
//nolint:gocritic // min, max говорящие названия для возвращаемых значений, хоть и затеняют встроенные функции.
func MinMax[T cmp.Ordered](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
) (min, max interfaces.ValueIterator[T]) {
	return MinMaxC(begin, end, comparator.DefaultComparator[T]())
}

// MinMaxC находит одновременно минимальный и максимальный элементы в диапазоне [begin, end)
// используя пользовательский компаратор.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// - cmp: компаратор, определяющий порядок элементов.
//
// Возвращает два итератора: первый на минимальный элемент, второй на максимальный.
// Если диапазон пуст, оба итератора указывают на begin.
//
// min, max говорящие названия для возвращаемых значений, хоть и затеняют встроенные функции.
//
//nolint:cyclop,gocritic // невозможно уменьшить сложность без потери производительности
func MinMaxC[T any](
	first interfaces.ValueIterator[T],
	last interfaces.Iterator,
	cmp comparator.Comparator[T],
) (min, max interfaces.ValueIterator[T]) {
	min = copiable.Copy[interfaces.ValueIterator[T]](first)
	max = copiable.Copy[interfaces.ValueIterator[T]](first)

	if first.Equals(last) {
		return min, max
	}
	first.Next()
	if first.Equals(last) {
		return min, max
	}

	if cmp(first.Value(), min.Value()) < 0 {
		min = copiable.Copy[interfaces.ValueIterator[T]](first)
	} else {
		max = copiable.Copy[interfaces.ValueIterator[T]](first)
	}

	first.Next()
	for !first.Equals(last) {
		i := copiable.Copy[interfaces.ValueIterator[T]](first)
		first.Next()
		if first.Equals(last) {
			if cmp(i.Value(), min.Value()) < 0 {
				min = copiable.Copy[interfaces.ValueIterator[T]](i)
			} else if cmp(i.Value(), max.Value()) > 0 {
				max = copiable.Copy[interfaces.ValueIterator[T]](i)
			}
			break
		}

		if cmp(first.Value(), i.Value()) < 0 {
			if cmp(first.Value(), min.Value()) < 0 {
				min = copiable.Copy[interfaces.ValueIterator[T]](first)
			}
			if cmp(i.Value(), max.Value()) > 0 {
				max = copiable.Copy[interfaces.ValueIterator[T]](i)
			}
		} else {
			if cmp(i.Value(), min.Value()) < 0 {
				min = copiable.Copy[interfaces.ValueIterator[T]](i)
			}
			if cmp(first.Value(), max.Value()) > 0 {
				max = copiable.Copy[interfaces.ValueIterator[T]](first)
			}
		}
	}

	return min, max
}
