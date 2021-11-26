package Factoril

import "testing"

func TestFactorial(t *testing.T) {
	fac := NewFactorial(10)
	for i := 1; i < 15; i++ {
		fac.Factorial(i)
		fac.Get(i)
	}
}
