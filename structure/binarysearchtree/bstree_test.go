package binarysearchtree_test

import (
	"reflect"
	"testing"

	bst "github.com/huxulm/alogrithms/structure/binarysearchtree"
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
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(20)
	tree.Insert(80)
	tree.Insert(60)
	tree.Insert(70)

	expect := []int{20, 30, 40, 50, 60, 70, 80}
	result := bst.InOrder(tree.Root)
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expect: %v but got: %v", expect, result)
	}
}
