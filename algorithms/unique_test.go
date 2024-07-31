package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestUnique(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 2, 3, 3, 3, 4, 5, 5})

	begin := vec.Begin()
	end := vec.End()

	newEnd := Unique[int](begin, end)

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, uint(9), vec.Size())
	ForEachIdx(vec.Begin(), newEnd, func(idx uint, value int) {
		assert.Equal(t, expected[idx], value)
	})
}

func TestUniqueC(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 2, 3, 3, 3, 4, 5, 5})

	begin := vec.Begin()
	end := vec.End()

	newEnd := UniqueC[int](begin, end, comparator.DefaultComparator[int]())

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, uint(9), vec.Size())
	ForEachIdx(vec.Begin(), newEnd, func(idx uint, value int) {
		assert.Equal(t, expected[idx], value)
	})
}

func TestUniqueIf(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 2, 3, 3, 3, 4, 5, 5})

	begin := vec.Begin()
	end := vec.End()

	newEnd := UniqueIf[int](begin, end, func(a int, b int) bool {
		return a == b
	})

	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, uint(9), vec.Size())
	ForEachIdx(vec.Begin(), newEnd, func(idx uint, value int) {
		assert.Equal(t, expected[idx], value)
	})
}
