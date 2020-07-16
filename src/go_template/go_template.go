package main

import (
   "fmt"
)

// 在完成包的 import 之后，开始对常量、变量和类型的定义或声明

const c = "C"

var v int = 5

type T struct{}

// 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）

func init() { // initialization of package
}

// 如果当前包是 main 包，则定义 main 函数
func main() {
   var a int
   Func1()
   // ...
   fmt.Println(a)
}


// 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数
// 如果有很多函数，则可以按照字母顺序来进行排序

func (t T) Method1() {
   //...
}

func Func1() { // exported function Func1
   //...
}