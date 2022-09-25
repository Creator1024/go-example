package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello() {
	for i := 0; i < 10000; i++ {
		fmt.Println("hello", i)
	}
}

func main() {

	// Example1
	//for i := 0; i < 10000; i++ {
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}
	// 此时部分i无法被打印出来，因为main已结束

	// Example2
	//for i := 0; i < 10000; i++ {
	//	go func(i int) {
	//		fmt.Println(i)
	//	}(i)
	//}
	// time.Sleep(1 * time.Second)
	// i都能打印出来（不是必定），因为main等待的时间足够goroutine去执行

	// Example3
	//go printHello()
	//time.Sleep(2 * time.Millisecond)

	// Example4
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
