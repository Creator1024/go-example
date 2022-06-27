package constant

import "testing"

var a string

func TestLocalScope(t *testing.T) {
	a = "G"
	println(a)
	n()
}

func n() {
	a := "O"
	println(a)
	m()
}

func m() {
	println(a)
}
