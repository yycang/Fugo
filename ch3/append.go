// append方法实现
package main

func appendInt(x []int, y int) ([]int) {
	var z []int
	zlen := len(x) + 1
	// 是否需要扩容
	if zlen <= cap(x) {
		// 仍有扩展空间
		z = x[:zlen]
	} else {
		// 需要重新分配一个数组
		zcap := zlen
		// 分摊线性复杂，容量扩充一倍
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		// 将原切片copy过去
		copy(z, x)
	}
	z[len(x)] = y
	return z
}