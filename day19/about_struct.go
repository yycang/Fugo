// 创建

type struct1 struct {
	field1 type1
	field2 type2
}

ms := new(struct1)

// 初始化

ms := &struct1{10, 15.5, "Chris"}


// 结构体命名开头大写，则结构体在包外可见
ms := Newstruct1{10, 15.5, "Chris"}
func Newstruct1(n int, f float32, name string) *struct1 {
	return &struct1{n, f, name}
}