package main

import (
	"fmt"
)

type Camera struct { }

func (c *Camera) TakeAPicture() string {
	return "click"
}

type Phone struct { }

func (p *Phone) Call() string {
	return "ringring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)

	fmt.Println("camera: ", cp.TakeAPicture())
	fmt.Println("phone: ", cp.Call())
}