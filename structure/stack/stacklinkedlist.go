package stack

// Node structure
type Node struct {
	Val  interface{}
	Next *Node
}

// Stack just has top of node with length
type Stack struct {
	top    *Node
	length int
}

func (ll *Stack) push(n interface{}) {
	newStack := &Node{}
	newStack.Val = n
	newStack.Next = ll.top
	ll.top = newStack
	ll.length++
}

// pop remove last item as first output
func (ll *Stack) pop() interface{} {
	if ll.length == 0 {
		return nil
	}
	result := ll.top.Val
	if ll.top.Next == nil {
		ll.top = nil
	} else {
		ll.top.Val, ll.top.Next = ll.top.Next.Val, ll.top.Next.Next
	}
	ll.length--
	return result
}

// isEmpty to check our stack is empty or not
func (ll *Stack) isEmpty() bool {
	return ll.length == 0
}

// len returns length of our stack
func (ll *Stack) len() int {
	return ll.length
}

// peek returns last input value
func (ll *Stack) peek() interface{} {
	if ll.length == 0 {
		return nil
	}
	return ll.top.Val
}

// show all value as an interface
func (ll *Stack) show() (in []interface{}) {
	current := ll.top
	for current != nil {
		in = append(in, current.Val)
		current = current.Next
	}
	return
}
