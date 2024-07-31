package algorithms

import (
	"github.com/Delisa-sama/collections/interfaces"
)

// Accumulate вычисляет сумму всех элементов в диапазоне [begin, end) начиная с начального значения init.
// Функция проходит по всем элементам диапазона, добавляя каждый элемент к переменной init.
//
// Параметры:
// - begin: итератор, указывающий на начало диапазона.
// - end: итератор, указывающий на конец диапазона (не включительно).
// - init: начальное значение, с которого начинается накопление.
//
// Возвращает:
// - итоговую сумму всех элементов диапазона, включая начальное значение init.
func Accumulate[T interfaces.Numeric](
	begin interfaces.ValueIterator[T],
	end interfaces.Iterator,
	init T,
) T {
	ForEach(begin, end, func(value T) {
		init += value
	})
	return init
}
