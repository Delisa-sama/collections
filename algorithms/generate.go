package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// generatorFunc представляет собой функцию, которая генерирует значения типа T.
type generatorFunc[T any] func() T

// Generate заполняет последовательность значений, начиная с итератора begin и заканчивая итератором end,
// значениями, генерируемыми функцией g.
//
// Параметры:
// - begin: итератор на начало последовательности, куда будут записаны результаты.
// - end: итератор на конец последовательности.
// - g: функция-генератор, создающая значения типа T.
func Generate[T any](
	begin interfaces.PointerIterator[T],
	end interfaces.Iterator,
	g generatorFunc[T],
) {
	for ; !begin.Equals(end); begin.Next() {
		*begin.Ptr() = g()
	}
}

// GenerateN заполняет n элементов последовательности, начиная с итератора begin,
// значениями, генерируемыми функцией g.
//
// Параметры:
// - begin: итератор на начало последовательности, куда будут записаны результаты.
// - n: количество элементов для заполнения.
// - g: функция-генератор, создающая значения типа T.
func GenerateN[T any](
	begin interfaces.PointerIterator[T],
	n uint,
	g generatorFunc[T],
) {
	for i := uint(0); i < n; i++ {
		*begin.Ptr() = g()
		begin.Next()
	}
}
