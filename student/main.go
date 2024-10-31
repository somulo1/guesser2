package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess-it-2/helperfunctions"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	var sliceNum []float64

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content := scanner.Text()
		if content == "" {
			return
		}

		num, err := strconv.ParseFloat(content, 64)
		if err != nil {
			fmt.Println("Error converting to float64", err)
			return
		}
		var xi []float64
		sliceNum = append(sliceNum, num)
		for i := 0; i < len(sliceNum); i++ {
			xi = append(xi, float64(i))
		}

		if len(sliceNum) > 1 {

			lowerLimit, upperLimit := helperfunctions.Limit(xi, sliceNum)

			fmt.Println((lowerLimit), int(upperLimit))
		}
	}
}
