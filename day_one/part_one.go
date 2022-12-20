package main

import (
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strconv"
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

	elves := groupElves(parsedInput)
	for _, elf := range elves {
		fmt.Println(elf)
	}

	elf := findElfWithMostest(elves)
	fmt.Println("Elf with the mostest: ", elf)

	topThreeElves := findThreeElvesWithTheMostest(elves)

	fmt.Println("Top three elves: ", topThreeElves)
}

func findThreeElvesWithTheMostest(elves []int) []int {
	elves = sortElves(elves)
	return elves[:3]
}

func parseInput(input string) []int {
	parsed := make([]int, 0)
	for _, row := range strings.Split(input, "\n") {
		if row != "" {
			i, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal(err)
			}
			parsed = append(parsed, i)
		} else {
			parsed = append(parsed, 0)
		}
	}
	return parsed
}

func findElfWithMostest(elves []int) int {
	elves = sortElves(elves)
	return elves[0]
}

func sortElves(elves []int) []int {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	return elves
}

func groupElves(input []int) []int {
	elves := make([]int, 0)
	currentElf := 0
	for _, row := range input {
		if row == 0 {
			elves = append(elves, currentElf)
			currentElf = 0
		} else {
			currentElf = currentElf + row
		}
	}

	return elves
}
