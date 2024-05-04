package queue

type Queue struct {
	queue []int
}

func (Q *Queue) Init() {
	Q.queue = []int{}
}

func (Q *Queue) Push(value int) {
	Q.queue = append(Q.queue, value)
}

func (Q *Queue) Pop() {
	Q.queue = Q.queue[1:]
}

func (Q *Queue) First() int {
	return Q.queue[0]
}

func (Q *Queue) Size() int {
	return len(Q.queue)
}
