package main

import "fmt"

const (
	Win  string = "Win"
	Draw string = "Draw"
	Loss string = "Loss"
)

func PrintTotalFixedScoreOfAllGames() {
	input := LoadInput("input.txt")

	inputOptionMap := make(map[byte]string)
	inputOptionMap['A'], inputOptionMap['B'], inputOptionMap['C'] = Rock, Paper, Scissors

	fixedOutcomeMap := make(map[byte]string)
	fixedOutcomeMap['X'], fixedOutcomeMap['Y'], fixedOutcomeMap['Z'] = Loss, Draw, Win

	scoreSum := 0

	for _, line := range input {
		firstPlayerOption := inputOptionMap[line[0]]
		fixedOutcome := fixedOutcomeMap[line[2]]
		scoreSum += int(GetFixedRoundScore(firstPlayerOption, fixedOutcome))
	}

	fmt.Printf("Total score of all games is %v\n", scoreSum)
}

func GetFixedRoundScore(firstPlayerOption string, fixedOutcome string) uint8 {
	optionPointsMap := make(map[string]uint8)
	optionPointsMap[Rock], optionPointsMap[Paper], optionPointsMap[Scissors] = 1, 2, 3

	fixedOutcomePoints := make(map[string]uint8)
	fixedOutcomePoints[Win], fixedOutcomePoints[Draw], fixedOutcomePoints[Loss] = 6, 3, 0

	secondPlayerOptionMap := make(map[string]string)
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Rock, Win)] = Paper
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Paper, Win)] = Scissors
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Scissors, Win)] = Rock
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Rock, Draw)] = Rock
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Paper, Draw)] = Paper
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Scissors, Draw)] = Scissors
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Rock, Loss)] = Scissors
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Paper, Loss)] = Rock
	secondPlayerOptionMap[fmt.Sprintf("%v %v", Scissors, Loss)] = Paper

	return fixedOutcomePoints[fixedOutcome] + optionPointsMap[secondPlayerOptionMap[fmt.Sprintf("%v %v", firstPlayerOption, fixedOutcome)]]
}
