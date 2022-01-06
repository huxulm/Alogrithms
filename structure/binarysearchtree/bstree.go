package binarysearchtree

import "errors"

// Preorder: root --> left --> right
// Inorder: left --> root --> right
// Postorder: left --> right --> root
// Levelorder

type Node struct {
	val         int
	left, right *Node
}

func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}

type BSTree struct {
	Root *Node
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// calculateDepth helper function for BSTree's depth()
func calculateDepth(n *Node, depth int) int {
	if n == nil {
		return depth
	}
	return Max(calculateDepth(n.left, depth+1), calculateDepth(n.right, depth+1))
}

// Insert a value in the BSTree
func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

func (bst *BSTree) Depth() int {
	return calculateDepth(bst.Root, 0)
}

func InOrderSuccessor(root *Node) *Node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func BstDelete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = BstDelete(root.left, val)
	} else if val > root.val {
		root.right = BstDelete(root.right, val)
	} else {
		// this is the node to delete
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			n := root.right
			d := InOrderSuccessor(n)
			d.left = root.left
			return root.right
		}
	}
	return root
}

// left --> root --> right
func inOrderRecursive(n *Node, traversal *[]int) {
	if n != nil {
		inOrderRecursive(n.left, traversal)
		*traversal = append(*traversal, n.val)
		inOrderRecursive(n.right, traversal)
	}
}

func InOrder(root *Node) []int {
	traversal := make([]int, 0)
	inOrderRecursive(root, &traversal)
	return traversal
}

func (bst *BSTree) Insert(val int) {
	bst.Root = Insert(bst.Root, val)
}

func (bst *BSTree) InOrder() []int {
	return InOrder(bst.Root)
}

// root --> left --> right
func preOrderRecursive(n *Node, traversal *[]int) {
	if n == nil {
		return
	}
	*traversal = append(*traversal, n.val)
	preOrderRecursive(n.left, traversal)
	preOrderRecursive(n.right, traversal)
}

func PreOrder(root *Node) []int {
	traversal := make([]int, 0)
	preOrderRecursive(root, &traversal)
	return traversal
}

func (bst *BSTree) PreOrder() []int {
	return PreOrder(bst.Root)
}

// left --> right --> root
func postOrderRecursive(n *Node, traversal *[]int) {
	if n == nil {
		return
	}
	postOrderRecursive(n.left, traversal)
	postOrderRecursive(n.right, traversal)
	*traversal = append(*traversal, n.val)
}

func PostOrder(root *Node) []int {
	traversal := make([]int, 0)
	postOrderRecursive(root, &traversal)
	return traversal
}

func (bst *BSTree) PostOrder() []int {
	return PostOrder(bst.Root)
}

func least(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return least(node.left)
}

func removeNode(n *Node, val int) (*Node, error) {
	if n == nil {
		return nil, errors.New("Not found")
	}

	if val > n.val {
		right, err := removeNode(n.right, val)
		if err != nil {
			return nil, err
		}
		n.right = right
	} else if val < n.val {
		left, err := removeNode(n.left, val)
		if err != nil {
			return nil, err
		}
		n.left = left
	} else {
		if n.left != nil && n.right != nil {
			// has two childs
			successor := least(n.right)
			value := successor.val
			// remove successor
			right, err := removeNode(n.right, value)
			if err != nil {
				return nil, err
			}
			n.right = right
			n.val = value
		} else if n.left != nil || n.right != nil {
			// has 1 child
			if n.left != nil {
				n = n.left
			} else {
				n = n.right
			}
		} else {
			// no child
			n = nil
		}
	}
	return n, nil
}

func (bst *BSTree) Remove(val int) (*Node, error) {
	return removeNode(bst.Root, val)
}
