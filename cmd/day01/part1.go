package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getModuleMasses() []int {
	file, err := os.Open("assets/day01input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	masses := make([]int, 0, 10)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		masses = append(masses, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return masses
}

func calculateFuel(masses []int) []int {
	fuelLoads := make([]int, 0, 10)

	for _, mass := range masses {
		fuelLoads = append(fuelLoads, (mass/3)-2)
	}

	return fuelLoads
}

func part1() {
	fuelTotal := 0

	masses := getModuleMasses()

	fuelLoads := calculateFuel(masses)

	for _, fuelLoad := range fuelLoads {
		fuelTotal += fuelLoad
	}

	fmt.Println(fuelTotal)
}
