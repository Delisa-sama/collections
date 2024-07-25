package algorithms

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// EqualsC сравнивает 2 неограниченных диапазона с использованием заданного компаратора.
// Возвращает true, если все элементы одинаковы, иначе false.
func EqualsC[T any](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T], cmp comparator.Comparator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// Equals сравнивает 2 неограниченных диапазона.
// Возвращает true, если все элементы одинаковы, иначе false.
func Equals[T comparable](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if a.Value() != b.Value() {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// EqualsRangesC сравнивает диапазоны двух итераторов с использованием заданного компаратора.
// Возвращает true, если все элементы одинаковы, иначе false.
func EqualsRangesC[T any](
	aBegin interfaces.ValueIterator[T], aEnd interfaces.Iterator,
	bBegin interfaces.ValueIterator[T], bEnd interfaces.Iterator,
	cmp comparator.Comparator[T],
) bool {
	a := aBegin
	b := bBegin
	for !a.Equals(aEnd) || !b.Equals(bEnd) {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// EqualsRanges сравнивает диапазоны двух итераторов.
// Возвращает true, если все элементы одинаковы, иначе false.
func EqualsRanges[T comparable](
	aBegin interfaces.ValueIterator[T], aEnd interfaces.Iterator,
	bBegin interfaces.ValueIterator[T], bEnd interfaces.Iterator,
) bool {
	a := aBegin
	b := bBegin
	for !a.Equals(aEnd) || !b.Equals(bEnd) {
		if a.Value() != b.Value() {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// xor возвращает true, если один из аргументов true, но не оба.
func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
