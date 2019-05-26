// 修改字符串中的字符
str := "hello"
c := []byte(str)
c[0] = 'c'
s2 := string(c)

// 获取字符串中的字串

substr := str[n:m]

// 使用for遍历一个字符串

for i:=0; i<len(str); i++ {
	print(str[i])
}

// 如何获取一个字符串的字节数
len(str)

// 如何获取一个字符串的字符数
len([]rune(str))
utf8.RuneCountInString(str)

// 连接字符串
/* 	1.+=
	2.with a bytes.Buffer	比+=节省内存和cpu
	3.Strings.Join()
 */

