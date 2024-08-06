package algorithms

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// binaryFoldFunc - это тип функции, которая принимает два значения типа T и возвращает значение типа T.
// Она используется в функциях свёртки для комбинирования элементов.
type binaryFoldFunc[T any] func(T, T) T

// FoldLeft выполняет левую свёртку последовательности, начиная с инициализирующего значения init
// и применяя функцию f последовательно ко всем элементам от begin до end.
// Если итератор begin равен итератору end, функция вернёт значение init.
//
// Аргументы:
// - begin: итератор, указывающий на первый элемент последовательности.
// - end: итератор, указывающий на конец последовательности (не включается в свёртку).
// - init: начальное значение для свёртки.
// - f: функция, которая будет применяться к элементам последовательности.
//
// Возвращает:
// - значение типа T, полученное после применения функции f ко всем элементам последовательности.
func FoldLeft[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	init T,
	f binaryFoldFunc[T],
) T {
	if begin.Equals(end) {
		return init
	}
	accum := f(init, begin.Value())
	begin.Next()
	for ; !begin.Equals(end); begin.Next() {
		accum = f(accum, begin.Value())
	}
	return accum
}

// FoldLeftFirst выполняет левую свёртку последовательности, используя в качестве начального значения
// первый элемент последовательности. Если последовательность пуста, возвращается nil.
//
// Аргументы:
// - begin: итератор, указывающий на первый элемент последовательности.
// - end: итератор, указывающий на конец последовательности (не включается в свёртку).
// - f: функция, которая будет применяться к элементам последовательности.
//
// Возвращает:
// - указатель на значение типа T, полученное после применения функции f ко всем элементам последовательности.
// - nil, если последовательность пуста.
func FoldLeftFirst[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	f binaryFoldFunc[T],
) *T {
	if begin.Equals(end) {
		return nil
	}
	accum := begin.Value()
	begin.Next()
	for ; !begin.Equals(end); begin.Next() {
		accum = f(accum, begin.Value())
	}
	return &accum
}

// FoldRight выполняет правую свёртку последовательности, начиная с инициализирующего значения init
// и применяя функцию f последовательно ко всем элементам от end до begin в обратном порядке.
// Если итератор begin равен итератору end, функция вернёт значение init.
//
// Аргументы:
// - begin: двунаправленный итератор, указывающий на первый элемент последовательности.
// - end: итератор, указывающий на конец последовательности (не включается в свёртку).
// - init: начальное значение для свёртки.
// - f: функция, которая будет применяться к элементам последовательности.
//
// Возвращает:
// - значение типа T, полученное после применения функции f ко всем элементам последовательности в обратном порядке.
func FoldRight[T any](
	begin interfaces.BidirectionalIterator[T],
	end interfaces.Iterator,
	init T,
	f binaryFoldFunc[T],
) T {
	if begin.Equals(end) {
		return init
	}

	tail := copiable.Copy[interfaces.BidirectionalIterator[T]](begin)
	NextBound(tail, end)
	tail.Prev()
	accum := f(tail.Value(), init)
	for !begin.Equals(tail) {
		tail.Prev()
		accum = f(tail.Value(), accum)
	}
	return accum
}

// FoldRightLast выполняет правую свёртку последовательности, используя в качестве начального значения
// последний элемент последовательности. Если последовательность пуста, возвращается nil.
//
// Аргументы:
// - begin: двунаправленный итератор, указывающий на первый элемент последовательности.
// - end: итератор, указывающий на конец последовательности (не включается в свёртку).
// - f: функция, которая будет применяться к элементам последовательности.
//
// Возвращает:
// - указатель на значение типа T, полученное после применения функции f ко всем элементам последовательности
// в обратном порядке.
// - nil, если последовательность пуста.
func FoldRightLast[T any](
	begin interfaces.BidirectionalIterator[T],
	end interfaces.Iterator,
	f binaryFoldFunc[T],
) *T {
	if begin.Equals(end) {
		return nil
	}
	tail := copiable.Copy[interfaces.BidirectionalIterator[T]](begin)
	NextBound(tail, end)
	tail.Prev()

	result := FoldRight(begin, tail, tail.Value(), f)
	return &result
}
