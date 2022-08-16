package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	//initial variables for the counting
	var prev = 0
	var counter = 0

	//open the file with a scanner
	scanner := bufio.NewScanner(file)

	//error checking new scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(scanner.Text())
	}

	//for each line in the file
	for scanner.Scan() {
		//integer taken from new line
		var current = 0
		//get the new line integer string and convert it to current integer
		if cur, err := strconv.Atoi(scanner.Text()); err == nil {
			current = cur
		}

		//if there is a previous number
		if prev != 0 {
			//compare if greater than previous
			if current > prev {
				//print increased line
				fmt.Println(current, " (increased)")
				//increment counter of increases
				counter++
				//if prev is greater than current
			} else if current < prev {
				fmt.Println(current, " (decreased)")
				//if the same
			} else {
				fmt.Println(current, " (no change)")
			}
		} else {
			fmt.Println(current, "(N/A - no previous measurement")

		}
		//reset the previous to be the old current before new line read
		prev = current

	}
	fmt.Println("Total number of increases", counter)
}
