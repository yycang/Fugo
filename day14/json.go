package main

import (
	"encoding/json"
	"fmt"
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
	Addresses []*Address
	Remark string
}

func main() {
	pa := &Address{"home", "beijing", "china"}
	wa := &Address{"work", "tianjin", "china"}
	va := VCard{"c", "yy", []*Address{pa, wa}, "none"}

	js, _ := json.MarshalIndent(va, "", "\t")
	fmt.Printf("json format: %s", js)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()

	enc := json.NewEncoder(file)
	err := enc.Encode(va)
	if err != nil {
		log.Println("error in encoding json")
	}
}