package main

import (
	"testing"
)

var example = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func Test_part1(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		threshold int
	}{
		{
			name:      "example",
			input:     example,
			want:      1,
			threshold: 64,
		},
		{
			name:      "example",
			input:     example,
			want:      2,
			threshold: 40,
		},
		{
			name:      "example",
			input:     example,
			want:      2,
			threshold: 39,
		},
		{
			name:      "example",
			input:     example,
			want:      3,
			threshold: 38,
		},
		{
			name:      "example",
			input:     example,
			want:      5,
			threshold: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.threshold); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      int
		threshold int
	}{
		{
			name:      "example",
			input:     example,
			want:      3,
			threshold: 76,
		},
		{
			name:      "example",
			input:     example,
			want:      7,
			threshold: 74,
		},
		{
			name:      "example",
			input:     example,
			want:      29,
			threshold: 72,
		},
		{
			name:      "example",
			input:     example,
			want:      29,
			threshold: 71,
		},
		{
			name:      "example",
			input:     example,
			want:      41,
			threshold: 70,
		},
		{
			name:      "example",
			input:     example,
			want:      55,
			threshold: 68,
		},
		{
			name:      "example",
			input:     example,
			want:      67,
			threshold: 66,
		},
		{
			name:      "example",
			input:     example,
			want:      86,
			threshold: 64,
		},
		{
			name:      "example",
			input:     example,
			want:      106,
			threshold: 62,
		},
		{
			name:      "example",
			input:     example,
			want:      129,
			threshold: 60,
		},
		{
			name:      "example",
			input:     example,
			want:      154,
			threshold: 58,
		},
		{
			name:      "example",
			input:     example,
			want:      193,
			threshold: 56,
		},
		{
			name:      "example",
			input:     example,
			want:      222,
			threshold: 54,
		},
		{
			name:      "example",
			input:     example,
			want:      253,
			threshold: 52,
		},
		{
			name:      "example",
			input:     example,
			want:      285,
			threshold: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.threshold); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
