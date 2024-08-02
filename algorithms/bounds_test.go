package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestLowerBound(t *testing.T) {
	data := vector.NewVectorFromSlice([]int{1, 2, 2, 4, 5, 6, 8})
	begin := data.Begin()
	end := data.End()

	it := LowerBound(begin, end, 4)
	assert.Equal(t, 4, it.Value())

	it = LowerBound(begin, end, 3)
	assert.Equal(t, 4, it.Value())

	it = LowerBound(begin, end, 0)
	assert.Equal(t, 1, it.Value())

	it = LowerBound(begin, end, 10)
	assert.True(t, it.Equals(end))
}

func TestUpperBound(t *testing.T) {
	data := vector.NewVectorFromSlice([]int{1, 2, 2, 4, 5, 6, 8})
	begin := data.Begin()
	end := data.End()

	it := UpperBound(begin, end, 4)
	assert.Equal(t, 5, it.Value())

	it = UpperBound(begin, end, 3)
	assert.Equal(t, 4, it.Value())

	it = UpperBound(begin, end, 0)
	assert.Equal(t, 1, it.Value())

	it = UpperBound(begin, end, 10)
	assert.True(t, it.Equals(end))
}
