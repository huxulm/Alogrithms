package stack

type MonotoneStack struct {
	data []int
}

func (ms *MonotoneStack) Len() int {
	return len(ms.data)
}

func (ms *MonotoneStack) Top() int {
	return ms.data[0]
}

func (ms *MonotoneStack) Push(v int) {
	for ms.Len() > 0 && ms.Top() < v {
		ms.Pop()
	}
	ms.data = append([]int{v}, ms.data...)
}

func (ms *MonotoneStack) Pop() {
	if ms.Len() > 0 {
		ms.data = ms.data[1:]
	}
}

func (ms *MonotoneStack) Data() []int {
	return ms.data
}
