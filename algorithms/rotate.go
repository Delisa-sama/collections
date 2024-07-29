package algorithms

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Rotate выполняет циклический сдвиг элементов в диапазоне [begin, end) так,
// чтобы элемент middle стал первым элементом диапазона.
// Элементы, находящиеся между begin и middle, будут перемещены в конец диапазона.
//
// begin - итератор, указывающий на начало диапазона.
// middle - итератор, указывающий на элемент, который станет первым после сдвига.
// end - итератор, указывающий на конец диапазона.
//
// Возвращает итератор на новый конец диапазона.
func Rotate[T any](begin, middle interfaces.ForwardIterator[T], end interfaces.Iterator) interfaces.Iterator {
	if begin.Equals(middle) {
		return end
	}
	if middle.Equals(end) {
		return begin
	}

	write := copiable.Copy[interfaces.ForwardIterator[T]](begin)
	nextRead := copiable.Copy[interfaces.ForwardIterator[T]](begin)

	for read := copiable.Copy[interfaces.ForwardIterator[T]](middle); !read.Equals(end); {
		if write.Equals(nextRead) {
			nextRead = copiable.Copy[interfaces.ForwardIterator[T]](read)
		}
		SwapIter[T](write, read)

		write.Next()
		read.Next()
	}
	Rotate(write, nextRead, end)
	return write
}

// RotateCopy выполняет циклический сдвиг элементов в диапазоне [begin, end) так,
// чтобы элемент nBegin стал первым элементом диапазона,
// и копирует результат в другой диапазон, начинающийся с destBegin.
//
// begin - итератор, указывающий на начало исходного диапазона.
// nBegin - итератор, указывающий на элемент, который станет первым после сдвига.
// end - итератор, указывающий на конец исходного диапазона.
// destBegin - итератор, указывающий на начало целевого диапазона, в который будет скопирован результат.
//
// Возвращает итератор на конец целевого диапазона после копирования.
func RotateCopy[T any](
	begin, nBegin interfaces.ForwardIterator[T],
	end interfaces.Iterator,
	destBegin interfaces.PointerIterator[T],
) interfaces.PointerIterator[T] {
	copiable.Copy[interfaces.ForwardIterator[T]](nBegin)
	destBegin = Copy(
		copiable.Copy[interfaces.ForwardIterator[T]](nBegin),
		copiable.Copy[interfaces.Iterator](end),
		copiable.Copy[interfaces.PointerIterator[T]](destBegin),
	)
	return Copy(begin, nBegin, destBegin)
}
