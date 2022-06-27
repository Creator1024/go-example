package deferfn

import (
	"fmt"
	"testing"
)

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("do something...")
	panic("error")
	fmt.Println("can not exec...")
}
