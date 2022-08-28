package mystruct

import "fmt"

type innerStruct struct {
	int1 int
	int2 int
}

// ExpStruct is a test struct
type ExpStruct struct {
	Mi1 int
	Mf1 float32
	int // 匿名字段
	innerStruct
}

// Employee is xx
type Employee struct {
	name   string
	salary float32
}

// NewEmployee is a construct for employee
func NewEmployee(name string, salary float32) *Employee {
	return &Employee{name, salary}
}

// AddEmployeeSalary is x
func (e *Employee) AddEmployeeSalary(percent float32) {
	// 指针作为方法接收者（会自动解指针）
	// 可以改变对象的值
	e.salary = e.salary + e.salary*percent
}

// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.6.md#1065-%E5%86%85%E5%B5%8C%E7%B1%BB%E5%9E%8B%E7%9A%84%E6%96%B9%E6%B3%95%E5%92%8C%E7%BB%A7%E6%89%BF
// 内嵌结构体和继承

// Engine is a interface x
type Engine interface {
	Start()
	Stop()
}

// Car has a Engine
type Car struct {
	wheelCount int
	Engine
}

// SetWheels used for set wheel
func (c *Car) SetWheels(num int) {
	c.wheelCount = num
}
func (c *Car) numberOfWheels() {
	fmt.Printf("this car has %d wheels", c.wheelCount)
}

// Mercedes is type of car
type Mercedes struct {
	Car
}

// GoToWorkIn is a Car func
func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}
