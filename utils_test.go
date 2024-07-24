package containers

import (
	"testing"

	"github.com/Delisa-sama/collections/sequence/forwardlist"
	"github.com/Delisa-sama/collections/sequence/list"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func Test_EqualsByIterators(t *testing.T) {
	t.Run("list, forward list and vector compatibility", func(t *testing.T) {
		l := forwardlist.NewForwardList[int](1, 2, 3)
		if l.Size() != 3 {
			t.Fatalf("bad list size, got (%d), expected (%d)", l.Size(), 3)
		}

		l2 := list.NewList[int](1, 2, 3)
		if !EqualsByIterators(l.Begin(), l2.Begin(), DefaultComparator[int]()) {
			t.Fatalf("lists not equals")
		}

		l3 := vector.NewVector(1, 2, 3)
		if !EqualsByIterators(l.Begin(), l3.Begin(), DefaultComparator[int]()) {
			t.Fatalf("list and vector not equals")
		}

		vec := vector.NewVector(1, 2, 3)
		vecBegin := vec.Begin()
		vec.Append(4)
		if EqualsByIterators(vecBegin, l3.Begin(), DefaultComparator[int]()) {
			t.Fatalf("iterator not invalidated after vector was modified")
		}
	})
}
