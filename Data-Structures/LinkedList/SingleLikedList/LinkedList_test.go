package LikedList

import (
	"fmt"
	"testing"
)

func TestInsertToHead(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToHead(i + 1)
	}
	fmt.Println("Insert To Head:", list.GetItems())
}

func TestInsertToTail(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i + 1)
	}
	fmt.Println("Insert To Tail:", list.GetItems())
}

func TestFindByIndex(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i + 1)
	}
	t.Log(list.FindByIndex(0))
	t.Log(list.FindByIndex(9))
	t.Log(list.FindByIndex(5))
	t.Log(list.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	list := NewLinkedList()
	for i := 0; i < 10; i++ {
		list.InsertToTail(i + 1)
	}
	fmt.Println("Before Delete:", list.GetItems())

	t.Log(list.DeleteNode(list.head.next)) // delete 1
	fmt.Println("After Delete 1:", list.GetItems())

	t.Log(list.DeleteNode(list.head.next.next)) // delete 3
	fmt.Println("After Delete 3:", list.GetItems())
}
