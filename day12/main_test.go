package main

import (
	"testing"
)

var example = `AAAA
BBCD
BBCC
EEEC`

var example_2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example",
			input: example,
			want:  140,
		},
		{
			name:  "example_2",
			input: example_2,
			want:  772,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example",
			input: example,
			want:  80,
		},
		{
			name:  "example_2",
			input: example_2,
			want:  436,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
