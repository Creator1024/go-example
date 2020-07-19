package funcparamtest

import (
	"fmt"
	"testing"
)

func AddTen(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func callback(n int, f func(int, int)) {
	f(n, 10)
}

// 关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值。
func f() (ret int) {
	defer func() {
		ret++
	}()
	// defer ret++
	return 1
}
func TestMain(t *testing.T) {
	// 函数作为参数传递
	callback(1, AddTen)
	// 匿名函数，或者成为闭包
	func() {
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()
	fmt.Println(f())
}