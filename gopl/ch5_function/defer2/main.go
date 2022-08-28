package main

import (
	"fmt"
	"time"
)

func doSomething() {
	defer printTimeUsage("doSomething func")()
	time.Sleep(5 * time.Second)
}

func printTimeUsage(funcName string) func() {
	start := time.Now()
	fmt.Printf("Func %s start.\n", funcName)
	return func() {
		fmt.Printf("Func %s end, time usage: %s.\n", funcName, time.Since(start))
	}
}

func main() {
	doSomething()
}
