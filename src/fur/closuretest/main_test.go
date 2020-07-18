package closuretest

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"testing"
	"time"
)

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func Adder() func(int) int {
	var x int
	return func(n int) int {
		x += n
		return x
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if !strings.HasSuffix(filename, suffix) {
			return filename + suffix
		}
		return filename
	}

}

// 返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数

func TestMain(t *testing.T) {
	t.Log("Start..")
	// ------------
	fmt.Println(f()) // output 2

	// ------------
	var add = Adder()
	fmt.Println(add(1))
	fmt.Println(add(10))
	fmt.Println(add(100))

	// ------------
	suffixMp3 := MakeAddSuffix(".mp3")
	fmt.Println(suffixMp3("hello"))

	// 利用闭包调试
	// var where = log.Print
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	// some code
	where()
	// some more code
	where()

	// 计算函数执行时间
	start := time.Now()
	fmt.Println(suffixMp3("hello"))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println("cost time:", delta)
}
