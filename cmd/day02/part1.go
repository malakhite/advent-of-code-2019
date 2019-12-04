package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readValues() []int {
	file, err := os.Open("assets/day02input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	input := scanner.Text()

	stringValues := strings.Split(input, ",")

	intValues := make([]int, 0, 10)

	for _, value := range stringValues {
		number, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		intValues = append(intValues, number)
	}

	return intValues
}

func calculateMemory(input []int) int {
	for index, value := range input {
		if index%4 == 0 {
			// if input[index+3] >= cap(input) && value != 99 {
			// 	newInput := make([]int, (input[index+3])*2)
			// 	copy(newInput, input)
			// 	input = newInput
			// }

			if value == 1 {
				input[input[index+3]] = input[input[index+1]] + input[input[index+2]]
			}

			if value == 2 {
				input[input[index+3]] = input[input[index+1]] * input[input[index+2]]
			}

			if value == 99 {
				break
			}
		}
	}

	return input[0]
}

func part1() {
	input := readValues()

	input[1] = 12
	input[2] = 2

	result := calculateMemory(input)

	fmt.Println(result)
}
