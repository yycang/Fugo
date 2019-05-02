package main

import (
	"fmt"
	"time"
)


type Address struct {
	Street string
	HouseNumber uint32
	ZipCode string
	City string
	Country string
}

type Vcard struct {
	FirstName string
	LastName string
	NickName string
	BirthDate time.Time
	Photo	string
	Address map[string]*Address
}

func main() {
	addr1 := &Address{"shifoying", 12, "", "beijing", "china"}
	addr2 := &Address{"chaoyang", 15, "", "beijing", "china"}

	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birthdt := time.Date(1998, 1, 15, 14, 4, 5, 0, time.Local)
	photo := "photo1.jpg"
	vcard := &Vcard{"cang", "yuanyuan", "", birthdt, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}