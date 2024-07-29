package iterators

import (
	"github.com/Delisa-sama/collections/copiable"
	"github.com/Delisa-sama/collections/interfaces"
)

// EndIterator представляет собой пустой конечный итератор.
// Его можно использовать в случаях когда
// итератор на элемент после конца невозможен из-за устройства контейнера.
// В таком случае итератор обязан реализовать в методу Equals сравнение с типом EndIterator,
// которое вернет true если итератор достиг конца.
type EndIterator struct{}

// NewEndIterator возвращает новый конечный итератор.
func NewEndIterator() interfaces.Iterator {
	return &EndIterator{}
}

// Equals проверяет, равны ли два итератора.
func (it *EndIterator) Equals(another interfaces.Iterator) bool {
	switch another.(type) {
	case *EndIterator:
		return true
	default:
		return another.Equals(it)
	}
}

// HasNext проверяет, есть ли еще элементы для перебора.
func (it *EndIterator) HasNext() bool {
	return false
}

// Next переходит к следующему элементу.
func (it *EndIterator) Next() {}

// Copy возвращает новый экземпляр EndIterator, поскольку не имеет состояния.
func (it *EndIterator) Copy() copiable.Copiable {
	return NewEndIterator()
}
