package main


type Foo map[string]string
type Bar struct {
	one string
	two int
}

func main() {
	y := new(Bar)
	(*y).one = "hello"
	(*y).two = 5

	z := make(Bar)	// 无法make一个结构体

	x := make(Foo)
	x["x"] = "hello"
	x["y"] = "world"

	u := new(Foo)
	(*u)["x"] = "hello"		// 会出现错误， new(Foo)返回一个指向nil的指针，尚未被分配内存，无法填充
}