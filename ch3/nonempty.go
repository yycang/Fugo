//slick修改
package main

// 去除字符串列表中空字符串
func nonempty(strings []string) []string {
	i := 0
	for _, item := range strings {
		if item != "" {
			strings[i] = item
			i++
		}
	}
	return strings[:i]
}