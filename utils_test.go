package containers

import (
	"fmt"
	"testing"

	"github.com/Delisa-sama/collections/adapters/stack"
	"github.com/Delisa-sama/collections/associative/bst"
	"github.com/Delisa-sama/collections/associative/set"
	"github.com/Delisa-sama/collections/comparator"
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
		if EqualsByIterators(binaryTree.InOrderIterator(), l3.Begin(), comparator.DefaultComparator[int]()) {
			t.Fatalf("wtf")
		}
		Print(binaryTree.InOrderIterator())
		fmt.Println()
	})
}
