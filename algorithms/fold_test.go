package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestFoldLeft(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4})
	result := FoldLeft(vec.Begin(), vec.End(), 0, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, 10, result)
}

func TestFoldLeftFirst(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4})
	result := FoldLeftFirst(vec.Begin(), vec.End(), func(a, b int) int {
		return a + b
	})
	assert.NotNil(t, result)
	assert.Equal(t, 10, *result)
}

func TestFoldRight(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4})
	result := FoldRight(vec.Begin(), vec.End(), 0, func(a, b int) int {
		return a + b
	})
	assert.Equal(t, 10, result)
}

func TestFoldRightLast(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4})
	result := FoldRightLast(vec.Begin(), vec.End(), func(a, b int) int {
		return a + b
	})
	assert.NotNil(t, result)
	assert.Equal(t, 10, *result)
}
