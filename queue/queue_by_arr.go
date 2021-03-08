package queue

import "fmt"

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
	if (q.tail+1)%q.cap == q.head {
		return false
	}
	q.queue[q.tail] = i
	fmt.Println(q.queue)
	q.tail++
	if q.tail >= q.cap {
		q.tail = 0
	}
	return true
}

func (q *QueueByArr) Pop() (interface{}, bool) {
	if q.head == q.tail {
		return 0, false
	}
	ret := q.queue[q.head]
	q.queue[q.head] = nil
	q.head++
	if q.head >= q.cap {
		q.head = 0
	}
	return ret, true
}

func (q *QueueByArr) Front() interface{} {
	return q.queue[q.head]
}

func (q *QueueByArr) Len() int {
	return len(q.queue)
}
