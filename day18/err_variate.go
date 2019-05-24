
var remeber bool = false

if something {
	remember := true 	// 错误用法
}

if something {
	remember = true		// 正确用法
}

for _, file := range files {
	if f, err = os.Open(file); err != nil {
		return
	}
	defer f.Close()		// 错误用法

	f.Close()			// 正确用法，defer仅在函数返回时才执行
}


fun findBiggest(listofNumbers []int) int {}		// 切片实际上是指向数组的指针，直接传递给函数即可

// 同样的接口已经是指针，不用将指针指向接口类型

type nexter interface {
	next() byte
}

func next1(n nexter, num int) []byte {
	var b []byte
	for i := 0; i<num; i++ {
		b[i] = n.next()
	}
	return b
}

func next2(n *nexter, num int) []byte {
	var b []byte
	for i := 0; i<num; i++ {
		b[i] = n.next()		// 编译错误
	}
	return b
}