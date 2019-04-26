package main

import "fmt"

// GOTO或者标签是不被鼓励的！
func main()  {
	LABEL1:
		for i:=0; i<=5; i++ {
			for j:=0; j<=5; j++ {
				if j == 4 {
					continue LABEL1
				}
				fmt.Println(i, j)
			}
		}

	i := 0
	HEAR:
		print(i)
		i++
		if i == 5 { return }
		goto HEAR
}