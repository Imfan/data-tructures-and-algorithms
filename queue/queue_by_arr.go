package queue

type QueueByArr struct {
	queue           []interface{}
	head, tail, cap int
}

// num传多少都行，切片会自动扩展
func Init(num int) *QueueByArr {
	return &QueueByArr{
		queue: make([]interface{}, num),
		cap:   num,
	}
}
func (q *QueueByArr) Push(i interface{}) bool {
	if len(q.queue) >= q.cap {
		return false
	}
	q.queue[q.tail] = i
	q.tail++
	if q.tail >= q.cap {
		q.tail = 0
	}
	return true
}

func (q *QueueByArr) Pop() (interface{}, bool) {
	if len(q.queue) <= 0 {
		return 0, false
	}
	ret := q.queue[q.head]
	q.head++
	return ret, true
}

func (q *QueueByArr) Front() interface{} {
	return q.queue[q.head]
}

func (q *QueueByArr) Len() int {
	return len(q.queue)
}
