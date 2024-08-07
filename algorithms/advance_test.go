package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestAdvance(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{10, 20, 30, 40, 50})

	vecIt := vec.Begin()
	Advance[int](vecIt, 3)
	assert.Equal(t, uint(3), vecIt.Index())

	// Test RandomAccessIterator backward shift
	Advance[int](vecIt, -2)
	assert.Equal(t, uint(1), vecIt.Index())

	Advance[int](vecIt, 10)
	assert.Equal(t, uint(11), vecIt.Index())

	bidirectionalIt := vec.Begin().(interfaces.BidirectionalIterator[int])
	Advance[int](bidirectionalIt, 2)
	assert.Equal(t, 30, bidirectionalIt.Value())

	Advance[int](bidirectionalIt, -1)
	assert.Equal(t, 20, bidirectionalIt.Value())
}

func TestAdvanceCopy(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{10, 20, 30, 40, 50})

	vecIt := vec.Begin()
	vecIt3 := AdvanceCopy[int](vecIt, 3)
	assert.Equal(t, uint(3), vecIt3.Index())
	assert.Equal(t, uint(0), vecIt.Index())
}
