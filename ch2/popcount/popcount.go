package popcount

// 包的初始化要解决变量的依赖顺序，然后根据包级变量声明出现的顺序
var a = b + c	// 第三个初始化
var b = PopCount(1)	// 第二个初始化
var c = 1		// 第一个初始化


var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 返回x的种群统计(置位的个数） 	种群统计不明白。。
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

