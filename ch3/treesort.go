// 利用二叉树实现插入排序
package main

import "fmt"

type tree struct {
	value 		int
	left, right *tree
}

func Sort(values []int)  {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	// 传入空切片，好往里面塞值
	appendValues(values[:0], root)
}


// 将元素按照顺序追加到values中，返回slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	lst := []int{1, 4, 10, 2, 0}
	Sort(lst)
	fmt.Println(lst)
}
