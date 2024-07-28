package algorithms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name        string
		items       []int
		middleIndex uint
		expected    []int
	}{
		{
			name:        "Rotate empty range",
			items:       []int{},
			middleIndex: 0,
			expected:    []int{},
		},
		{
			name:        "Rotate full range",
			items:       []int{1, 2, 3, 4, 5},
			middleIndex: 2,
			expected:    []int{3, 4, 5, 1, 2},
		},
		{
			name:        "Rotate single element",
			items:       []int{1},
			middleIndex: 0,
			expected:    []int{1},
		},
		{
			name:        "Rotate half range",
			items:       []int{1, 2, 3, 4, 5},
			middleIndex: 1,
			expected:    []int{2, 3, 4, 5, 1},
		},
		{
			name:        "Rotate with middle at end",
			items:       []int{1, 2, 3, 4, 5},
			middleIndex: 5,
			expected:    []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := vector.NewVector(tt.items...)
			begin := vec.Begin()
			middle := vec.At(tt.middleIndex)
			end := vec.End()

			_ = algorithms.Rotate[int](begin, middle, end)

			result := make([]int, 0)
			for it := vec.Begin(); !it.Equals(vec.End()); it.Next() {
				result = append(result, it.(interfaces.ForwardIterator[int]).Value())
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRotateCopy(t *testing.T) {
	tests := []struct {
		name        string
		items       []int
		middleIndex uint
		expected    []int
	}{
		{
			name:        "Rotate empty range",
			items:       []int{},
			middleIndex: 0,
			expected:    []int{},
		},
		{
			name:        "Rotate full range",
			items:       []int{1, 2, 3, 4, 5},
			middleIndex: 2,
			expected:    []int{3, 4, 5, 1, 2},
		},
		{
			name:        "Rotate single element",
			items:       []int{1},
			middleIndex: 0,
			expected:    []int{1},
		},
		{
			name:        "Rotate half range",
			items:       []int{1, 2, 3, 4, 5},
			middleIndex: 1,
			expected:    []int{2, 3, 4, 5, 1},
		},
		{
			name:        "Rotate with middle at end",
			items:       []int{1, 2, 3, 4, 5},
			middleIndex: 5,
			expected:    []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vec := vector.NewVector(tt.items...)
			begin := vec.Begin()
			middle := vec.At(tt.middleIndex)
			end := vec.End()
			destVec := vector.NewVectorFromSlice(make([]int, len(tt.items)))

			_ = algorithms.RotateCopy[int](begin, middle, end, destVec.Begin())

			result := make([]int, 0)
			for it := destVec.Begin(); !it.Equals(destVec.End()); it.Next() {
				result = append(result, it.(interfaces.ForwardIterator[int]).Value())
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
