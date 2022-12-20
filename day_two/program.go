package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type Play struct {
	opponent string
	player   string
}

func parseInput(input string) (parsed []Play) {
	for _, row := range strings.Split(input, "\n") {
		split := strings.Split(row, " ")
		parsed = append(parsed, Play{
			opponent: split[0],
			player:   split[1],
		})
	}
	return
}

const (
	Win  = 6
	Loss = 0
	Draw = 3

	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func main() {
	parsedInput := parseInput(input)

	partOneOutput := partOne(parsedInput)
	partTwoOutput := partTwo(parsedInput)
	fmt.Println(partOneOutput)
	fmt.Println(partTwoOutput)
}

func partTwo(input []Play) (totalScore int) {
	choices := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Loss",
		"Y": "Draw",
		"Z": "Win",
	}

	scores := map[string]int{
		"Rock":     Rock,
		"Paper":    Paper,
		"Scissors": Scissors,
	}

	for _, play := range input {
		score := 0
		// convert to choices
		result := choices[play.player]
		opponent := choices[play.opponent]

		switch result {
		case "Win":
			score += Win
			switch opponent {
			case "Rock":
				score += scores["Paper"]
			case "Paper":
				score += scores["Scissors"]
			case "Scissors":
				score += scores["Rock"]
			}
		case "Loss":
			score += Loss
			switch opponent {
			case "Rock":
				score += scores["Scissors"]
			case "Paper":
				score += scores["Rock"]
			case "Scissors":
				score += scores["Paper"]
			}
		case "Draw":
			score += Draw
			score += scores[opponent]
		}

		totalScore += score

	}

	return totalScore

}

func partOne(input []Play) (totalScore int) {
	choices := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}
	scores := map[string]int{
		"Rock":     Rock,
		"Paper":    Paper,
		"Scissors": Scissors,
	}

	for _, play := range input {

		score := 0
		// convert to choices
		player := choices[play.player]
		opponent := choices[play.opponent]

		score += scores[player]

		if player == opponent {
			score += Draw
		} else {
			switch player {
			case "Rock":
				if opponent == "Scissors" {
					score += Win
				}
			case "Paper":
				if opponent == "Rock" {
					score += Win
				}
			case "Scissors":
				if opponent == "Paper" {
					score += Win
				}

			}
		}
		totalScore += score
	}

	return totalScore
}
