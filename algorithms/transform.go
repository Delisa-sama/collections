package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// unaryTransformFunc представляет собой унарную функцию, которая принимает значение типа T
// и возвращает значение типа O. Используется для преобразования элементов.
type unaryTransformFunc[T any, O any] func(T) O

// TransformUnary применяет унарную функцию f к каждому элементу последовательности,
// начиная с итератора begin и заканчивая итератором end, и записывает результаты в
// последовательность, начинающуюся с destBegin.
//
// Параметры:
// - begin: итератор на начало последовательности исходных значений.
// - end: итератор на конец последовательности исходных значений.
// - destBegin: итератор на начало последовательности, куда будут записаны результаты.
// - f: унарная функция для преобразования.
//
// Возвращает итератор на конец последовательности результатов.
func TransformUnary[T any, O any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	destBegin interfaces.PointerIterator[O],
	f unaryTransformFunc[T, O],
) interfaces.PointerIterator[O] {
	for !begin.Equals(end) {
		*destBegin.Ptr() = f(begin.Value())
		destBegin.Next()
		begin.Next()
	}

	return destBegin
}

// binaryTransformFunc представляет собой бинарную функцию, которая принимает два значения типов T1 и T2
// и возвращает значение типа O. Используется для преобразования пар элементов.
type binaryTransformFunc[T1 any, T2 any, O any] func(T1, T2) O

// TransformBinary применяет бинарную функцию f к парам элементов из двух последовательностей,
// начиная с итераторов begin1 и begin2, и записывает результаты в последовательность, начинающуюся с destBegin.
//
// Параметры:
// - begin1: итератор на начало первой последовательности.
// - end1: итератор на конец первой последовательности.
// - begin2: итератор на начало второй последовательности.
// - destBegin: итератор на начало последовательности, куда будут записаны результаты.
// - f: бинарная функция для преобразования.
//
// Возвращает итератор на конец последовательности результатов.
func TransformBinary[T1 any, T2 any, O any](
	begin1 interfaces.ValueIterator[T1],
	end1 interfaces.Iterator,
	begin2 interfaces.ValueIterator[T2],
	destBegin interfaces.PointerIterator[O],
	f binaryTransformFunc[T1, T2, O],
) interfaces.PointerIterator[O] {
	for !begin1.Equals(end1) {
		*destBegin.Ptr() = f(begin1.Value(), begin2.Value())

		begin1.Next()
		begin2.Next()
		destBegin.Next()
	}
	return destBegin
}
