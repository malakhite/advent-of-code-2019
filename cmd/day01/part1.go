package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getModuleWeights() []int {
	file, err := os.Open("assets/day01input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	weights := make([]int, 0, 10)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		weights = append(weights, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return weights
}

func part1() {
	fuelSubtotals := make([]int, 0, 10)
	fuelTotal := 0

	weights := getModuleWeights()

	for _, subtotal := range subtotals {
		total += subtotal
	}

	fmt.Println(total)
}
