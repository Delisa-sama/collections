package interfaces

// Container - это обобщенный интерфейс для контейнера, который содержит элементы любого типа.
type Container[T any] interface {
	// Size возвращает количество элементов в контейнере.
	Size() uint
	// IsEmpty проверяет пуст ли контейнер.
	IsEmpty() bool
	// Copy возвращает копию контейнера.
	// Реализация обязана вернуть тот же тип которым она является,
	// чтобы можно было "безопасно" расширить интерфейс в месте вызова копирования, пример:
	// func (s *Stack[T, C]) Copy() *Stack[T, C] {
	//	return &Stack[T, C]{
	//		c: s.c.Copy().(C),
	//	}
	// }
	Copy() Container[T]
}
