package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type Girl struct {
	Name string
	Age  int
}

func NewGirl() *Girl {
	return &Girl{
		Name: "Girl",
		Age:  18,
	}
}

type Gay struct {
	Girl *Girl
}

func NewGay(girl *Girl) *Gay {
	return &Gay{girl}
}
func main() {
	invoke := func(gay *Gay) {
		fmt.Println(gay.Girl) // &{Girl 18}
	}
	girl := NewGirl() //直接提供对象
	app := fx.New(
		// Provide将被依赖的对象的构造函数传进去，传进去的函数必须是个待返回值的函数指针
		fx.Provide(NewGay),
		//fx.Provide(NewGirl),
		// 直接提供被依赖的对象
		fx.Supply(girl),
		// Invoke将函数依赖的对象作为参数传进函数然后调用函数
		fx.Invoke(invoke),
	)
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
