// Stack Linked-List with standard library (Container/List)
// description: based on `geeksforgeeks` description Stack is a linear data structure which follows a particular order in which the operations are performed.
//	The order may be LIFO(Last In First Out) or FILO(First In Last Out).
// details:
// 	Stack Data Structure : https://www.geeksforgeeks.org/stack-data-structure-introduction-program/
// 	Stack (abstract data type) : https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
// author [huxulm](https://github.com/huxulm)
// see stacklinkedlist.go, stackarray.go, stack_test.go
package stack

import "container/list"

type SList struct {
	stack *list.List
}

func (sl *SList) push(v interface{}) {
	sl.stack.PushFront(v)
}

func (sl *SList) pop() interface{} {
	f := sl.stack.Front()
	sl.stack.Remove(f)
	return f.Value
}

func (sl *SList) peek() interface{} {
	return sl.stack.Front().Value
}

func (sl *SList) len() int {
	return sl.stack.Len()
}

func (sl *SList) isEmpty() bool {
	return sl.stack.Len() == 0
}
