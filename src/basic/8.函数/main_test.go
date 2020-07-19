package mainfn

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

type IntConv func(op int) int

func timeCost(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time Cost: ", time.Since(start).Seconds())
		return ret
	}
}

func timeSleepOne(op int) int {
	time.Sleep(time.Second * 1)
	return op + 1
}

// 可变参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret

}

func TestFunc(t *testing.T) {
	t.Log(returnMultiValues())
	t.Log(timeCost(timeSleepOne)(1))
	tsSF := timeCost(timeSleepOne)
	t.Log(tsSF(1))
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}
