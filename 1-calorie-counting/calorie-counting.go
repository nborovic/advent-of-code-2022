package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := LoadInput("input.txt")

	var n uint = 3

	sumOfAllMostCarriedCalories := 0
	topThreeMostCarriedCalories := GetTopNMostCarriedCalories(input, n)

	for _, nthMostCarriedCalories := range topThreeMostCarriedCalories {
		sumOfAllMostCarriedCalories += nthMostCarriedCalories
	}

	fmt.Printf("Sum of the top %v most carried calories is %v\n", n, sumOfAllMostCarriedCalories)
}

func LoadInput(fileName string) []string {
	input := []string{}

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		log.Panic(err)
	}

	return input
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
