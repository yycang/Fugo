// 使用内建函数recover 终止panic过程

func protect(f func()) {
	defer func() {
		log.Println("done")
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}
}