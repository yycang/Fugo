// 检测一个值是否实现了接口Stringer

if v，ok := v.(Stringer); ok {
	fmt.Printf("implements String(): %s\n", v.String())
}

// 使用接口实现一个类型分类函数

func classifier(items ...interface{})  {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("bool")
		case float64:
			fmt.Printf("float64")
		case int, int64:
			fmt.Printf("int")
		case nil:
			fmt.Printf("nil")
		case string:
			fmt.Printf("string")
		default:
			fmt.Printf("unknown")
		}
	}
}