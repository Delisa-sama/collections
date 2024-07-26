package containers

import (
	"fmt"
	"testing"

	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/algorithms"
	"github.com/Delisa-sama/collections/associative/bst"
	"github.com/Delisa-sama/collections/associative/set"
	"github.com/Delisa-sama/collections/comparator"
	"github.com/Delisa-sama/collections/pair"
	"github.com/Delisa-sama/collections/sequence/forwardlist"
	"github.com/Delisa-sama/collections/sequence/list"
	"github.com/Delisa-sama/collections/sequence/vector"
)

func Test_Examples(t *testing.T) {
	t.Run("examples", func(t *testing.T) {
		l := forwardlist.NewForwardList[int](1, 2, 3)
		if l.Size() != 3 {
			t.Fatalf("bad list size, got (%d), expected (%d)", l.Size(), 3)
		}
		algorithms.ForEach(l.Begin(), l.End(), func(i int) {
			fmt.Println(i)
		})
		fmt.Println()

		l2 := list.NewList[int](1, 2, 3)
		if !algorithms.Equals(l.Begin(), l2.Begin()) {
			t.Fatalf("lists not equals")
		}
		algorithms.ForEach(l2.Begin(), l2.End(), func(i int) {
			fmt.Println(i)
		})
		fmt.Println()

		l3 := vector.NewVector(1, 2, 3)
		if !algorithms.Equals(l.Begin(), l3.Begin()) {
			t.Fatalf("list and vector not equals")
		}
		algorithms.ForEach(l3.Begin(), l3.End(), func(i int) {
			fmt.Println(i)
		})
		fmt.Println()

		s := set.NewSet(1, 2, 3, 3)
		if !algorithms.Equals(l.Begin(), s.Begin()) {
			t.Fatalf("list and set not equals")
		}

		algorithms.ForEach(s.Begin(), s.End(), func(i int) {
			fmt.Println(i)
		})
		fmt.Println()

		vec := vector.NewVector(1, 2, 3)
		vecBegin := vec.Begin()
		vec.PushBack(4)
		if algorithms.Equals[int](vecBegin, l3.Begin()) {
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
		if algorithms.Equals(binaryTree.InOrderBegin(), l3.Begin()) {
			t.Fatalf("wtf")
		}
		algorithms.ForEach(binaryTree.InOrderBegin(), binaryTree.InOrderEnd(), func(i int) {
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
		algorithms.ForEach(kvBST.InOrderBegin(), kvBST.InOrderEnd(), func(p pair.Pair[int, string]) {
			fmt.Println(p)
		})
		fmt.Println()

		fmt.Println(kvBST.Find(pair.NewPair(8, "8")))
		kvBST.Delete(pair.NewPair(8, "8"))
		fmt.Println(kvBST.Find(pair.NewPair(8, "8")))
		fmt.Println()

		algorithms.ForEach(kvBST.PreOrderBegin(), kvBST.PreOrderEnd(), func(p pair.Pair[int, string]) {
			fmt.Println(p)
		})
		fmt.Println()

		algorithms.ForEach(kvBST.PostOrderBegin(), kvBST.PostOrderEnd(), func(p pair.Pair[int, string]) {
			fmt.Println(p)
		})
		fmt.Println()

		sl1 := []int{3, 2, 1, 4, 6}
		sl2 := []int{1, 2, 3, 5, 4}

		// O(n + m)
		bst1 := bst.NewBST(comparator.DefaultComparator[int](), sl1...)
		bst2 := bst.NewBST(comparator.DefaultComparator[int](), sl2...)
		if algorithms.EqualsRanges[int](
			bst1.InOrderBegin(), bst1.InOrderEnd(),
			bst2.InOrderBegin(), bst2.InOrderEnd(),
		) {
			fmt.Println("equal")
		} else {
			fmt.Println("not equal")
		}
		fmt.Println()

		count := algorithms.CountIf(bst1.InOrderBegin(), bst1.InOrderEnd(), func(value int) bool {
			return bst2.Find(value) != nil
		})
		fmt.Println(count)

		vec2 := vector.NewVectorFromSlice([]int{1, 2, 3, 4, 6})
		if algorithms.EqualsRanges[int](
			bst1.InOrderBegin(), bst1.InOrderEnd(),
			vec2.Begin(), vec2.End(),
		) {
			fmt.Println("equal")
		} else {
			fmt.Println("not equal")
		}
		fmt.Println()
	})
}
