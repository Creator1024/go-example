package factory_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type file struct{		// 使用小写的类，外部包不可见，强制使用工厂方法创建对象（禁止使用new）
	fd	int		// 8bit
	name string	// 16bit
}

func NewFile(fd int, name string) *file{
	if fd < 0 {
        return nil
	}
    return &file{fd, name}
}
func TestMain(t *testing.T){
	f := NewFile(123, "test.txt")
	t.Log(f.fd)
	t.Log(unsafe.Sizeof(File{}))	// 查看一个结构体占用多少内存   24bit
}