package containers

import (
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/interfaces"
)

// Equals сравнивает 2 неограниченных диапазона с использованием заданного компаратора.
// Возвращает true, если все элементы одинаковы, иначе false.
func Equals[T any](a interfaces.ForwardIterator[T], b interfaces.ForwardIterator[T], cmp comparator.Comparator[T]) bool {
	for a.HasNext() && b.HasNext() {
		if cmp(a.Value(), b.Value()) != 0 {
			return false
		}
		a.Next()
		b.Next()
	}

	return !xor(a.HasNext(), b.HasNext())
}

// EqualsRanges сравнивает диапазоны двух итераторов с использованием заданного компаратора.
// Возвращает true, если все элементы одинаковы, иначе false.
func EqualsRanges[T any](
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

// forEachFunc представляет функцию, применяемую к каждому элементу контейнера.
type forEachFunc[T any] func(T)

// ForEach применяет функцию f ко всем элементам в диапазоне от begin до end.
func ForEach[T any](begin interfaces.ForwardIterator[T], end interfaces.Iterator, f forEachFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Value())
	}
}

// xor возвращает true, если один из аргументов true, но не оба.
func xor(x, y bool) bool {
	return (x || y) && !(x && y)
}
