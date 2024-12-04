package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func find_mas(board [][]rune, x, y int, dir []int, match string) int {
	n, m := len(board), len(board[0])
	final_x, final_y := x+dir[0]*len(match), y+dir[1]*len(match)
	if final_x < 0 || final_x >= n || final_y < 0 || final_y >= m {
		return 0
	}
	for i, r := range match {
		if board[x+dir[0]*(i+1)][y+dir[1]*(i+1)] != r {
			return 0
		}
	}
	return 1
}

func part1(s string) int {
	board := make([][]rune, 0)
	for i, line := range strings.Split(s, "\n") {
		board = append(board, make([]rune, 0, len(line)))
		for _, r := range line {
			board[i] = append(board[i], r)
		}
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	sum := 0
	for i, line := range board {
		for j := range line {
			if board[i][j] == 'X' {
				for _, dir := range dirs {
					sum += find_mas(board, i, j, dir, "MAS")
				}
			}
		}
	}
	return sum
}

func find_mmss(board [][]rune, x, y int) int {
	dirs := [][]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}, {0, 0}}
	match := []rune{'M', 'M', 'S', 'S', 'M', 'M', 'S'}
	for offset := 0; offset < 4; offset += 1 {
		for i, dir := range dirs {
			if i == 4 {
				return 1
			}
			if board[x+dir[0]][y+dir[1]] != match[i+offset] {
				break
			}
		}
	}
	return 0
}

func part2(s string) int {
	board := make([][]rune, 0)
	for i, line := range strings.Split(s, "\n") {
		board = append(board, make([]rune, 0, len(line)))
		for _, r := range line {
			board[i] = append(board[i], r)
		}
	}
	sum := 0
	for i, line := range board[1 : len(board)-1] {
		for j := range line[1 : len(line)-1] {
			if board[i+1][j+1] == 'A' {
				sum += find_mmss(board, i+1, j+1)
			}
		}
	}
	return sum
}
