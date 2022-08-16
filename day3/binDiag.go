package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	//Need a 2d array to store the data
	//var binary = [][]int{}

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

	binaryMatrix := make([][]int, 0)
	//for each line in the file
	for scanner.Scan() {
		//get the new line integer string and convert it to current integer
		temp := make([]int, 0)
		var line = scanner.Text()
		//this converts the string to runes
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			integerVal := int(runes[i] - '0')
			temp = append(temp, integerVal)
		}

		binaryMatrix = append(binaryMatrix, [][]int{temp}...)
	}
	fmt.Println(len(binaryMatrix))

	//gamma rate is the most common bit in the corresponding position
	var gammaRate [12]int
	for xx := 0; xx < 12; xx++ {
		sum := 0
		for yy := 0; yy < len(binaryMatrix); yy++ {
			sum += binaryMatrix[yy][xx]
		}
		if sum > 500 {
			sum = 1
		} else {
			sum = 0
		}
		gammaRate[xx] = sum

	}
	fmt.Println("Gamma Rate: ", gammaRate)
	//sum them and know the length of the splice, specifically the number of splices it containts
	//epsilon rate is calculated by the opposit of the first (xor)
	var epsilonRate [12]int
	for gg := 0; gg < 12; gg++ {
		if gammaRate[gg] == 1 {
			epsilonRate[gg] = 0
		} else if gammaRate[gg] == 0 {
			epsilonRate[gg] = 1
		}
	}
	fmt.Println("Epsilon Rate: ", epsilonRate)
	var ep float64 = 0
	for oo := 11; oo > -1; oo-- {
		if epsilonRate[oo] == 1 {
			ep += math.Pow(2, float64(11-oo))
		}
	}
	fmt.Println("Epsilon: ", ep)

	var gr float64 = 0
	for oo := 11; oo > -1; oo-- {
		if gammaRate[oo] == 1 {
			gr += math.Pow(2, float64(11-oo))
		}
	}
	fmt.Println("Gamma: ", gr)

	finalResult := gr * ep
	fmt.Printf("Final Result: %f", finalResult)
}
