package main


// 常见用法1
/*
value, err := pack1.Function1(param1)

if err != nil {
	fmt.Printf("An error occured in pack1.Function1 with %v", param1)
	return err
}

 */


// 常见用法2（终止运行程序
/*
if err != nil {
	fmt.Printf("Program stopped with error %v", err)
	os.Exit(1)
}
 */


// 无需else
/*
f, err := os.Open(name)
if err != nil {
	return err
}
doSomething(f)

 */


// 常见用法3（记得要为多返回值的函数准备变量
/*
if value, ok := readData(); ok {
	doSomething.....
}
 */


