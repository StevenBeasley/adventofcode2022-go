package main_test

import (
	_ "embed"
	"strings"
	"testing"

	main "github.com/StevenBeasley/adventofcode2022-go/day_two"
	"github.com/stretchr/testify/require"
)

var parsedInput []main.Play

func init() {
	p, err := main.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedInput = p
}

var parsed = []main.Play{
	{
		Opponent: 1,
		Player:   2,
	},
	{
		Opponent: 1,
		Player:   0,
	},
	{
		Opponent: 2,
		Player:   1,
	},
	{
		Opponent: 1,
		Player:   1,
	},
	{
		Opponent: 1,
		Player:   1,
	},
	{
		Opponent: 0,
		Player:   0,
	},
	{
		Opponent: 0,
		Player:   0,
	},
}

func TestParse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantParsed []main.Play
	}{
		{
			name:       "Normal Input",
			args:       args{"B Z\nB X\nC Y\nB Y\nB Y\nA X\nA X"},
			wantParsed: parsed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := main.Parse(strings.NewReader(tt.args.input))
			require.Nil(t, err)
			require.Equal(t, tt.wantParsed, got)
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		input []main.Play
	}
	tests := []struct {
		name           string
		args           args
		wantTotalScore int
	}{
		{
			name: "AOC input",
			args: args{
				input: parsedInput,
			},
			wantTotalScore: 12767,
		}, {
			name: "Small input",
			args: args{
				input: parsed,
			},
			wantTotalScore: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantTotalScore, main.PartTwo(tt.args.input))
		})
	}
}

func TestPartOne(t *testing.T) {
	type args struct {
		input []main.Play
	}
	tests := []struct {
		name           string
		args           args
		wantTotalScore int
	}{
		{
			name: "AOC input",
			args: args{
				input: parsedInput,
			},
			wantTotalScore: 11666,
		}, {
			name: "Small input",
			args: args{
				input: parsed,
			},
			wantTotalScore: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantTotalScore, main.PartOne(tt.args.input))
		})
	}
}
