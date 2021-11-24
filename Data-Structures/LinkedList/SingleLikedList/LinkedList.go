package LikedList

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length uint
}

func NewNode(value interface{}) *Node {
	return &Node{nil, value}
}

func (node *Node) GetNext() *Node {
	return node.next
}

func (node *Node) GetValue() interface{} {
	return node.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

// Insert a new node after a certain node
func (list *LinkedList) InsertAfter(node *Node, value interface{}) bool {
	if node == nil {
		return false
	}
	newNode := NewNode(value)
	oldNext := node.next
	node.next = newNode
	newNode.next = oldNext
	list.length++
	return true
}

// Insert a new node before a certain node
func (list *LinkedList) InsertBefore(node *Node, value interface{}) bool {
	if node == nil || node == list.head {
		return false
	}
	current := list.head.next
	prev := list.head
	if current == nil {
		return false
	}
	for current != nil {
		if current == node {
			break
		}
		prev = current
		current = current.next
	}
	newNode := NewNode(value)
	prev.next = newNode
	newNode.next = current
	list.length++
	return true
}

func (list *LinkedList) InsertToHead(value interface{}) bool {
	return list.InsertAfter(list.head, value)
}

func (list *LinkedList) InsertToTail(value interface{}) bool {
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return list.InsertAfter(current, value)
}

func (list *LinkedList) FindByIndex(index uint) *Node {
	if index >= list.length {
		return nil
	}
	current := list.head.next
	var i uint
	for i = 0; i < index; i++ {
		current = current.next
	}
	return current
}

func (list *LinkedList) DeleteNode(node *Node) bool {
	if node == nil {
		return false
	}
	current := list.head.next
	prev := list.head
	if current == nil {
		return false
	}
	for current != nil {
		if current == node {
			break
		}
		prev = current
		current = current.next
	}
	prev.next = node.next
	node = nil
	list.length--
	return true
}

func (list *LinkedList) GetSize() uint {
	return list.length
}

func (list *LinkedList) GetItems() string {
	current := list.head.next
	format := "<head>->"
	for current != nil {
		format += fmt.Sprintf("%+v", current.GetValue())
		current = current.next
		format += "->"
	}
	format += "<nil>"
	return format
}
