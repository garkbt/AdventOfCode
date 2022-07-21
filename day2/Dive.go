package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//assign the files and err variables the same value - opening the input file and handling the error != nil
	file, err := os.Open("input")
	if err != nil {
		//if error print it
		fmt.Println(err)
	}
	//close the file
	defer file.Close()

	//horizontal from 0
	//depth from 0
	var horizontal = 0
	var depth = 0

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
		var lineSplit = strings.Split(line, " ")
		var direction = lineSplit[0]

		amount, _ := strconv.Atoi(lineSplit[1])
		if err != nil {
		}

		switch direction {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}

		//switch on first word of line

	}
	fmt.Println("Horizontal: ", horizontal, "Depth: ", depth)

}

//read word and value from line
