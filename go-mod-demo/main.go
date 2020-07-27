package main

import (
	"fmt"
	"test/mystruct"
)

func main() {
	myExpStruct := new(mystruct.ExpStruct)
	employeeWW := mystruct.NewEmployee("ww", 1000.0)
	fmt.Println(employeeWW)
	employeeWW.AddEmployeeSalary(0.2)
	fmt.Println(employeeWW)
	fmt.Println(myExpStruct)
	fmt.Println(IntVendor{1, 2, 3}.Sum())
	// ---
	car := &mystruct.Car{}
	car.SetWheels(4)
	fmt.Println(car)
	fmt.Println(car.Stop())
}

// IntVendor is xx
type IntVendor []int

// Sum is a function for IntVendor
func (v IntVendor) Sum() (s int) {
	for _, value := range v {
		s += value
	}
	return
}

// https://zhuanlan.zhihu.com/p/109828249
