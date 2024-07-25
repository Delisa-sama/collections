package containers

import (
	"fmt"
	"testing"

	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/associative/bst"
	"github.com/Delisa-sama/collections/associative/set"
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/pair"
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
		if !EqualsByIterators(l.Begin(), l2.Begin(), comparator.DefaultComparator[int]()) {
			t.Fatalf("lists not equals")
		}

		l3 := vector.NewVector(1, 2, 3)
		if !EqualsByIterators(l.Begin(), l3.Begin(), comparator.DefaultComparator[int]()) {
			t.Fatalf("list and vector not equals")
		}

		s := set.NewSet(1, 2, 3, 3)
		if !EqualsByIterators(l.Begin(), s.Begin(), comparator.DefaultComparator[int]()) {
			t.Fatalf("list and set not equals")
		}

		vec := vector.NewVector(1, 2, 3)
		vecBegin := vec.Begin()
		vec.PushBack(4)
		if EqualsByIterators(vecBegin, l3.Begin(), comparator.DefaultComparator[int]()) {
			t.Fatalf("iterator not invalidated after vector was modified")
		}

		vecStack := stack.NewStack(vector.NewVector[int], 1, 2, 3)
		for ; !vecStack.IsEmpty(); vecStack.Pop() {
			fmt.Println(vecStack.Top())
		}
		fmt.Println()

		listStack := stack.NewStack(list.NewList[int], 1, 2, 3)
		for ; !listStack.IsEmpty(); listStack.Pop() {
			fmt.Println(listStack.Top())
		}
		fmt.Println()

		fwlistStack := stack.NewStack(forwardlist.NewForwardList[int], 1, 2, 3)
		for ; !fwlistStack.IsEmpty(); fwlistStack.Pop() {
			fmt.Println(fwlistStack.Top())
		}
		fmt.Println()

		binaryTree := bst.NewBST(comparator.DefaultComparator[int](), 4, 3, 5, 1, 0, 2, 6, 8, 7)
		// In-order обход
		if EqualsByIterators(binaryTree.InOrderIteratorBegin(), l3.Begin(), comparator.DefaultComparator[int]()) {
			t.Fatalf("wtf")
		}
		ForEach(binaryTree.InOrderIteratorBegin(), binaryTree.InOrderIteratorEnd(), func(i int) {
			fmt.Println(i)
		})
		fmt.Println()

		comparePairByFirstInt := func(a, b pair.Pair[int, string]) int {
			return comparator.DefaultComparator[int]()(a.First, b.First)
		}
		kvBST := bst.NewBST[pair.Pair[int, string]](
			comparePairByFirstInt,
			pair.NewPair(4, "some"),
			pair.NewPair(3, "aaaa"),
			pair.NewPair(5, "bbb"),
			pair.NewPair(1, "ccccc"),
			pair.NewPair(0, "ddd"),
			pair.NewPair(2, "ff"),
			pair.NewPair(6, "6"),
			pair.NewPair(8, "8"),
			pair.NewPair(7, "7"),
		)
		ForEach(kvBST.InOrderIteratorBegin(), kvBST.InOrderIteratorEnd(), func(p pair.Pair[int, string]) {
			fmt.Println(p)
		})
		fmt.Println()
		fmt.Println(kvBST.Find(pair.NewPair(8, "8")))
		kvBST.Delete(pair.NewPair(8, "8"))
		fmt.Println(kvBST.Find(pair.NewPair(8, "8")))
	})
}
