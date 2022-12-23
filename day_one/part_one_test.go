package main_test

import (
	_ "embed"
	"strings"
	"testing"

	main "github.com/StevenBeasley/adventofcode2022-go/day_one"
	"github.com/stretchr/testify/require"
)

var elves []int

func init() {
	parsedInput, err := main.ParseFile("input.txt")
	if err != nil {
		panic(err)
	}

	elves = main.GroupElves(parsedInput)
}

func TestTopThreeElves(t *testing.T) {
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
			require.Equal(t, tt.want, main.TopThreeElves(tt.args.elves))
		})
	}
}

func TestParse(t *testing.T) {
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
			got, err := main.Parse(strings.NewReader(tt.args.input))
			require.Nil(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestTopElf(t *testing.T) {
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
			require.Equal(t, tt.want, main.TopElf(tt.args.elves))
		})
	}
}

func TestGroupElves(t *testing.T) {
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
			require.Equal(t, tt.want, main.GroupElves(tt.args.input))
		})
	}
}
