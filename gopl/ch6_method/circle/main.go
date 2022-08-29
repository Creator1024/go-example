package main

import "fmt"

type Circle struct {
	radius float64
}

// 将方法绑定到Circle类型，Circle对象就可以使用这个方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var c1 Circle
	var c2 Circle
	c1.radius = 10.00
	c2.radius = 12.00

	fmt.Println("c1 area is: ", c1.getArea())
	fmt.Println("c2 area is: ", c2.getArea())
	fmt.Printf("%T \n", Circle.getArea)
}
