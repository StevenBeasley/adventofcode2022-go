package main

import (
	_ "embed"
	"reflect"
	"testing"
)

var parsed = []Play{
	{
		opponent: "B",
		player:   "Z",
	},
	{
		opponent: "B",
		player:   "X",
	},
	{
		opponent: "C",
		player:   "Y",
	},
	{
		opponent: "B",
		player:   "Y",
	},
	{
		opponent: "B",
		player:   "Y",
	},
	{
		opponent: "A",
		player:   "X",
	},
	{
		opponent: "A",
		player:   "X",
	},
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantParsed []Play
	}{
		{
			name:       "Normal Input",
			args:       args{"B Z\nB X\nC Y\nB Y\nB Y\nA X\nA X"},
			wantParsed: parsed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParsed := parseInput(tt.args.input); !reflect.DeepEqual(gotParsed, tt.wantParsed) {
				t.Errorf("parseInput() = %v, want %v", gotParsed, tt.wantParsed)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		input []Play
	}
	tests := []struct {
		name           string
		args           args
		wantTotalScore int
	}{
		{
			name: "AOC input",
			args: args{
				input: parseInput(input),
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
			if gotTotalScore := partTwo(tt.args.input); gotTotalScore != tt.wantTotalScore {
				t.Errorf("partTwo() = %v, want %v", gotTotalScore, tt.wantTotalScore)
			}
		})
	}
}

func Test_partOne(t *testing.T) {
	type args struct {
		input []Play
	}
	tests := []struct {
		name           string
		args           args
		wantTotalScore int
	}{
		{
			name: "AOC input",
			args: args{
				input: parseInput(input),
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
			if gotTotalScore := partOne(tt.args.input); gotTotalScore != tt.wantTotalScore {
				t.Errorf("partOne() = %v, want %v", gotTotalScore, tt.wantTotalScore)
			}
		})
	}
}
