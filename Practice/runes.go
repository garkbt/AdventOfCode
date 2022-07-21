package main

import (
	"fmt"
)

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	//this prints ��=� ⌘
	fmt.Println(sample)

	//prints out in spaced output
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	//prints out in normal format
	fmt.Printf("\n%x\n", sample)

	fmt.Printf("% x", sample)

	fmt.Printf("\n%q\n", sample)

}
