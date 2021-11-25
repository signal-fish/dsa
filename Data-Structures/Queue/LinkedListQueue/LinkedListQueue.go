package LinkedListQueue

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedListQueue struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (lq *LinkedListQueue) EnQueue(value interface{}) {
	node := &Node{value, nil}
	if lq.tail == nil {
		lq.tail = node
		lq.head = node
	} else {
		lq.tail.next = node
		lq.head = node
	}
	lq.length++
}

func (lq *LinkedListQueue) DeQueue() interface{} {
	if lq.head == nil {
		return nil
	}
	value := lq.head.value
	lq.head = lq.head.next
	lq.length++
	return value
}

func (lq *LinkedListQueue) Get() string {
	if lq.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for node := lq.head; node != nil; node = node.next {
		result += fmt.Sprintf("<-%+v", node.value)
	}
	result += "<-tail"
	return result
}
