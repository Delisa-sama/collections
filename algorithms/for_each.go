package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// forEachFunc представляет функцию, применяемую к каждому элементу контейнера.
type forEachFunc[T any] func(T)

// ForEach применяет функцию f ко всем элементам в диапазоне от begin до end.
// На вход функции f будет копия элемента.
func ForEach[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, f forEachFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Value())
	}
}

// forEachFunc представляет функцию, применяемую к каждому элементу контейнера.
type forEachPtrFunc[T any] func(*T)

// ForEachPtr применяет функцию f ко всем элементам в диапазоне от begin до end.
// На вход функции f будет указатель на элемент.
func ForEachPtr[T any](begin interfaces.PointerIterator[T], end interfaces.Iterator, f forEachPtrFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Ptr())
	}
}
