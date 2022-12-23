package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	parsedInput, err := ParseFile(os.Args[1])
	if err != nil {
		log.Fatalf("parse input: %v", err)
	}

	elves := GroupElves(parsedInput)
	for _, elf := range elves {
		fmt.Println(elf)
	}

	elf := TopElf(elves)
	fmt.Println("Elf with the mostest: ", elf)

	topThreeElves := TopThreeElves(elves)
	fmt.Println("Top three elves: ", topThreeElves)
}

func TopThreeElves(elves []int) []int {
	m := elves[:3]
	slices.Sort(m)

	for _, elf := range elves[3:] {
		if elf > m[2] {
			copy(m, m[1:])
			m[2] = elf
			continue
		}

		if elf > m[1] {
			m[0] = m[1]
			m[1] = elf
			continue
		}

		if elf > m[0] {
			m[0] = elf
		}
	}
	m[0], m[2] = m[2], m[0]

	return m
}

func ParseFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Parse(file)
}

func Parse(r io.Reader) (nums []int, err error) {
	s := bufio.NewScanner(r)

	var line int
	for s.Scan() {
		line++

		row := strings.TrimSpace(s.Text())
		if row == "" {
			nums = append(nums, 0)
			continue
		}

		i, err := strconv.ParseInt(row, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("parse line %v: %w", line, err)
		}
		nums = append(nums, int(i))
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scan: %w", err)
	}

	return nums, nil
}

func TopElf(elves []int) int {
	m := elves[0]
	for _, elf := range elves[1:] {
		if elf > m {
			m = elf
		}
	}
	return m
}

func GroupElves(input []int) (elves []int) {
	currentElf := 0
	for _, row := range input {
		if row != 0 {
			currentElf = currentElf + row
			continue
		}

		elves = append(elves, currentElf)
		currentElf = 0
	}

	return elves
}
