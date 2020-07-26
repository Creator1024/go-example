package csptest

import (
	"testing"
	"time"
	"fmt"
)

func service() string{
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask()  {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask is done")
}


func TestService(t *testing.T){
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string{
	retCh := make(chan string)
	// retCh := make(chan string, 1)
	go func(){
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret		
		fmt.Println("service exited")	
		// 由于没有设置缓冲区，service exited只有在channel的另一端把ret拿走之后才执行（阻塞）
		// 定义缓存型channel，retCh := make(chan string,1 )，即可立即执行service exited
	}()
	// 定义好了channel之后，直接返回
	return retCh
}

func TestAsyncService(t *testing.T){
	retCh:= AsyncService()
	otherTask()
	fmt.Println(<- retCh)
}

