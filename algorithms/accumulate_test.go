package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/interfaces"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestAccumulate(t *testing.T) {
	type args[T interfaces.Numeric] struct {
		vec  *vector.Vector[T]
		init T
	}
	type testCase[T interfaces.Numeric] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Accumulate should return the sum of all elements",
			args: args[int]{
				vec:  vector.NewVector(1, 2, 3, 4, 5),
				init: 10,
			},
			want: 25,
		},
		{
			name: "Accumulate should return the sum of all elements plus init value",
			args: args[int]{
				vec:  vector.NewVector(1, 2, 3, 4, 5),
				init: 0,
			},
			want: 15,
		},
		{
			name: "Accumulate should return the sum of all elements plus init value",
			args: args[int]{
				vec:  vector.NewVector(1, 2, 3, 4, 5),
				init: 0,
			},
			want: 15,
		},
		{
			name: "Accumulate on an empty vector should return the init value",
			args: args[int]{
				vec:  vector.NewVector[int](),
				init: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Accumulate(tt.args.vec.Begin(), tt.args.vec.End(), tt.args.init)
			assert.Equal(t, tt.want, got)
		})
	}
}
