package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestFill(t *testing.T) {
	type testCase struct {
		name  string
		vec   *vector.Vector[int]
		value int
	}
	tests := []testCase{
		{
			name:  "fill ten integers",
			vec:   vector.NewVectorFromSlice(make([]int, 10)),
			value: 765,
		},
		{
			name:  "fill empty vector",
			vec:   vector.NewVector[int](),
			value: 765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fill(tt.vec.Begin(), tt.vec.End(), tt.value)
			assert.True(t, AllOf(tt.vec.Begin(), tt.vec.End(), func(value int) bool {
				return tt.value == value
			}))
		})
	}
}
