package stack

var stackArray []interface{}

// stackPush push to first index of array
func stackPush(n interface{}) {
	stackArray = append([]interface{}{n}, stackArray...)
}

// stackLength return length of stack
func stackLength() int {
	return len(stackArray)
}

// stackPeek return last input of stack
func stackPeek() interface{} {
	return stackArray[0]
}

// stackPop return last input of array and remove it in array
func stackPop() interface{} {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}

// stackEmpty check array is empty or not
func stackEmpty() bool {
	return len(stackArray) == 0
}
