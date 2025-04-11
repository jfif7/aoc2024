package main

import (
	"testing"
)

var example = `029A
980A
179A
456A
379A`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example_1",
			input: `029A`,
			want:  68*29,
		},{
			name:  "example_2",
			input: `980A`,
			want:  60*980,
		},{
			name:  "example_3",
			input: `179A`,
			want:  68*179,
		},{
			name:  "example_4",
			input: `456A`,
			want:  64*456,
		},{
			name:  "example_5",
			input: `379A`,
			want:  64*379,
		},
		{
			name:  "example",
			input: example,
			want:  126384,
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
		name      string
		input     string
		want      int
		threshold int
	}{
		{
			name:      "example",
			input:     example,
			want:      154115708116294,
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
