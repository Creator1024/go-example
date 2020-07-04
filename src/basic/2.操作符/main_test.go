// ==比较数组
// 相同维数切含有相同个数的数组才能比较
// 每个元素相同的才相等
package operator

import "testing"

func TestArrayEqual(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	t.Log(a == b)
}

// &^：与0操作，结果为本身，与1操作，结果变为0
func TestBitClear(t *testing.T) {
	a := 7        // 0111
	x := 1        // 0001
	y := 2        // 0010
	t.Log(a &^ x) // 第一位清0 结果为6
	t.Log(a &^ y) // 第二位清0 结果为5
}
