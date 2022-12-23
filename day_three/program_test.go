package main

import (
	_ "embed"
	"reflect"
	"strings"
	"testing"
)

func Test_partOne(t *testing.T) {
	type args struct {
		parsedInput []Rucksack
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AOC input",
			args: args{
				parsedInput: parseInput(input),
			},
			want: 7917,
		},
		{
			name: "simple input",
			args: args{
				parsedInput: []Rucksack{
					{
						a: "abc",
						b: "ade",
					},
					{
						a: "ABC",
						b: "CDE",
					},
				},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.parsedInput); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMatchingAsciiValue(t *testing.T) {
	type args struct {
		rucksack Rucksack
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lower case a",
			args: args{
				rucksack: Rucksack{
					a: "abcde",
					b: "itsnotBisa",
				},
			},
			want: 1,
		},
		{
			name: "upper case A",
			args: args{
				rucksack: Rucksack{
					a: "Abcde",
					b: "itsnotBisA",
				},
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatchingAsciiValue(tt.args.rucksack); got != tt.want {
				t.Errorf("findMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsString(t *testing.T) {
	type args struct {
		b     []rune
		aChar rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				b:     []rune{97, 98, 99, 100, 101, 102, 103, 104},
				aChar: 100,
			},
			want: true,
		},
		{
			name: "should return false",
			args: args{
				b:     []rune{97, 98, 99, 100, 101, 102, 103, 104},
				aChar: 105,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsString(tt.args.b, tt.args.aChar); got != tt.want {
				t.Errorf("containsString() = %v, want %v", got, tt.want)
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
		want []Rucksack
	}{
		{
			name: "should split each line evenly",
			args: args{
				input: "abcd\nefgh\nijlkmnop",
			},
			want: []Rucksack{
				{
					a: "ab",
					b: "cd",
				},
				{
					a: "ef",
					b: "gh",
				},
				{
					a: "ijlk",
					b: "mnop",
				},
			},
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

func Test_getAsciiValue(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 1 for a",
			args: args{
				char: rune('a'),
			},
			want: 1,
		},
		{
			name: "should return 27 for A",
			args: args{
				char: rune('A'),
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAsciiValue(tt.args.char); got != tt.want {
				t.Errorf("getAsciiValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		parsedInput []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "AOC input",
			args: args{
				parsedInput: strings.Split(input, "\n"),
			},
			want: 2585,
		},
		{
			name: "Simple input",
			args: args{
				parsedInput: []string{"abc", "AbC", "ZQb"},
			},
			want: 2, // want b
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.parsedInput); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBadgesFromGroup(t *testing.T) {
	type args struct {
		currentGroup []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "non error input",
			args: args{
				currentGroup: []string{"AbC", "abc", "zxb"},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "error input",
			args: args{
				currentGroup: []string{"abC", "abc", "axb"},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getBadgesFromGroup(tt.args.currentGroup)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBadgesFromGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getBadgesFromGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deDupe(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "de dupe ints",
			args: args{
				items: []int{1, 2, 3, 3, 4, 5, 3},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deDupe(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deDupe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMatchingValues(t *testing.T) {
	type args struct {
		a []rune
		b []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatchingValues(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatchingValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
