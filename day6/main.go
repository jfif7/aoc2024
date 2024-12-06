package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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

func part1(s string) int {
	board := make([][]rune, 0)
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var pos []int
	for i, line := range strings.Split(s, "\n") {
		board = append(board, make([]rune, 0, len(line)))
		for j, r := range line {
			board[i] = append(board[i], r)
			if board[i][j] == '^' {
				pos = []int{i, j}
				board[i][j] = 'X'
			}
		}
	}
	n, m := len(board), len(board[0])
	dir := 0
	sum := 1
	for {
		new_x, new_y := pos[0]+dirs[dir][0], pos[1]+dirs[dir][1]
		if new_x < 0 || new_x >= n ||
			new_y < 0 || new_y >= m {
			break
		}
		switch board[new_x][new_y] {
		case '.':
			board[new_x][new_y] = 'X'
			sum += 1
			pos = []int{new_x, new_y}
		case 'X':
			pos = []int{new_x, new_y}
		case '#':
			dir = (dir + 1) % 4
		}
	}
	return sum
}

func test_loop(ori [][]int, dir int, pos []int) int {
	board := make([][]int, len(ori))
	for i := range board {
		board[i] = make([]int, len(ori[i]))
		copy(board[i], ori[i])
	}

	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	n, m := len(board), len(board[0])
	for {
		new_x, new_y := pos[0]+dirs[dir][0], pos[1]+dirs[dir][1]
		if new_x < 0 || new_x >= n ||
			new_y < 0 || new_y >= m {
			return 0
		}
		tile := board[new_x][new_y]
		switch {
		case tile >= 0:
			board[new_x][new_y] += 1
			pos = []int{new_x, new_y}
			if board[new_x][new_y] >= 5 {
				return 1
			}
		case tile == -1:
			dir = (dir + 1) % 4
		}
	}
}

func part2(s string) int {
	board := make([][]int, 0)
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var pos []int
	for i, line := range strings.Split(s, "\n") {
		board = append(board, make([]int, 0, len(line)))
		for j, r := range line {
			switch r {
			case '^':
				board[i] = append(board[i], 0)
				pos = []int{i, j}
			case '.':
				board[i] = append(board[i], 0)
			case '#':
				board[i] = append(board[i], -1)
			}
		}
	}
	n, m := len(board), len(board[0])
	dir := 0
	yes := make(map[int]bool)
	ori_pos := []int{pos[0], pos[1]}
	yes[(pos[0]<<8)+pos[1]] = true
	for {
		new_x, new_y := pos[0]+dirs[dir][0], pos[1]+dirs[dir][1]
		if new_x < 0 || new_x >= n ||
			new_y < 0 || new_y >= m {
			break
		}
		tile := board[new_x][new_y]
		switch {
		case tile == 0:
			board[new_x][new_y] = -1
			if test_loop(board, 0, ori_pos) == 1 {
				yes[(new_x<<8)+new_y] = true
				// fmt.Println(new_x, new_y)
			}
			board[new_x][new_y] = 0
			pos = []int{new_x, new_y}
		case tile == -1:
			dir = (dir + 1) % 4
		default:
			log.Fatalf("what?")
		}
	}
	return len(yes) - 1
}
