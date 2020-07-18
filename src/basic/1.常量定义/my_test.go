package constant

import "testing"

func TestMyTest(t *testing.T) {
	var str string = "你好"
	t.Logf("%c", str[0])
}
