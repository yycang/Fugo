package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type string
	City string
	Country string
}

type VCard struct {
	FirstName string
	LastName string
	Address []*Address
	Remark string
}

var content string

func main() {
	pa := &Address{"home", "beijing", "china"}
	wa := &Address{"work", "tianjin", "china"}
	va := VCard{"c", "yy", []*Address{pa, wa}, "none"}

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(va)
	if err != nil {
		log.Println("error in encoding gob")
	}
}