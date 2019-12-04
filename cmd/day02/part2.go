package main

import (
	"errors"
	"fmt"
	"log"
)

func calculateAnswer() (int, error) {
	answer := 19690720
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			input := readValues()
			input[1] = noun
			input[2] = verb

			if calculateMemory(input) == answer {
				fmt.Printf("Noun: %d, Verb: %d\n", noun, verb)
				return (100*noun + verb), nil
			}
		}
	}

	return 0, errors.New("No solution found")
}

func part2() {
	answer, err := calculateAnswer()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
