package main

import (
	_ "embed"
	"errors"
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

func main() {
	parsedInput := parseInput(input)

	resultOne := partOne(parsedInput)
	resultTwo := partTwo(strings.Split(input, "\n"))
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func partTwo(parsedInput []string) int {
	total := 0
	currentGroup := make([]string, 0)
	for i, rucksack := range parsedInput {
		currentGroup = append(currentGroup, rucksack)
		if (i+1)%3 == 0 && i != 0 {
			groupResult, err := getBadgesFromGroup(currentGroup)
			if err != nil {
				panic(err)
			}
			total += groupResult
			currentGroup = make([]string, 0)
		}
	}
	return total
}

var ErrTooManyMatches = errors.New("too many matches in the group")

type MatchingBadges map[int]bool

func getBadgesFromGroup(currentGroup []string) (int, error) {
	// find matching item in all groups
	splitElves := make([][]rune, 0)
	for _, elf := range currentGroup {
		splitElves = append(splitElves, []rune(elf))
	}

	matches := make([]rune, 0)
	for i, elf := range splitElves {
		if i == 0 {
			matches = splitElves[0]
		} else {
			matches = findMatchingValues(matches, elf)
		}
	}

	deDuped := deDupe(matches)
	if len(deDuped) != 1 {
		return 0, ErrTooManyMatches
	}

	return getAsciiValue(deDuped[0]), nil
}

type DeDupable interface{ string | int | rune }

func deDupe[T DeDupable](items []T) []T {
	// add items to allKeys as they are found
	// when initially checking (and not set) the map will return false
	allKeys := make(map[T]bool)
	// add items to list when first found
	list := []T{}
	for _, item := range items {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func partOne(parsedInput []Rucksack) int {
	total := 0
	// find matching in each Rucksack

	for _, rucksack := range parsedInput {
		total += findMatchingAsciiValue(rucksack)
	}

	return total
}

func getAsciiValue(char rune) int {
	if char >= 'a' && char <= 'z' {
		// between a -> z then 1 => 26
		return int(char-'a') + 1
	}
	// A -> Z 27 -> 52
	return int(char-'A') + 27
}

func findMatchingAsciiValue(rucksack Rucksack) int {
	a := []rune(rucksack.a)
	b := []rune(rucksack.b)

	var match rune

	for _, aChar := range a {
		if containsString(b, aChar) {
			match = aChar
			break
		}
	}
	return getAsciiValue(match)
}

func findMatchingValues(a, b []rune) []rune {

	var matches []rune

	for _, aChar := range a {
		if containsString(b, aChar) {
			matches = append(matches, aChar)
		}
	}

	return matches
}

func containsString(b []rune, aChar rune) bool {
	for _, s := range b {
		if s == aChar {
			return true
		}
	}
	return false
}

type Rucksack struct {
	a string
	b string
}

func parseInput(input string) []Rucksack {
	parsed := make([]Rucksack, 0)
	for _, rucksack := range strings.Split(input, "\n") {
		half := len(rucksack) / 2
		parsed = append(parsed, Rucksack{
			a: rucksack[:half],
			b: rucksack[half:],
		})
	}
	return parsed
}
