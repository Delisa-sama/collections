package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestReverse(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4, 5})
	Reverse[int](vec.Begin(), vec.End())

	expected := vector.NewVectorFromSlice([]int{5, 4, 3, 2, 1})
	assert.True(t, Equals[int](vec.Begin(), expected.Begin()))
}

func TestReverseCopy(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 4, 5})
	dest := vector.NewVectorFromSlice(make([]int, vec.Size()))
	ReverseCopy[int](vec.Begin(), vec.End(), dest.Begin())

	expected := vector.NewVectorFromSlice([]int{5, 4, 3, 2, 1})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}
