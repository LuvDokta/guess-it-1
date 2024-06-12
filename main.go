package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"guess/calculations"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Can not read file", err)
	}
	defer file.Close()
	var numbers []float64
	scanner := bufio.NewScanner(file)
	var lines string
	var words []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		words = strings.Split(line, " ")
		if len(words) > 1 {
			fmt.Println("wrong format")
			return
		}
		for _, v := range words {
			if !(v == "") {
				num, err := strconv.ParseFloat(line, 64)
				if err != nil {
					fmt.Println("The data contains non-integer", err)
				}
				numbers = append(numbers, num)
			}
		}
		lines = lines + line
	}
	if lines == "" {
		fmt.Println("the file is empty")
		return
	}

	if len(numbers) < 2 {
		fmt.Println("Have at least 2 numbers")
		return
	}

	for i := 1; i < len(numbers); i++ {
		currentNum := numbers[i-1]
		nextNum := numbers[i]
		rangeStart, rangeEnd := calculations.GuessRange(numbers, currentNum, nextNum)
		fmt.Printf("%d--> the standard input\n", int(currentNum))
		fmt.Printf("%d %d--> the range for the next input, in this case for the number\n", rangeStart, rangeEnd)
	}
}