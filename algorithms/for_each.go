package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// forEachFunc представляет функцию, принимающую элемент типа T.
type forEachFunc[T any] func(T)

// ForEach применяет функцию к каждому элементу в диапазоне [begin, end).
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в применение функции).
// - f: функция, которая применяется к каждому элементу.
func ForEach[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, f forEachFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Value())
	}
}

// forEachPtrFunc представляет функцию, принимающую указатель на элемент типа T.
type forEachPtrFunc[T any] func(*T)

// ForEachPtr применяет функцию к указателю на каждый элемент в диапазоне [begin, end).
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в применение функции).
// - f: функция, которая применяется к каждому указателю на элемент.
func ForEachPtr[T any](begin interfaces.PointerIterator[T], end interfaces.Iterator, f forEachPtrFunc[T]) {
	for it := begin; !it.Equals(end); it.Next() {
		f(it.Ptr())
	}
}

// forEachFunc представляет функцию, принимающую uint индекс и элемент типа T.
type forEachIdxFunc[T any] func(uint, T)

// ForEachIdx применяет функцию к каждому элементу в диапазоне [begin, end),
// передавая в функцию элемент и его индекс.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в применение функции).
// - f: функция, которая применяется к каждому элементу.
func ForEachIdx[T any](begin interfaces.ValueIterator[T], end interfaces.Iterator, f forEachIdxFunc[T]) {
	var idx uint
	for it := begin; !it.Equals(end); it.Next() {
		f(idx, it.Value())
		idx++
	}
}

// forEachFunc представляет функцию, принимающую uint индекс и указатель на элемент типа T.
type forEachIdxPtrFunc[T any] func(uint, *T)

// ForEachIdxPtr применяет функцию к указателю каждый элемент в диапазоне [begin, end),
// передавая в функцию указатель на элемент и его индекс.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включается в применение функции).
// - f: функция, которая применяется к каждому элементу.
func ForEachIdxPtr[T any](begin interfaces.PointerIterator[T], end interfaces.Iterator, f forEachIdxPtrFunc[T]) {
	var idx uint
	for it := begin; !it.Equals(end); it.Next() {
		f(idx, it.Ptr())
		idx++
	}
}
