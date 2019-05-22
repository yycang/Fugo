package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)


type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status Status
}

func main() {
	// 发起对一个url的请求
	response, _ := http.Get("http://twitter.com/users/Googland.xml")

	user := User{xml.Name{"", "user"}, Status{""}}
	// 由于版本更新，传入的第一个参数应该是一个[]byte类型
	xml.Unmarshal(response.Body, &user)
	fmt.Printf("status: %s", user.Status.Text)

}