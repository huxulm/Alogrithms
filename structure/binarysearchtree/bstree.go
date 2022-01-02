package binarysearchtree

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
