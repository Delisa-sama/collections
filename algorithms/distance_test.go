package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestDistanceWithRandomAccessIterator(t *testing.T) {
	vec := vector.NewVector(1, 2, 3, 4, 5)

	begin := vec.Begin()
	end := vec.End()

	// Проверка полного диапазона
	result := Distance[int](begin, end)
	expected := uint(5)
	assert.Equal(t, expected, result, "Distance should return the number of elements in the range")

	// Проверка диапазона с середины до конца
	mid := vec.Begin()
	mid.Shift(2)

	result = Distance[int](mid, end)
	expected = uint(3)
	assert.Equal(t, expected, result, "Distance should return the number of elements from mid to end")

	// Проверка диапазона от начала до середины
	result = Distance[int](begin, mid)
	expected = uint(2)
	assert.Equal(t, expected, result, "Distance should return the number of elements from begin to mid")
}

func TestDistanceEmptyRange(t *testing.T) {
	vec := vector.NewVector[int](1, 2, 3, 4, 5)

	begin := vec.Begin()
	result := Distance[int](begin, begin)
	expected := uint(0)
	assert.Equal(t, expected, result, "Distance for an empty range should be 0")
}
