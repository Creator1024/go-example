package array_test

import "testing"

//1. 数组不可伸缩，切片可以
//2. 切片不可比较，数组可以

func TestArray(t *testing.T) {
	//数组声明，默认赋值0
	var a [3]int
	a[0] = 1
	t.Log(a[0], a[1], a[2])
	//声明时赋值
	b := [3]int{1, 2, 3}
	c := [2][2]int{{1, 2}, {3, 4}}
	t.Log(b, c)
	// 切片和py类似
	t.Log(b[1:])
	//不支持负索引
	//t.log(b[-1:])
}
func TestArrayTravel(t *testing.T) {
	// 使用...定义的长度，会跟赋值一样
	d := [...]int{10, 20, 3, 4, 5}
	// for i := 0; i < len(d); i++ {
	// 	t.Log(d[i])
	// }
	for _, v := range d {
		t.Log(v)
	}
}

//切片的内部结构
// ptr, len, cap
func TestSlice(t *testing.T) {
	// 切片（slice）的声明，无长度，而数组需要定义长度
	var s0 []int
	t.Log(len(s0), cap(s0)) // 0 0
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) // 1 1

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1)) // 4 4

	//make(type, len, cap)
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))    // 3 5
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0不能访问s2[3]+
	s2 = append(s2, 1)         // 填充之后，就可以访问s2[3]
	t.Log(len(s2), cap(s2))    // 4 5  length变长
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // cap 以2的指数增加，1、2、4、8、16...
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec"}
	summary := year[5:8]
	t.Log(summary)                    //[Jun Jul Aug]
	t.Log(len(summary), cap(summary)) // 3 7
	Q2 := year[3:6]
	t.Log(Q2)               // [Apr May Jun]
	t.Log(len(Q2), cap(Q2)) // 3 9
	Q2[2] = "Modify"
	t.Log(summary) // [Modify Jul Aug]
	t.Log(year)
}
