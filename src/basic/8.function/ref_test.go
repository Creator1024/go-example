package reffn

import (
	"fmt"
	"testing"
)

func doSomething(a *int) {
	b := a
	fmt.Println(b)
	fmt.Println(*b)
}
func doSomething2(a int) {
	// 这里是值传递，复制了一份a，所以b的值与main里面&a的值不一样，也改变不到
	b := &a
	*b = 4
	fmt.Println(b)
	fmt.Println(*b)
}

func TestRef(t *testing.T) {
	var a int = 10
	fmt.Println(&a)
	doSomething(&a)
	doSomething2(a)
	fmt.Println(a)

}

/*
output:
0xc000092080
0xc000092080
10
0xc000092090
4
10
*/
