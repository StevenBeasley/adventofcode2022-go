package main

import (
	_ "embed"
	"fmt"
	"log"
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

	resultOne := partOne(parsedInput)
	resultTwo := partTwo(parsedInput)
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func partTwo(parsedInput []Pairs) (numberOfOverlaps int) {
	for _, pair := range parsedInput {
		if eitherOverlaps(pair) {
			numberOfOverlaps += 1
		}
	}
	return
}

func eitherOverlaps(pair Pairs) bool {
	a := pair[0]
	b := pair[1]
	if eitherContains(pair) {
		return true
	}
	if between(a.start, a.end, b.start) {
		return true
	}
	if between(a.start, a.end, b.end) {
		return true
	}
	return false
}

func between(start, finish, num int) bool {
	return start <= num && num <= finish
}

func partOne(parsedInput []Pairs) (numberOfContains int) {
	for _, pair := range parsedInput {
		if eitherContains(pair) {
			numberOfContains += 1
		}
	}
	return
}

func eitherContains(pair Pairs) bool {
	if contains(pair[0], pair[1]) {
		return true
	} else if contains(pair[1], pair[0]) {
		return true
	}
	return false
}

func contains(elf1, elf2 IdRange) bool {
	return (elf1.start <= elf2.start && elf1.end >= elf2.end)
}

type IdRange struct {
	start int
	end   int
}
type Pairs map[int]IdRange

func parseInput(input string) (pairs []Pairs) {
	for _, row := range strings.Split(input, "\n") {
		pair := Pairs{}
		splitRow := strings.Split(row, ",")
		for i, elf := range splitRow {
			splitElf := strings.Split(elf, "-")
			start, err := strconv.Atoi(splitElf[0])
			if err != nil {
				log.Fatal(err)
			}
			end, err := strconv.Atoi(splitElf[1])
			if err != nil {
				log.Fatal(err)
			}
			pair[i] = IdRange{
				start,
				end,
			}
		}
		pairs = append(pairs, pair)
	}
	return
}
