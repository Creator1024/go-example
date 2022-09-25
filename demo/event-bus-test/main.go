package main

import (
	"fmt"
	"github.com/asaskevich/EventBus"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	bus := EventBus.New()
	_ = bus.Subscribe("main:calculator", calculator)
	_ = bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	bus.Publish("main:calculator", 220, 40)
	//_ = bus.Unsubscribe("main:calculator", calculator)
}
