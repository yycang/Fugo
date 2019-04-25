package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string = "Hi, how are you? Hi"

	fmt.Printf("%t\n", strings.HasPrefix(str, "Hi"))	// 判断字符串开头

	fmt.Printf("%t\n", strings.HasSuffix(str, "Hi"))	// 判断字符串结尾

	fmt.Printf("%t\n", strings.Contains(str, "Hi"))	// 判断字符串是否包含

	fmt.Printf("%d\n", strings.Index(str, "Hi"))		// 返回字符串第一个出现的索引

	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))	// 返回字符串最后出现的索引

	fmt.Printf("%s\n", strings.Replace(str, "Hi", "hi", -1))	// 替换父字符串中的字符串，n为替换的个数，为-1时替换全部

	fmt.Printf("%d\n", strings.Count(str, "Hi"))	// 返回字符串出现的次数

	fmt.Printf("%s\n", strings.Repeat(str, 2))	// 重复字符串2次

	fmt.Printf("%s\n", strings.ToLower(str))				// 将字符串全部小写

	fmt.Printf("%s\n", strings.ToUpper(str))				// 将字符串全部大写

	fmt.Printf("%s\n", strings.TrimSpace(str))			// 将字符串前后的空格剔除

	fmt.Printf("%s\n", strings.Trim(str, "Hi"))			// 将字符串前后的指定字符去除

	fmt.Printf("%q\n", strings.Fields(str))			// 将字符串以空格分割成列表

	fmt.Printf("%q\n", strings.Split(str, ","))			// 将字符串以指定字符分割成列表

	sl := strings.Fields(str)

	fmt.Printf("%s\n", strings.Join(sl, "-"))			// 将字符串以指定字符拼接
}