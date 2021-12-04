package LinkedListStack

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type LinkedListStack struct {
	topNode *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (ls *LinkedListStack) IsEmpty() bool {
	return ls.topNode == nil
}

func (ls *LinkedListStack) Push(value interface{}) {
	ls.topNode = &Node{next: ls.topNode, value: value}
}

func (ls *LinkedListStack) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	value := ls.topNode.value
	ls.topNode = ls.topNode.next
	return value
}

func (ls *LinkedListStack) Peek() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	return ls.topNode.value
}

func (ls *LinkedListStack) Empty() {
	ls.topNode = nil
}

func (ls *LinkedListStack) Get() string {
	if ls.IsEmpty() {
		fmt.Println("empty stack")
	}
	result := "["
	current := ls.topNode
	for current != nil {
		value := fmt.Sprintf("%+v", current.value)
		result += value
		if current.next != nil {
			result += "->"
		}
		current = current.next
	}
	result += "(top)"
	return result
}
