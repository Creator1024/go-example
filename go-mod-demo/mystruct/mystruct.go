package mystruct

type innerStruct struct {
	int1 int
	int2 int
}

// ExpStruct is a test struct
type ExpStruct struct {
	Mi1 int
	Mf1 float32
	int // 匿名字段
	innerStruct
}
