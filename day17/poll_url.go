package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.baidu.com/",
	"http://bilibili.com/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("error", url, err)
		}
		fmt.Println(url, ":", resp.Status)
	}
}