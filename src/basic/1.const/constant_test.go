package constant

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstant(t *testing.T)  {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable)
}

const (
	Mon = iota + 1
	Tue
	Wed
)

func TestConstant2(t *testing.T)  {
	t.Log(Mon, Tue, Wed)
}