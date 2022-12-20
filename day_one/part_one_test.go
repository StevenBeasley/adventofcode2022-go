package main

import (
	_ "embed"
	"reflect"
	"testing"
)

var parsedInput []int = parseInput(input)
var elves []int = groupElves(parsedInput)

func Test_findThreeElvesWithTheMostest(t *testing.T) {
	type args struct {
		elves []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Small set",
			args: args{
				elves: []int{10, 20, 30, 40},
			},
			want: []int{40, 30, 20},
		},
		{
			name: "AOC code",
			args: args{
				elves: elves,
			},
			want: []int{70613, 68330, 66862},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThreeElvesWithTheMostest(tt.args.elves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThreeElvesWithTheMostest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Small snippet",
			args: args{
				input: "12\n23\n\n56\n76",
			},
			want: []int{12, 23, 0, 56, 76},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findElfWithMostest(t *testing.T) {
	type args struct {
		elves []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "small snippet",
			args: args{
				elves: []int{50, 10, 100, 99},
			},
			want: 100,
		}, {
			name: "AOC code",
			args: args{
				elves: elves,
			},
			want: 70613,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findElfWithMostest(tt.args.elves); got != tt.want {
				t.Errorf("findElfWithMostest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupElves(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "small snippet",
			args: args{
				input: []int{1, 2, 0, 3, 4, 0, 5, 6, 0},
			},
			want: []int{3, 7, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupElves(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupElves() = %v, want %v", got, tt.want)
			}
		})
	}
}
