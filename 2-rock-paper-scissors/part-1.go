package main

import (
	"fmt"
)

func PrintTotalScoreOfAllGames() {
	input := LoadInput("input.txt")

	inputOptionMap := make(map[byte]string)
	inputOptionMap['A'], inputOptionMap['B'], inputOptionMap['C'] = Rock, Paper, Scissors
	inputOptionMap['X'], inputOptionMap['Y'], inputOptionMap['Z'] = Rock, Paper, Scissors

	scoreSum := 0

	for _, line := range input {
		firstPlayerOption := inputOptionMap[line[0]]
		secondPlayerOption := inputOptionMap[line[2]]
		scoreSum += int(GetRoundScore(firstPlayerOption, secondPlayerOption))
	}

	fmt.Printf("Total fixed score of all games is %v\n", scoreSum)
}

func GetRoundScore(firstPlayerOption string, secondPlayerOption string) uint8 {
	optionPointsMap := make(map[string]uint8)
	optionPointsMap[Rock], optionPointsMap[Paper], optionPointsMap[Scissors] = 1, 2, 3

	resultPointsMap := make(map[string]uint8)
	resultPointsMap[fmt.Sprintf("%v %v", Rock, Rock)] = 3
	resultPointsMap[fmt.Sprintf("%v %v", Paper, Paper)] = 3
	resultPointsMap[fmt.Sprintf("%v %v", Scissors, Scissors)] = 3
	resultPointsMap[fmt.Sprintf("%v %v", Rock, Paper)] = 6
	resultPointsMap[fmt.Sprintf("%v %v", Paper, Scissors)] = 6
	resultPointsMap[fmt.Sprintf("%v %v", Scissors, Rock)] = 6
	resultPointsMap[fmt.Sprintf("%v %v", Rock, Scissors)] = 0
	resultPointsMap[fmt.Sprintf("%v %v", Paper, Rock)] = 0
	resultPointsMap[fmt.Sprintf("%v %v", Scissors, Paper)] = 0

	return resultPointsMap[fmt.Sprintf("%v %v", firstPlayerOption, secondPlayerOption)] + optionPointsMap[secondPlayerOption]
}
