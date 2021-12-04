package ArrayQueue

import "fmt"

type ArrayQueue struct {
	queues   []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (aq *ArrayQueue) EnQueue(value interface{}) bool {
	if aq.tail == aq.capacity {
		return false
	}
	aq.queues[aq.tail] = value
	aq.tail++
	return true
}

func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}
	value := aq.queues[aq.head]
	aq.head++
	return value
}

func (aq *ArrayQueue) Get() string {
	if aq.head == aq.tail {
		return "empty queue"
	}
	result := "head"
	for i := aq.head; i <= aq.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", aq.queues[i])
	}
	result += "<-tail"
	return result
}
