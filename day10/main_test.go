package main

import (
	"testing"
)

// 0..111....22222
// 022111222
var example_easy = `12345`

var example = `2333133121414131402`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example",
			input: example,
			want:  1928,
		},
		{
			name:  "example_easy",
			input: example_easy,
			want:  60,
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
			want:  2858,
		},
		{
			name:"example_easy",
			input: example_easy,
			want: 132,
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
