package main

import (
	_ "embed"
	"reflect"
	"testing"
)

func Test_between(t *testing.T) {
	type args struct {
		start  int
		finish int
		num    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "number is between others",
			args: args{
				start:  1,
				finish: 10,
				num:    5,
			},
			want: true,
		},
		{
			name: "number is between others",
			args: args{
				start:  1,
				finish: 10,
				num:    15,
			},
			want: false,
		},
		{
			name: "number is between others",
			args: args{
				start:  1,
				finish: 10,
				num:    1,
			},
			want: true,
		},
		{
			name: "number is between others",
			args: args{
				start:  1,
				finish: 10,
				num:    10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := between(tt.args.start, tt.args.finish, tt.args.num); got != tt.want {
				t.Errorf("between() = %v, want %v", got, tt.want)
			}
		})
	}
}

var simpleInput = "1-5,2-4\n1-10,9-12\n1-12,13-20\n2-2,1-4"

func Test_partTwo(t *testing.T) {
	type args struct {
		parsedInput []Pairs
	}
	tests := []struct {
		name                 string
		args                 args
		wantNumberOfOverlaps int
	}{
		{
			name: "simple input",
			args: args{
				parsedInput: parseInput(simpleInput),
			},
			wantNumberOfOverlaps: 3,
		},
		{
			name: "AOC input",
			args: args{
				parsedInput: parseInput(input),
			},
			wantNumberOfOverlaps: 909,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumberOfOverlaps := partTwo(tt.args.parsedInput); gotNumberOfOverlaps != tt.wantNumberOfOverlaps {
				t.Errorf("partTwo() = %v, want %v", gotNumberOfOverlaps, tt.wantNumberOfOverlaps)
			}
		})
	}
}

func Test_eitherOverlaps(t *testing.T) {
	type args struct {
		pair Pairs
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "start overlaps",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 5},
					1: {start: 4, end: 10},
				},
			},
			want: true,
		},
		{
			name: "end overlaps",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 5},
					1: {start: 0, end: 2},
				},
			},
			want: true,
		},
		{
			name: "full overlaps",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 5},
					1: {start: 3, end: 4},
				},
			},
			want: true,
		},
		{
			name: "full overlaps same number",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 5},
					1: {start: 2, end: 2},
				},
			},
			want: true,
		},
		{
			name: "not overlaps",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 5},
					1: {start: 6, end: 20},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eitherOverlaps(tt.args.pair); got != tt.want {
				t.Errorf("eitherOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partOne(t *testing.T) {
	type args struct {
		parsedInput []Pairs
	}
	tests := []struct {
		name                 string
		args                 args
		wantNumberOfContains int
	}{
		{
			name: "AOC input",
			args: args{
				parsedInput: parseInput(input),
			},
			wantNumberOfContains: 518,
		},
		{
			name: "simple input",
			args: args{
				parsedInput: parseInput(simpleInput),
			},
			wantNumberOfContains: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumberOfContains := partOne(tt.args.parsedInput); gotNumberOfContains != tt.wantNumberOfContains {
				t.Errorf("partOne() = %v, want %v", gotNumberOfContains, tt.wantNumberOfContains)
			}
		})
	}
}

func Test_eitherContains(t *testing.T) {
	type args struct {
		pair Pairs
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a in b",
			args: args{
				pair: map[int]IdRange{
					0: {start: 3, end: 5},
					1: {start: 1, end: 20},
				},
			},
			want: true,
		},
		{
			name: "b in a",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 50},
					1: {start: 10, end: 20},
				},
			},
			want: true,
		},
		{
			name: "b overlap a",
			args: args{
				pair: map[int]IdRange{
					0: {start: 1, end: 50},
					1: {start: 10, end: 70},
				},
			},
			want: false,
		},
		{
			name: "a overlap b",
			args: args{
				pair: map[int]IdRange{
					0: {start: 11, end: 80},
					1: {start: 10, end: 70},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eitherContains(tt.args.pair); got != tt.want {
				t.Errorf("eitherContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		elf1 IdRange
		elf2 IdRange
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 in 2",
			args: args{
				elf1: IdRange{
					start: 5,
					end:   10,
				},
				elf2: IdRange{
					start: 1,
					end:   12,
				},
			},
			want: false,
		},
		{
			name: "1 next to 2",
			args: args{
				elf1: IdRange{
					start: 5,
					end:   10,
				},
				elf2: IdRange{
					start: 11,
					end:   15,
				},
			},
			want: false,
		},
		{
			name: "2 in 1",
			args: args{
				elf1: IdRange{
					start: 5,
					end:   10,
				},
				elf2: IdRange{
					start: 7,
					end:   9,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.elf1, tt.args.elf2); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name      string
		args      args
		wantPairs []Pairs
	}{
		{
			name: "simple input",
			args: args{
				input: simpleInput,
			},
			wantPairs: []Pairs{
				map[int]IdRange{
					0: {start: 1, end: 5},
					1: {start: 2, end: 4},
				},
				map[int]IdRange{
					0: {start: 1, end: 10},
					1: {start: 9, end: 12},
				},
				map[int]IdRange{
					0: {start: 1, end: 12},
					1: {start: 13, end: 20},
				},
				map[int]IdRange{
					0: {start: 2, end: 2},
					1: {start: 1, end: 4},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPairs := parseInput(tt.args.input); !reflect.DeepEqual(gotPairs, tt.wantPairs) {
				t.Errorf("parseInput() = %v, want %v", gotPairs, tt.wantPairs)
			}
		})
	}
}
