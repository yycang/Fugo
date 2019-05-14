package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

// hash 实现adler32, crc32, crc64, fnv
// crypto 实现md4, md5, sha1, aes, blowfish, rc4, rsa, xtea等加密算法
func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "1234")
	b := []byte{}

	fmt.Printf("result: %x\n", hasher.Sum(b))
	fmt.Printf("result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("we shall overcome")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("hash write error %v / %v", n, err)
	}

	checksum := hasher.Sum(b)
	fmt.Printf("result: %x\n", checksum)
}