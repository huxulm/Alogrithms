package linkedlist

type Singly struct {
	length int

	// Note that node here holds both Next an Prev node
	// however only Next node is used in singly methods.
	Head *Node
}

func NewSingly() *Singly {
	return &Singly{}
}

func (ll *Singly) AddAtBegain(val interface{}) {
	n := NewNode(val)
	n.Next, ll.Head = ll.Head, n
	ll.length++
}

func (ll *Singly) AddAtEnd(val interface{}) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head

	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = n
	ll.length++
}

func (ll *Singly) DelAtBegain() interface{} {
	if ll.Head == nil {
		return -1
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--

	return cur.Val
}

func (ll *Singly) DelAtEnd() interface{} {
	if ll.Head == nil {
		return -1
	}
	if ll.Head.Next == nil {
		return ll.DelAtBegain()
	}

	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retval := cur.Next.Val
	cur.Next = nil
	ll.length--
	return retval
}

func (ll *Singly) Count() int {
	return ll.length
}

// Reverse reverses the list
func (ll *Singly) Reverse() {

	var prev, next *Node
	cur := ll.Head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	ll.Head = prev
}
