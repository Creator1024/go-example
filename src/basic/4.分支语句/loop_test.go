package loop

import (
	"fmt"
	"testing"
	"time"
)

func TestMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}
}

func TestForAndSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("0 or 2")
		case 1, 3:
			t.Log("1 or 3")
		default:
			t.Log("not 0-3")
		}
	}
	for i := 0; i < 5; i++ {
		switch { //类似 if else-if else
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")
		default:
			t.Log("unknow")
		}
	}

	a := 10
	for true {
		fmt.Println(a)
		a--
		time.Sleep(time.Second * 3)
	}
}
