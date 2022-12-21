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
	state, moves := parseInput(input)

	resultOne := partOne(state, moves)
	resultTwo := partTwo(state, moves)

	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func partTwo(state State, moves []Move) string {
	// take a copy of state
	stateCopy := createCopyOfState(state)

	for _, move := range moves {
		makeStackMove(stateCopy, move)
	}
	return getTopOfEachStack(stateCopy)
}

func makeStackMove(state State, move Move) {
	removedCrate := removeXfromSlice(move.number, &state[move.from])
	// add crate(s) to the to slice
	state[move.to] = append(state[move.to], removedCrate...)
}

func partOne(state State, moves []Move) string {
	// take a copy of state
	stateCopy := createCopyOfState(state)

	for _, move := range moves {
		makeMove(&stateCopy, move)
	}

	return getTopOfEachStack(stateCopy)
}
func createCopyOfState(state State) State {
	stateCopy := make(State, len(state))
	for i, stack := range state {
		stackCopy := make([]rune, 0)
		stackCopy = append(stackCopy, stack...)
		stateCopy[i] = stackCopy
	}
	return stateCopy
}

func getTopOfEachStack(state State) string {
	letters := make([]string, 0)
	for _, stack := range state {
		topLetter := stack[len(stack)-1:][0]
		letters = append(letters, string(topLetter))
	}
	return strings.Join(letters, "")
}

func makeMove(state *State, move Move) {
	for i := 1; i <= move.number; i++ {
		removedCrate := removeXfromSlice(1, &(*state)[move.from])
		// add crate to the to slice
		(*state)[move.to] = append((*state)[move.to], removedCrate...)
		// fmt.Println(removedCrate)
	}
}

func removeXfromSlice(i int, r *[]rune) []rune {
	lengthOfSlice := len((*r))
	removedCrate := (*r)[lengthOfSlice-i:]
	*r = (*r)[:lengthOfSlice-i]
	return removedCrate
}

func parseInput(input string) (State, []Move) {
	stateInput := make([]string, 0)
	moveInput := make([]string, 0)

	upToMoves := false

	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			upToMoves = true
			continue
		}
		if !upToMoves {
			stateInput = append(stateInput, row)
		} else {
			moveInput = append(moveInput, row)
		}
	}
	state := processState(stateInput)
	moves := processMoves(moveInput)
	return state, moves
}

type Move struct {
	number int
	from   int
	to     int
}

func processMoves(moveInput []string) (moves []Move) {
	for _, row := range moveInput {
		moves = append(moves, extractMoveFromString(row))
	}
	return
}
func extractMoveFromString(row string) (move Move) {
	splitRow := strings.Split(row, " ")
	number, err := strconv.Atoi(splitRow[1])
	if err != nil {
		log.Fatal(err)
	}
	from, err := strconv.Atoi(splitRow[3])
	if err != nil {
		log.Fatal(err)
	}
	to, err := strconv.Atoi(splitRow[5])
	if err != nil {
		log.Fatal(err)
	}
	move = Move{
		number: number,
		from:   from - 1,
		to:     to - 1,
	}
	return
}

type State [][]rune

func processState(stateInput []string) State {
	// remove last row, is numbers
	stateInput = stateInput[:len(stateInput)-1]
	stateColumn := len(getLettersFromRow(stateInput[len(stateInput)-1]))
	state := make(State, stateColumn)

	// go over the slice in reverse
	for i := len(stateInput) - 1; i >= 0; i-- {
		row := getLettersFromRow(stateInput[i])
		for j, letter := range row {
			if letter != rune(' ') {
				state[j] = append(state[j], letter)
			}
		}
	}
	return state
}

func getLettersFromRow(s string) (letters []rune) {
	for i := 1; i <= len(s); i += 4 {
		letters = append(letters, rune(s[i]))
	}
	return
}
