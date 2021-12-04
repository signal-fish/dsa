package Fibonacci

import "fmt"

type Fibs struct {
	values map[int]int
}

func NewFibs(n int) *Fibs {
	return &Fibs{
		make(map[int]int, n),
	}
}

func (fibs *Fibs) Fibonacci(n int) int {
	if fibs.values[n] != 0 {
		return fibs.values[n]
	}
	if n <= 1 {
		fibs.values[1] = 1
		return 1
	} else if n == 2 {
		fibs.values[2] = 1
		return 1
	} else {
		res := fibs.Fibonacci(n-1) + fibs.Fibonacci(n-2)
		fibs.values[n] = res
		return res
	}
}

func (fibs *Fibs) Get(n int) {
	fmt.Print(fibs.values[n], " ")
}
