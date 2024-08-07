package algorithms

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// Advance продвигает итератор it на n шагов вперед или назад.
//
// Функция поддерживает три вида итераторов:
//   - RandomAccessIterator: итератор с произвольным доступом. Если итератор реализует этот интерфейс,
//     он сдвигается на n шагов с помощью метода Shift. Этот метод оптимален по скорости для итераторов
//     с произвольным доступом.
//   - BidirectionalIterator: итератор с двусторонним доступом. Если итератор реализует этот интерфейс,
//     и n отрицательное, итератор сдвигается назад на n шагов с помощью метода Prev.
//   - UnidirectionalIterator: итератор с однонаправленным доступом. Если итератор не поддерживает
//     произвольный или двусторонний доступ, он сдвигается на n шагов вперед с помощью метода Next.
//
// Параметры:
//   - it: итератор, который будет сдвинут.
//   - n: количество шагов для сдвига итератора. Если n положительное, итератор двигается вперед.
//     Если n отрицательное, итератор двигается назад (при поддержке двустороннего доступа).
//
// Поведение функции зависит от типа итератора:
// - RandomAccessIterator: сдвиг производится одним вызовом Shift.
// - BidirectionalIterator: сдвиг вперед производится через Next, назад — через Prev.
// - UnidirectionalIterator: поддерживается только сдвиг вперед через Next. Сдвиг назад невозможен.
func Advance[T any](it interfaces.Iterator, n int) {
	if randIt, ok := it.(interfaces.RandomAccessIterator[T]); ok {
		randIt.Shift(n)
		return
	}

	for n > 0 {
		n--
		it.Next()
	}
	if bidirectionalIt, ok := it.(interfaces.BidirectionalIterator[T]); ok {
		for n < 0 {
			n++
			bidirectionalIt.Prev()
		}
	}
}

// AdvanceCopy продвигает копию итератора it на n шагов вперед или назад и возвращает его.
func AdvanceCopy[T any, It interfaces.Iterator](it It, n int) It {
	c := copiable.Copy[It](it)
	Advance[T](c, n)
	return c
}
