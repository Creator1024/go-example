package main

import (
	"bytes"
	"fmt"
	"io"
)

//var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	//flag.Parsed()
	//fmt.Printf("Sleeping for %v...", *period)
	//time.Sleep(*period)
	//fmt.Println()

	//var w1 io.Writer
	//var w2 io.Writer
	//fmt.Println(w1 == w2)
	//w1 = os.Stdout
	//w2 = new(bytes.Buffer)
	//fmt.Println(w1 == w2)

	//var x interface{}
	//x = []int{1, 2, 3}
	//fmt.Println(x == x)

	//var w io.Writer
	//fmt.Printf("%T\n", w) // <nil>
	//fmt.Printf("%v\n", w) // <nil>
	//
	//w = os.Stdout
	//fmt.Printf("%T\n", w) // *os.File
	//fmt.Printf("%v\n", w) // &{0xc000022120}

	var i interface{}
	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)

	var buf io.Writer
	buf = new(bytes.Buffer)
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done\n"))
	}
}
