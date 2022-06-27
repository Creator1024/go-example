package loop

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestIfMultiReturn(t *testing.T) {
	var str string = "123"

	an, err := strconv.Atoi(str) // 转换失败，则an为0
	if err != nil {
		fmt.Printf("%s is not an integer -- error\n", str)
		os.Exit(1)
	}
	fmt.Printf("Integer is %d\n", an)
}
