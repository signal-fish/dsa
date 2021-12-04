package ArrayStack

import (
	"fmt"
	"testing"
)

func TestArrayStackPush(t *testing.T) {
	as := NewArrayStack()
	as.Push(1)
	as.Push(2)
	top := as.top
	if top != 1 {
		t.Fail()
	}
	v := as.Get()
	t.Log(v)
	fmt.Println("Test Push:", v)
}

func TestArrayStackPop(t *testing.T) {
	as := NewArrayStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	t.Log(as.Get())
	t.Log(as.Pop())
	t.Log(as.Pop())
	t.Log(as.Pop())
	t.Log(as.Pop())
}

func TestArrayStackPeek(t *testing.T) {
	as := NewArrayStack()
	as.Push(1)
	as.Push(2)
	as.Push(3)
	t.Log(as.Peek())
	as.Pop()
	t.Log(as.Peek())
	as.Pop()
	t.Log(as.Peek())
	as.Pop()
	t.Log(as.Peek())
	as.Pop()
}
