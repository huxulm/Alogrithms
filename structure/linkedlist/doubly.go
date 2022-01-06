package linkedlist

type Doubly struct {
	Head *Node
}

func NewDoubly() *Doubly {
	return &Doubly{nil}
}

func (ll *Doubly) AddAtBegain(val interface{}) {
	n := NewNode(val)
	n.Next = ll.Head

	if ll.Head != nil {
		ll.Head.Prev = n
	}

	ll.Head = n
}

func (ll *Doubly) AddAtEnd(val interface{}) {
	cur := ll.Head
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		return
	}

	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = n
	n.Prev = cur
}

func (ll *Doubly) Reverse() {
	var prev, next *Node
	cur := ll.Head

	for cur != nil {
		next = cur.Next
		next.Next = prev
		cur.Prev = next
		prev = cur
		cur = next
	}

	ll.Head = prev
}
