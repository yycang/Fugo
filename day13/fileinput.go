package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, errorFile := os.Open("input.dat")
	if errorFile != nil {
		fmt.Println("file not exists")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("the input was %s", inputString)
		if readerError == io.EOF {	// io.EOF值为true
			return
		}
	}
}