package ArrayStack

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (as *ArrayStack) IsEmpty() bool {
	if ok := as.top < 0; ok {
		return true
	}
	return false
}

func (as *ArrayStack) Push(value interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top += 1
	}
	if as.top > len(as.data)-1 {
		as.data = append(as.data, value)
	} else {
		as.data[as.top] = value
	}
}

func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	v := as.data[as.top]
	as.top -= 1
	return v
}

func (as *ArrayStack) Peek() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.data[as.top]
}

func (as *ArrayStack) Empty() {
	as.top = -1
}

func (as *ArrayStack) Get() []interface{} {
	return as.data[:as.top+1]
}
