package main

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(len(s))

	s = "中"
	t.Log(len(s)) // 3  len取的是bytes数

	c := []rune(s) // rune -- 取出字符串中的unicode
	t.Log(len(c))  // 1

	// unicode是一种字符集（定义）
	// UTF8 是 unicode的存储实现（实现）

	t.Logf("中 -- unicode： %x", c[0]) // 4e2d
	t.Logf("中 -- UTF8 %x", s)        // e4b8ad

}

func TestStringToRune(t *testing.T) {
	s := "字符串"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
