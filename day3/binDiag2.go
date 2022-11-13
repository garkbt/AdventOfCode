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
	//sum them and know the length of the splice, specifically the number of splices it contains
	//epsilon rate is calculated by the opposite of the first (xor)
	var epsilonRate [12]int
	for gg := 0; gg < 12; gg++ {
		if gammaRate[gg] == 1 {
			epsilonRate[gg] = 0
		} else if gammaRate[gg] == 0 {
			epsilonRate[gg] = 1
		}
	}
	fmt.Println("Epsilon Rate: ", epsilonRate)


    //use the oxygen generator rating from the gamma rate
    //iterate through the list

    matches := make([]int, 0)           
    var oxygenRate [12]int 
    var topdog = -1
    var prevcount = 0
    //loop through the binary matrix
    for ii := 0; ii < 12; ii++ {
        for oo := 0; oo < len(binaryMatrix); oo++ {
            if binaryMatrix[oo][ii]==gammaRate[ii]{
                //add it to the list
                matches = append(matches, oo)
            }
        }    
    }
    
    for dd:=0; dd < 12; dd++ {
        oxygenRate[dd]=binaryMatrix[topdog][dd]
    }

    fmt.Printf("Oxygen Rate: %v\n", oxygenRate)

    var co2Rate [12]int 
    var smalldog = -1
    var precount = 0

    //loop through the binary matrix
    for oo := 0; oo < len(binaryMatrix); oo++ {
        var count = 0
        //loop through each line
        for ii := 0; ii < 12; ii++ {
            if binaryMatrix[oo][ii]==epsilonRate[ii] {
                count++
            } else{
                ii = 12
            }
        }
        if count > precount {
            precount = count
            smalldog = oo
        }
    }
    
    for dd:=0; dd < 12; dd++ {
        co2Rate[dd]=binaryMatrix[smalldog][dd]
    }

    fmt.Printf("CO2 Rate: %v\n", co2Rate)
            //use the epsilon rate from the epsilon rate



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
