package algorithms

import (
	"fmt"

	"github.com/Delisa-sama/collections/interfaces"
)

// Определение типа функции для вывода, совместимой с fmt.Print.
type printFunc func(a ...any) (n int, err error)

// Обеспечение соответствия printFunc типу функции fmt.Print.
var _ printFunc = fmt.Print

// Print выводит значения, начиная с итератора begin до итератора end, используя fmt.Print.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// Возвращает:
// - n: количество записанных байтов.
// - err: ошибка, если она возникла.
func Print[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
) (n int, err error) {
	return PrintFunc(begin, end, fmt.Print)
}

// Println выводит значения, начиная с итератора begin до итератора end, используя fmt.Println.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// Возвращает:
// - n: количество записанных байтов.
// - err: ошибка, если она возникла.
func Println[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
) (n int, err error) {
	return PrintFunc(begin, end, fmt.Println)
}

// PrintFunc выводит значения, начиная с итератора begin до итератора end, используя заданную функцию вывода.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// - f: функция вывода, соответствующая типу printFunc.
// Возвращает:
// - totalN: общее количество записанных байтов.
// - err: ошибка, если она возникла.
func PrintFunc[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	f printFunc,
) (totalN int, err error) {
	for it := begin; !it.Equals(end); it.Next() {
		n, err := f(it.Value())
		if err != nil {
			return 0, err
		}
		totalN += n
	}

	return totalN, nil
}

// Определение типа функции для форматированного вывода, совместимой с fmt.Printf.
type printfFunc func(format string, a ...any) (n int, err error)

// Обеспечение соответствия printfFunc типу функции fmt.Printf.
var _ printfFunc = fmt.Printf

// PrintF выводит значения, начиная с итератора begin до итератора end, используя fmt.Printf и заданный формат.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// - format: строка формата для вывода.
// Возвращает:
// - n: количество записанных байтов.
// - err: ошибка, если она возникла.
func PrintF[T any](
	begin interfaces.ValueIterator[T], end interfaces.Iterator,
	format string,
) (n int, err error) {
	return PrintFFunc(begin, end, format, fmt.Printf)
}

// PrintFFunc выводит значения, начиная с итератора begin до итератора end,
// используя заданную функцию форматированного вывода.
//
// Параметры:
// - begin: итератор, реализующий интерфейс ValueIterator, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (невключительно).
// - format: строка формата для вывода.
// - f: функция форматированного вывода, соответствующая типу printfFunc.
// Возвращает:
// - totalN: общее количество записанных байтов.
// - err: ошибка, если она возникла.
func PrintFFunc[T any](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	format string,
	f printfFunc,
) (totalN int, err error) {
	for it := begin; !it.Equals(end); it.Next() {
		n, err := f(format, it.Value())
		if err != nil {
			return 0, err
		}
		totalN += n
	}

	return totalN, nil
}
