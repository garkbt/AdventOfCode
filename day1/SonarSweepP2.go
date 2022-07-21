package main

import (
	"bufio"
	"container/list"
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
	var prevSum = 0
	var counter = 0

	//open the file with a scanner
	scanner := bufio.NewScanner(file)

	//error checking new scanner
	if err := scanner.Err(); err != nil {
		fmt.Println(scanner.Text())
	}

	//using a container/list to get it all to work more nicely together
	sliding := list.New()
	//for each line in the file
	for scanner.Scan() {
		//integer taken from new line
		var current = 0
		var currentSum = 0
		//get the new line integer string and convert it to current integer
		if cur, err := strconv.Atoi(scanner.Text()); err == nil {
			current = cur
		}
		//add newly read integer to front of list
		sliding.PushBack(current)
		if sliding.Len() == 3 {
			for element := sliding.Front(); element != nil; element = element.Next() {
				currentSum = currentSum + element.Value.(int)
			}
		}

		//if there is a completed set
		if prevSum != 0 {
			//compare if greater than previous
			if currentSum > prevSum {
				//print increased line
				fmt.Println(currentSum, " (increased)")
				//increment counter of increases
				counter++
				//if prev is greater than current
			} else if currentSum < prevSum {
				fmt.Println(currentSum, " (decreased)")
				//if the same
			} else {
				fmt.Println(currentSum, " (no change)")
			}
		} else {
			fmt.Println(currentSum, "(N/A - no previous measurement)")

		}
		//reset the previous to be the old current before new line read
		prevSum = currentSum

		if sliding.Len() == 3 {
			sliding.Remove(sliding.Front())
		}

	}
	fmt.Println("Total number of increases", counter)
}
