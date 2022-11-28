package main

import (
	"fmt"
	"reflect"
)

func main() {
	atUserMap := map[string]*string{}

	fmt.Println(reflect.TypeOf(atUserMap))
	fmt.Println(reflect.ValueOf(atUserMap))
	fmt.Println(atUserMap == nil)
	fmt.Println(len(atUserMap) == 0)
}
