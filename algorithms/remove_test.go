package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestRemove(t *testing.T) {
	data := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4, 2, 5})
	first := data.Begin()
	last := data.End()

	newEnd := Remove(first, last, 2)

	expected := []int{1, 3, 4, 5}
	assert.Equal(t, uint(7), data.Size())
	ForEachIdx(data.Begin(), newEnd, func(idx uint, value int) {
		assert.Equal(t, expected[idx], value)
	})
}

func TestRemoveC(t *testing.T) {
	data := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4, 2, 5})
	first := data.Begin()
	last := data.End()

	newEnd := RemoveC(first, last, 2, comparator.DefaultComparator[int]())

	expected := []int{1, 3, 4, 5}
	assert.Equal(t, uint(7), data.Size())
	ForEachIdx(data.Begin(), newEnd, func(idx uint, value int) {
		assert.Equal(t, expected[idx], value)
	})
}

func TestRemoveIf(t *testing.T) {
	data := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4, 2, 5})
	first := data.Begin()
	last := data.End()

	newEnd := RemoveIf(first, last, func(v int) bool {
		return v == 2
	})

	expected := []int{1, 3, 4, 5}
	assert.Equal(t, uint(7), data.Size())
	ForEachIdx(data.Begin(), newEnd, func(idx uint, value int) {
		assert.Equal(t, expected[idx], value)
	})
}
