package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// forEachFunc представляет функцию, применяемую к каждому элементу контейнера.
type forEachFunc[T any] func(T)

// ForEach применяет функцию f ко всем элементам в диапазоне от begin до end.
func ForEach[T any](begin interfaces.ForwardIterator[T], end interfaces.Iterator, f forEachFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Value())
	}
}
