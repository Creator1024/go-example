package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a int = 1
		b int = 1
	)
	for i := 0; i < 5; i++ {
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(b)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
