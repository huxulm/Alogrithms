package stack

import (
	"reflect"
	"testing"
)

// TestStackLinkedList for testing Stack with LinkedList
func TestStackLinkedList(t *testing.T) {
	stack := Stack{}

	t.Run("Stack Push", func(t *testing.T) {
		stack.push(2)
		stack.push(3)
		result := stack.show()
		expected := []interface{}{3, 2}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Stack Push is not work we expected %v but got %v", expected, result)
		}
	})

	t.Run("Stack Pop", func(t *testing.T) {
		pop := stack.pop()
		if stack.len() == 2 && pop != 3 {
			t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
		}
	})

	t.Run("Stack Peek", func(t *testing.T) {
		stack.push(2)
		stack.push(83)
		if stack.peek() != 83 {
			t.Errorf("Stack Peek is not worked as we expected: %v but got: %v", 83, stackPeek())
		}
	})

	t.Run("Stack Length", func(t *testing.T) {
		if stack.len() != 3 {
			t.Errorf("Stack Length is not work we expected %v but got %v", 3, stack.len())
		}
	})

	t.Run("Stack Empty", func(t *testing.T) {
		if stack.isEmpty() {
			t.Errorf("Stack Empty is not work we expected %v but got %v", false, stack.isEmpty())
		}
		stack.pop()
		stack.pop()
		stack.pop()
		if stack.isEmpty() == false {
			t.Errorf("Stack Empty is not work we expected %v but got %v", true, stack.isEmpty())
		}
	})
}

func TestStackArray(t *testing.T) {
	t.Run("Stack With Array", func(t *testing.T) {
		stackPush(2)
		stackPush(3)

		t.Run("Stack Push", func(t *testing.T) {
			if !reflect.DeepEqual([]interface{}{3, 2}, stackArray) {
				t.Errorf("Stack Push is not work we expected %v but got %v", []interface{}{3, 2}, stackArray)
			}
		})

		pop := stackPop()

		t.Run("Stack Pop", func(t *testing.T) {
			if stackLength() == 2 && pop != 3 {
				t.Errorf("Stack Pop is not work we expected %v but got %v", 3, pop)
			}
		})

		stackPush(2)
		stackPush(83)

		t.Run("Stack Peek", func(t *testing.T) {
			if stackPeek() != 83 {
				t.Errorf("Stack Peek is not worked as we expected: %v but got: %v", 83, stackPeek())
			}
		})

		t.Run("Stack Length", func(t *testing.T) {
			if stackLength() != 3 {
				t.Errorf("Stack Length is not work we expected %v but got %v", 3, stackLength())
			}
		})

		t.Run("Stack Empty", func(t *testing.T) {
			if stackEmpty() {
				t.Errorf("Stack Empty is not work we expected %v but got %v", false, stackEmpty())
			}
			stackPop()
			stackPop()
			stackPop()
			if stackEmpty() == false {
				t.Errorf("Stack Empty is not work we expected %v but got %v", true, stackEmpty())
			}
		})
	})
}
