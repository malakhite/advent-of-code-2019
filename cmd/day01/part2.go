package main

import "fmt"

func calculateMoreFuel(thisFuel int, totalFuel int) int {
	if thisFuel/3 <= 2 {
		return totalFuel
	}
	newFuel := (thisFuel / 3) - 2

	return calculateMoreFuel(newFuel, (newFuel + totalFuel))
}

func part2() {
	masses := getModuleMasses()
	fuelLoads := calculateFuel(masses)
	newTotal := 0
	for _, fuelLoad := range fuelLoads {
		newTotal += calculateMoreFuel(fuelLoad, fuelLoad)
	}

	fmt.Println(newTotal)
}
