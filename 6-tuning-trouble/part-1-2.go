package main

import "fmt"

func PrintLengthUntilTheEndOfFirstMarker(numberOfDistinctCharacters int) {
	input := LoadInput("input.txt")

	for i := range input {
		markerIndexes := make(map[byte]int)

		for j := i; j <= i+numberOfDistinctCharacters-1 && j < len(input); j++ {
			markerIndexes[input[j]] = j
		}

		if len(markerIndexes) == numberOfDistinctCharacters {
			fmt.Printf("Length until the end of the first %v character marker is %v\n", numberOfDistinctCharacters, i+numberOfDistinctCharacters)
			break
		}
	}
}
