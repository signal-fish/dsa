package Factoril

import "fmt"

type Fac struct {
	values map[int]int
}

func NewFactorial(n int) *Fac {
	return &Fac{
		make(map[int]int, n),
	}
}

func (fac *Fac) Factorial(n int) int {
	if fac.values[n] != 0 {
		return fac.values[n]
	}
	if n <= 1 {
		fac.values[n] = 1
		return 1
	} else {
		res := n * fac.Factorial(n-1)
		fac.values[n] = res
		return res
	}
}

func (fac *Fac) Get(n int) {
	fmt.Println(fac.values[n])
}
