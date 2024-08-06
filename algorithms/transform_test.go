package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestTransformUnary(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4})
	dest := vector.NewVectorFromSlice(make([]int, vec.Size()))

	TransformUnary(vec.Begin(), vec.End(), dest.Begin(), func(a int) int {
		return a * 2
	})

	expected := vector.NewVectorFromSlice([]int{2, 4, 6, 8})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}

func TestTransformBinary(t *testing.T) {
	data1 := vector.NewVectorFromSlice([]int{1, 2, 3, 4})
	data2 := vector.NewVectorFromSlice([]int{10, 20, 30, 40})
	dest := vector.NewVectorFromSlice(make([]int, data1.Size()))

	TransformBinary(data1.Begin(), data1.End(), data2.Begin(), dest.Begin(), func(a, b int) int {
		return a + b
	})

	expected := vector.NewVectorFromSlice([]int{11, 22, 33, 44})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}
