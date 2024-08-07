package algorithms

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Delisa-sama/collections/sequence/vector"
)

func TestSort(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{5, 2, 9, 1, 5, 6})
	Sort[int](vec.Begin(), vec.End())
	expected := vector.NewVectorFromSlice([]int{1, 2, 5, 5, 6, 9})
	assert.Truef(t, Equals[int](vec.Begin(), expected.Begin()), "expected: %+v, got: %+v", expected, vec)

	vec = vector.NewVectorFromSlice(make([]int, 256))
	init := 256
	Generate(vec.Begin(), vec.End(), func() int {
		init--
		return init
	})
	Sort[int](vec.Begin(), vec.End())
	for it := vec.Begin(); !it.Equals(vec.End()); it.Next() {
		assert.Equal(t, init, it.Value())
		init++
	}
}

func TestSortC(t *testing.T) {
	vec := vector.NewVectorFromSlice([]int{5, 2, 9, 1, 5, 6})
	SortC[int](vec.Begin(), vec.End(), func(x, y int) bool {
		return x > y
	})
	expected := vector.NewVectorFromSlice([]int{9, 6, 5, 5, 2, 1})
	assert.Truef(t, Equals[int](vec.Begin(), expected.Begin()), "expected: %+v, got: %+v", expected, vec)
}

const sortBenchmarkLength = 10_000

func makeRandomIntSlice(n int) []int {
	numbers := make([]int, n)
	for i := range n {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		vec := vector.NewVectorFromSlice(makeRandomIntSlice(sortBenchmarkLength))

		b.StartTimer()
		Sort(vec.Begin(), vec.End())
	}
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := makeRandomIntSlice(sortBenchmarkLength)

		b.StartTimer()
		slices.Sort(s)
	}
}
