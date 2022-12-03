package main

import (
	"fmt"
	"log"
	"strconv"
)

func PrintTopNMostCarriedCalories(n uint) {
	input := LoadInput("input.txt")

	sumOfAllMostCarriedCalories := 0
	topThreeMostCarriedCalories := GetTopNMostCarriedCalories(input, n)

	for _, nthMostCarriedCalories := range topThreeMostCarriedCalories {
		sumOfAllMostCarriedCalories += nthMostCarriedCalories
	}

	fmt.Printf("Sum of the top %v most carried calories is %v\n", n, sumOfAllMostCarriedCalories)
}

func GetTopNMostCarriedCalories(input []string, n uint) []int {
	topNMostCarriedCalories := make([]int, n)

	minMostCarriedCaloriesIndex := 0
	currentCalorieCounter := 0

	for _, line := range input {
		if line == "" {
			if currentCalorieCounter > topNMostCarriedCalories[minMostCarriedCaloriesIndex] {
				topNMostCarriedCalories[minMostCarriedCaloriesIndex] = currentCalorieCounter

				for i, nthMostCarriedCalories := range topNMostCarriedCalories {
					if nthMostCarriedCalories < topNMostCarriedCalories[minMostCarriedCaloriesIndex] {
						minMostCarriedCaloriesIndex = i
					}
				}
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

	return topNMostCarriedCalories
}
