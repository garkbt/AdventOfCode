package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePtr, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	bufferedReader := bufio.NewReader(filePtr)

	for 
	data, err := bufferedReader.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
