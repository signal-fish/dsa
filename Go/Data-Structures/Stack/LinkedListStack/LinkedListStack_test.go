package LinkedListStack

import (
	"fmt"
	"testing"
)

func TestLinkedListStackPush(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Get())
	t.Log(s.Pop())
	t.Log(s.Pop())
	fmt.Println(s.Get())
	t.Log(s.Peek())
}
