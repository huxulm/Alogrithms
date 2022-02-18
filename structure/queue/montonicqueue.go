package queue

type MontonicQueue struct {
	data []int
}

func (mq *MontonicQueue) Push(val int) {
	for sz := len(mq.data); sz > 0 && mq.data[sz-1] < val; sz = len(mq.data) {
		mq.data = mq.data[:sz-1]
	}
	mq.data = append(mq.data, val)
}

func (mq *MontonicQueue) Max() int {
	return mq.data[0]
}

func (mq *MontonicQueue) Pop(val int) {
	if mq.data[0] == val {
		mq.data = mq.data[1:]
	}
}
