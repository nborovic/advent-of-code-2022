package main

import (
	"fmt"
	"log"
	"strconv"
)

func PrintMaxCarriedCalories() {
	input := LoadInput("input.txt")

	maxCarriedCalories := GetMaxCarriedCalories(input)

	fmt.Printf("Max carried calories is %v\n", maxCarriedCalories)
}

func GetMaxCarriedCalories(input []string) int {
	maxCarriedCalories := 0
	currentCalorieCounter := 0

	for _, line := range input {
		if line == "" {
			if currentCalorieCounter > maxCarriedCalories {
				maxCarriedCalories = currentCalorieCounter
			}
			currentCalorieCounter = 0
			continue
		}

		caloriesCount, err := strconv.Atoi(line)

		if err != nil {
			log.Panic(err)
		}

		currentCalorieCounter += caloriesCount
	}

	return maxCarriedCalories
}
