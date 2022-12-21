package main

import (
	_ "embed"
	"testing"
)

func Test_partTwo(t *testing.T) {
	state, moves := parseInput(input)
	type args struct {
		state State
		moves []Move
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AOC input",
			args: args{
				state,
				moves,
			},
			want: "BNTZFPMMW",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.state, tt.args.moves); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partOne(t *testing.T) {
	state, moves := parseInput(input)
	type args struct {
		state State
		moves []Move
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "AOC input",
			args: args{
				state,
				moves,
			},
			want: "PSNRGBTFT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.state, tt.args.moves); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
