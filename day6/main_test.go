package main

import (
	"testing"
)

var example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  41,
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

var tricky = `.....#......
.#........#.
......#.....
............
...#........
#..........#
.........#..
............
............
..#..^......
.....#......`

var tricky2 = `....#....
.#.....#.
.....#...
.....#...
#.......#
......#..
...#.....
..^....#.`

var guard = `##..
^..#
..#.`

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  6,
		},
		{
			name:  "tricky",
			input: tricky,
			want:  5,
		},
		{
			name:  "tricky2",
			input: tricky2,
			want:  0,
		},
		{
			name:  "guard",
			input: guard,
			want:  0,
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
