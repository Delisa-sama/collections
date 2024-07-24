package set

import (
	"github.com/elliotchance/orderedmap/v2"

	"github.com/Delisa-sama/collections/interfaces"
)

// Set представляет собой множество уникальных значений.
type Set[K comparable] struct {
	m *orderedmap.OrderedMap[K, struct{}]
}

// NewSet создает новое множество и заполняет его переданными значениями.
func NewSet[K comparable](items ...K) *Set[K] {
	m := orderedmap.NewOrderedMap[K, struct{}]()
	for i := range items {
		m.Set(items[i], struct{}{})
	}
	return &Set[K]{m: m}
}

// Size возвращает количество элементов в множестве.
func (s *Set[K]) Size() uint {
	return uint(s.m.Len())
}

// IsEmpty проверяет что множество пустое.
func (s *Set[K]) IsEmpty() bool {
	return s.Size() == 0
}

// Begin возвращает итератор на первый элемент множества.
func (s *Set[K]) Begin() interfaces.BidirectionalIterator[K] {
	front := s.m.Front()
	if front == nil {
		return newIterator[K](nil)
	}
	return newIterator[K](front)
}

// End возвращает итератор на последний элемент множества.
func (s *Set[K]) End() interfaces.BidirectionalIterator[K] {
	return newIterator(s.m.Back())
}

// Set добавляет новый элемент в множество.
func (s *Set[K]) Set(k K) {
	s.m.Set(k, struct{}{})
}
