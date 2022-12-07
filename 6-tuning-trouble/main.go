package main

import (
	"log"
	"os"
)

func main() {
	PrintLengthUntilTheEndOfFirstMarker(4)
	PrintLengthUntilTheEndOfFirstMarker(14)
}

func LoadInput(fileName string) string {
	file, err := os.ReadFile(fileName)

	if err != nil {
		log.Panic(err)
	}

	return string(file)
}
