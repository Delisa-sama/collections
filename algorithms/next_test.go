package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestNextBound(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{10, 20, 30, 40, 50})

	vecIt := vec.Begin()
	vecBound := vec.Begin()
	Advance[int](vecBound, 3)
	NextBound(vecIt, vecBound)
	assert.Equal(t, uint(3), vecIt.Index())

	NextBound(vecIt, vec.End())
	assert.Equal(t, uint(5), vecIt.Index())
}
