package main

import (
	"bufio"
	"log"
	"os"
)

const (
	Rock     string = "Rock"
	Paper    string = "Paper"
	Scissors string = "Scissors"
)

func main() {
	PrintTotalScoreOfAllGames()
	PrintTotalFixedScoreOfAllGames()
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
