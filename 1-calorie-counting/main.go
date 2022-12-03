package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	PrintMaxCarriedCalories()
	PrintTopNMostCarriedCalories(3)
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
