package CircularQueue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewCircularQueue(7)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	fmt.Println(q.Get())
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	fmt.Println(q.Get())
}
