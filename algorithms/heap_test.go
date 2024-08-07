package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestMakeHeap(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
	MakeHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())

	expected := vector.NewVectorFromSlice([]int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1})
	assert.Truef(t, Equals[int](vec.Begin(), expected.Begin()), "expected: %+v, got: %+v", expected, vec)
}

func TestSortHeap(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{3, 1, 4, 1, 5, 9})
	MakeHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())

	SortHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())

	expected := vector.NewVectorFromSlice([]int{1, 1, 3, 4, 5, 9})
	assert.Truef(t, Equals[int](vec.Begin(), expected.Begin()), "expected: %+v, got: %+v", expected, vec)
}

func TestPopHeap(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{3, 1, 4, 1, 5, 9})
	MakeHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())

	PopHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())
	largest := vec.Back()
	vec.PopBack()

	assert.Equal(t, 9, largest)
	expected := vector.NewVectorFromSlice([]int{5, 3, 4, 1, 1})
	assert.Truef(t, Equals[int](vec.Begin(), expected.Begin()), "expected: %+v, got: %+v", expected, vec)
}

func TestPushHeap(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{3, 1, 4, 1, 5, 9})
	MakeHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())
	vec.PushBack(6)

	PushHeap(vec.Begin(), vec.End(), comparator.DefaultLess[int]())

	expected := vector.NewVectorFromSlice([]int{9, 5, 6, 1, 1, 3, 4})
	assert.Truef(t, Equals[int](vec.Begin(), expected.Begin()), "expected: %+v, got: %+v", expected, vec)
}
