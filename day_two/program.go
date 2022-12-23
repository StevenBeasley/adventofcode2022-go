package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Play struct {
	Opponent int
	Player   int
}

func ParseFile(path string) ([]Play, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Parse(file)
}

func Parse(r io.Reader) (parsed []Play, err error) {
	s := bufio.NewScanner(r)

	var line int
	for s.Scan() {
		line++

		row := strings.Fields(s.Text())
		if len(row) < 2 {
			return nil, fmt.Errorf("line %v has unexpected number of fields: %v", line, len(row))
		}

		parsed = append(parsed, Play{
			Opponent: int(row[0][0]) - 'A',
			Player:   int(row[1][0]) - 'X',
		})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return parsed, nil
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
	parsedInput, err := ParseFile(os.Args[1])
	if err != nil {
		log.Fatalf("parse input: %v", err)
	}

	partOneOutput := PartOne(parsedInput)
	partTwoOutput := PartTwo(parsedInput)
	fmt.Println(partOneOutput)
	fmt.Println(partTwoOutput)
}

func PartTwo(input []Play) (score int) {
	results := [...]int{Loss, Draw, Win}
	scores := [...]int{Rock, Paper, Scissors}

	for _, play := range input {
		score += results[play.Player]
		offset := (play.Player + 2) % 3
		score += scores[(play.Opponent+offset)%len(scores)]
	}

	return score
}

func PartOne(input []Play) (score int) {
	results := [...]int{Win, Loss, Draw, Win, Loss}
	scores := [...]int{Rock, Paper, Scissors}

	for _, play := range input {
		score += scores[play.Player]
		score += results[play.Player-play.Opponent+2]
	}

	return score
}
