package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func main() {
	names := []string{
		"3.cx",
		"4.dx",
		"2.bx",
		"1.ax",
		"1.bx",
	}
	//sort.Sort(MyStringList(names))
	sort.Strings(names)
	ages := []int{18, 17, 19}
	sort.Ints(ages)
	scores := []float64{80.5, 75.0, 99.5}
	sort.Float64s(scores)
	for _, v := range scores {
		fmt.Printf("%s\n", v)
	}
}
