package interfaces

// Iterator - это интерфейс для перебора коллекции.
type Iterator interface {
	// Equals проверяет, равны ли два итератора.
	Equals(iterator Iterator) bool
	// HasNext проверяет, есть ли еще элементы для перебора.
	HasNext() bool
	// Next переходит к следующему элементу.
	Next()
	// Copy копирует итератор.
	// Реализация обязана вернуть тот же тип которым она является,
	// чтобы можно было "безопасно" расширить интерфейс в месте вызова копирования, пример:
	// foo[T any](begin ForwardIterator[T]) {
	//      // Copy вернет интерфейс Iterator, но мы знаем что на самом деле это ForwardIterator[T].
	//      beginCopy := begin.Copy().(ForwardIterator[T])
	//      fmt.Println(beginCopy.Value())
	// }
	Copy() Iterator
}

// ValueIterator - это интерфейс для итератора, который возвращает значение.
type ValueIterator[T any] interface {
	Iterator
	// Value возвращает текущее значение итератора.
	Value() T
}

// PointerIterator - это интерфейс для итератора, который возвращает указатель на значение.
type PointerIterator[T any] interface {
	Iterator
	// Ptr возвращает указатель на текущее значение итератора.
	Ptr() *T
}

// ForwardIterator - это интерфейс для прямого итератора, который объединяет ValueIterator и PointerIterator.
type ForwardIterator[T any] interface {
	ValueIterator[T]
	PointerIterator[T]
}

// BidirectionalIterator - это интерфейс для двунаправленного итератора, который объединяет ForwardIterator и добавляет методы для обратного перебора.
type BidirectionalIterator[T any] interface {
	ForwardIterator[T]
	// HasPrev проверяет, есть ли предыдущие элементы для перебора.
	HasPrev() bool
	// Prev переходит к предыдущему элементу.
	Prev()
}

// RandomAccessIterator - это интерфейс для итератора с произвольным доступом, который объединяет BidirectionalIterator и добавляет метод для доступа по индексу.
type RandomAccessIterator[T any] interface {
	BidirectionalIterator[T]
	// At проверяет, доступен ли элемент по заданному индексу.
	At(index uint) (*T, bool)
	Shift(offset int)
}
