package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	balance int
}

func (b *Bank) Deposit(amount int) { b.balance = b.balance + amount }

func (b *Bank) Balance() int { return b.balance }

func test() {
	// A
	var bank = Bank{0}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		bank.Deposit(100)
		wg.Done()
	}()
	// B
	go func() {
		bank.Deposit(200)
		wg.Done()
	}()
	wg.Wait()
	//fmt.Println(bank.Balance())
	if bank.Balance() != 300 {
		fmt.Println("最终余额不等于300，结果是：", bank.Balance())
	}
	// 如果是以下情形，最终的结果，有可能不是300
	// 当A在执行 balance + amount 的时候，B执行完了 balance = balance + amount
	// 然后A执行 balance = balance + amount
	// 这时候，最终的余额是100
}

func main() {

	for i := 0; i < 100000; i++ {
		test()
	}

	time.Sleep(5 * time.Second)
}
