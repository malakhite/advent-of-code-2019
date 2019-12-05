package main

import (
	"log"
	"strconv"
)

func recordPath(route []string) [][]int {
	location := [2]int{0, 0}
	path := make([][]int, 0, 10)
	for _, point := range route {
		direction := point[0]
		distance, err := strconv.Atoi(point[1:])
		if err != nil {
			log.Fatal(nil)
		}
		if direction == "U" {
		}
	}
}

func part1() {

}
