// dup2打印输入中多次出现的行的个数与文本

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filename := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filename)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts, filename)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d,%s,%s\n", n, line, filename[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename map[string]string)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filename[input.Text()] = f.Name()
	}
}