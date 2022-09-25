package main

import (
	"fmt"
	"go.uber.org/fx"
)

type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{B: b}
}

type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{c}
}

type C struct {
}

func NewC() *C {
	return &C{}
}

func normalCreatA() {
	b := NewB(NewC())
	a := NewA(b)
	_ = a
	PrintA(a)
}

func fxCreateA() {
	fx.New(
		// Provide传入的是一个返回指针的函数
		fx.Provide(NewA),
		fx.Provide(NewB),
		fx.Provide(NewC),
		fx.Invoke(PrintA),
	)
}

func main() {
	normalCreatA()
	fxCreateA()
}
func PrintA(a *A) {
	fmt.Println(*a)
}
