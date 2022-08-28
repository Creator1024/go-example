
package main

import "fmt"

func intPrinter(x int){
	fmt.Println(x)
}

func main() {
	var f1 func(int) int
	fmt.Println(f1 == nil)  // true

	f2 := intPrinter
	f2(100)  // 100

	fmt.Printf("%T\n", f2) // func(int)
}