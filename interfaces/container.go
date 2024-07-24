package interfaces

// Container - это обобщенный интерфейс для контейнера, который содержит элементы любого типа.
type Container[T any] interface {
	// Size возвращает количество элементов в контейнере.
	Size() uint
	IsEmpty() bool
}
