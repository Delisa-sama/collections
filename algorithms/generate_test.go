package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestGenerate(t *testing.T) {
	dest := vector.NewVectorFromSlice(make([]int, 5))

	generatorState := 0
	myGenerator := func() int {
		generatorState++
		return generatorState
	}

	Generate(dest.Begin(), dest.End(), myGenerator)
	expected := vector.NewVectorFromSlice([]int{1, 2, 3, 4, 5})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}

func TestGenerateN(t *testing.T) {
	dest := vector.NewVectorFromSlice(make([]int, 5))
	GenerateN(dest.Begin(), 3, func() int {
		return 7
	})
	expected := vector.NewVectorFromSlice([]int{7, 7, 7, 0, 0})
	assert.True(t, Equals[int](dest.Begin(), expected.Begin()))
}
