// 创建
arr1 := new([len]type)
slice1 := make([]type, len)

// 初始化
arr1 := [...]type{i1, i2, i3, i4, i5}
arrKeyValue := [len]type{i1: val1, i2: val2}
var slice1 []type = arr1[start:end]

// 如何截断数组或者切片的最后一个元素

line = line[:len(line)-1]

// 遍历数组

for i:=0; i<len(arr); i++ {
	print(arr[i])
}

// 如何在一个二维数组或者切片中的查找指定值

found := false
Found: for row := range arr2Dim {
	for column := range arr2Dim[row] {
		if arr2Dim[row][column] == v {
			found = true
			break Found
		}
	}
}