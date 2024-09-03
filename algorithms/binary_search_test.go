package algorithms

import (
	"testing"

	"github.com/Delisa-sama/collections/sequence/vector"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	data := vector.NewVector(1, 3, 5, 7, 9)

	found := BinarySearch(data.Begin(), data.End(), 5)
	assert.True(t, found, "Элемент 5 должен быть найден в массиве")

	notFound := BinarySearch(data.Begin(), data.End(), 2)
	assert.False(t, notFound, "Элемент 2 не должен быть найден в массиве")
}
