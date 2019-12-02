package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	subtotals := make([]int, 0, 10)
	total := 0
	file, err := os.Open("assets/day01input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result := int(math.Floor(float64(number)/float64(3))) - 2
		subtotals = append(subtotals, result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, subtotal := range subtotals {
		total += subtotal
	}

	fmt.Println(total)
}
