package iterators

import (
	"github.com/Delisa-sama/collections/interfaces"
)

type EndIterator struct{}

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
