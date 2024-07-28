package interfaces

import (
	"github.com/Delisa-sama/collections/copiable"
)

// Container - это обобщенный интерфейс для контейнера, который содержит элементы любого типа.
type Container[T any] interface {
	copiable.Copiable
	// Size возвращает количество элементов в контейнере.
	Size() uint
	// IsEmpty проверяет пуст ли контейнер.
	IsEmpty() bool
}
