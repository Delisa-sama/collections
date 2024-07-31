package algorithms

import "github.com/Delisa-sama/collections/interfaces"

// Fill заполняет диапазон значений от итератора `begin` до итератора `end`
// значением `value`. Итераторы должны поддерживать интерфейс PointerIterator и Iterator,
// предоставляющий методы Equals, Next и Ptr.
//
// Параметры:
// - begin: начальный итератор, откуда начинается заполнение.
// - end: конечный итератор, где заканчивается заполнение. Значение по этому итератору не включается.
// - value: значение, которым заполняется диапазон.
func Fill[T any](begin interfaces.PointerIterator[T], end interfaces.Iterator, value T) {
	for it := begin; !it.Equals(end); it.Next() {
		*it.Ptr() = value
	}
}
