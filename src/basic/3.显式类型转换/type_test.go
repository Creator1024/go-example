package typeexchange

import "testing"

type MyInt int64

func TestTypeExchange(t *testing.T) {
	// go不支持任何隐式转换，包括别名
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)

	t.Log(a, b, c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	// 指针不能运算
	// aPtr += 1
	t.Logf("%T, %T", a, aPtr)
}

func TestStringInit(t *testing.T) {
	var a string
	// string类型默认为空字符串""
	t.Log("*" + a + "*")
}