package binarysearchtree_test

import (
	"reflect"
	"testing"

	bst "github.com/huxulm/algorithms/structure/binarysearchtree"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &bst.BSTree{}
	/* Let us create following BST
				50
			/      \
		30       70
		/  \     /  \
	20   40   60   80
	*/
	tree.Insert(1)
	tree.Insert(6)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(7)

	t.Run("Test in-order traversal", func(t *testing.T) {
		expect := []int{1, 2, 3, 4, 5, 6, 7, 8}
		result := bst.InOrder(tree.Root)
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Expect: %v but got: %v", expect, result)
		}
	})

}

func TestPreorder(t *testing.T) {
	tree := bst.BSTree{}
	nums := []int{80, 100, 70, 85, 95, 105, 90}

	for _, n := range nums {
		tree.Insert(n)
	}

	result := tree.PreOrder()
	expect := []int{80, 70, 100, 85, 95, 90, 105}
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expect: %+v, but got: %+v", expect, result)
	}
}
