package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestMin(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{10, 20, 5, 40, 15})
	it := Min[int](vec.Begin(), vec.End())
	assert.Equal(t, 5, it.Value())
}

func TestMinC(t *testing.T) {
	vec := vector.NewVectorFromSlice([]string{"apple", "banana", "cherry"})
	cmp := func(a, b string) int {
		return len(a) - len(b)
	}
	it := MinC(vec.Begin(), vec.End(), cmp)
	assert.Equal(t, "apple", it.Value())
}

func TestMax(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{10, 20, 5, 40, 15})
	it := Max[int](vec.Begin(), vec.End())
	assert.Equal(t, 40, it.Value())
}

func TestMaxC(t *testing.T) {
	vec := vector.NewVectorFromSlice([]string{"apple", "banana", "cherry"})
	cmp := func(a, b string) int {
		return len(a) - len(b)
	}
	it := MaxC(vec.Begin(), vec.End(), cmp)
	assert.Equal(t, "banana", it.Value())
}

func TestMinMax(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{10, 20, 5, 40, 15})
	minIt, maxIt := MinMax[int](vec.Begin(), vec.End())
	assert.Equal(t, 5, minIt.Value())
	assert.Equal(t, 40, maxIt.Value())
}

func TestMinMaxC(t *testing.T) {
	vec := vector.NewVectorFromSlice([]string{"apple", "banana", "cherry"})
	cmp := func(a, b string) int {
		return len(a) - len(b)
	}
	minIt, maxIt := MinMaxC(vec.Begin(), vec.End(), cmp)
	assert.Equal(t, "apple", minIt.Value())
	assert.Equal(t, "banana", maxIt.Value())
}
