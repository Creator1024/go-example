package maptest

import "testing"

// map定义
func TestMap(t *testing.T)  {
		m0 := map[string]int{"one":1,"two":2}
		m1 := map[string]int{}
		m1["one"] = 1
		m2 := make(map[string]int, 10/*initial Capacity*/)
		t.Log(m0, m1, m2)
}

//遍历
func TestTraveMap(t *testing.T){
	m1 := map[int]int {1:1, 2:2 ,3:3}
	for k,v := range m1{
		t.Log(k, v)
	}
}

// 访问不存在的key
func TestAccessNotExistingKey(t *testing.T)  {
	m1:=map[int]int{}
	t.Log(m1[0] ) // key不存在，返回0
	m1[1] = 0 
	t.Log(m1[1]) // key的值为0，返回0
	if v, ok := m1[2]; ok{
		t.Logf("Key 2's value is %d", v)
	} else {
		t.Log("Key 2 is not existing")
	}
}

// 工厂模式
func TestMapWithFunVaule(t *testing.T){
	m := map[int]func(op int) int{}
	m[0] = func(op int) int {return op}
	m[1] = func(op int) int {return op*op}
	m[2] = func(op int) int {return op*op*op}
	t.Log(m[0](2), m[1](2), m[2](2))
}

// map实现set
func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n]{
		t.Logf("%d is existing", n)
	}else{
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 3)
	t.Log(len(mySet))
}