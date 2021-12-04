package CircularQueue

import "fmt"

type CircularQueue struct {
	queues   []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

func (cq *CircularQueue) EnQueue(value interface{}) bool {
	if ok := cq.head == (cq.tail+1)%cq.capacity; ok {
		return false
	}
	cq.queues[cq.tail] = value
	cq.tail = (cq.tail + 1) % cq.capacity
	return true
}

func (cq *CircularQueue) DeQueue() interface{} {
	if ok := cq.head == cq.tail; ok {
		return nil
	}
	value := cq.queues[cq.head]
	cq.head = (cq.head + 1) % cq.capacity
	return value
}

func (cq *CircularQueue) Get() string {
	if ok := cq.head == cq.tail; ok {
		return "empty queue"
	}
	result := "head"
	for i := cq.head; i != cq.tail; i = (i + 1) % cq.capacity {
		result += fmt.Sprintf("<-%+v", cq.queues[i])
	}
	result += "<-tail"
	return result
}
