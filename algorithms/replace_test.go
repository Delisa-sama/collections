package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestReplace(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4})
	Replace(vec.Begin(), vec.End(), 2, 5)

	expected := vector.NewVectorFromSlice([]int{1, 5, 3, 5, 4})
	assert.True(t, Equals[int](vec.Begin(), expected.Begin()))
}

func TestReplaceIf(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4})

	ReplaceIf(vec.Begin(), vec.End(), func(val int) bool {
		return val%2 == 0
	}, 0)

	expected := vector.NewVectorFromSlice([]int{1, 0, 3, 0, 0})
	assert.True(t, Equals[int](vec.Begin(), expected.Begin()))
}

func TestReplaceCopy(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4})
	dest := vector.NewVectorFromSlice(make([]int, vec.Size()))

	ReplaceCopy(vec.Begin(), vec.End(), dest.Begin(), 2, 5)

	expected := vector.NewVectorFromSlice([]int{1, 5, 3, 5, 4})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}

func TestReplaceCopyIf(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{1, 2, 3, 2, 4})
	dest := vector.NewVectorFromSlice(make([]int, vec.Size()))

	ReplaceCopyIf(vec.Begin(), vec.End(), dest.Begin(), func(val int) bool {
		return val%2 == 0
	}, 0)

	expected := vector.NewVectorFromSlice([]int{1, 0, 3, 0, 0})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}
