package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Need a 2d array to store the data
	//var binary = [][]int{}

	//var counter
	//var count = 0

	//file to open
	file, err := os.Open("input")
	if err != nil {
		//if error print it
		fmt.Println(err)
	}

	//open the file with a scanner
	scanner := bufio.NewScanner(file)

	//error checking new scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(scanner.Text())
	}

	//for each line in the file
	for scanner.Scan() {
		//get the new line integer string and convert it to current integer

		var line = scanner.Text()
		fmt.Println(line)
	}

	//gamma rate is the most common bit in the corresponding position

	//epsilon rate is calculated by the opposit of the first (xor)
}
