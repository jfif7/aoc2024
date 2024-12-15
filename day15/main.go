package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"github.com/jfif7/aoc2024/utils"
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

func part1(s string) int64 {
	regions := strings.Split(s, "\n\n")
	_, _, board := utils.ReadInputAsBoard(regions[0])
	var pos []int
	for i, line := range board {
		for j, r := range line {
			if r == '@' {
				pos = []int{i, j}
				break
			}
		}
		if len(pos) == 2 {
			break
		}
	}
	for _, r := range regions[1] {
		switch r {
		case '^':
			if move(board, pos[0], pos[1], -1, 0) {
				pos[0]--
			}
		case '<':
			if move(board, pos[0], pos[1], 0, -1) {
				pos[1]--
			}
		case 'v':
			if move(board, pos[0], pos[1], 1, 0) {
				pos[0]++
			}
		case '>':
			if move(board, pos[0], pos[1], 0, 1) {
				pos[1]++
			}
		default:
		}
	}
	sum := 0
	for i, line := range board {
		for j, r := range line {
			fmt.Printf("%c", r)
			if r == 'O' {
				sum += 100*i + j
			}
		}
		fmt.Printf("\n")
	}
	return int64(sum)
}

func move(board [][]rune, x, y, dx, dy int) bool {
	nx, ny := x+dx, y+dy
	switch board[nx][ny] {
	case '#':
		return false
	case '.':
		board[nx][ny] = board[x][y]
		board[x][y] = '.'
		return true
	case 'O':
		ok := move(board, nx, ny, dx, dy)
		if ok {
			board[nx][ny] = board[x][y]
			board[x][y] = '.'
		}
		return ok
	default:
		return false
	}
}

func part2(s string) int64 {
	regions := strings.Split(s, "\n\n")
	_, _, board := utils.ReadInputAsBoard(regions[0])
	board = widen(board)
	var pos []int
	for i, line := range board {
		for j, r := range line {
			if r == '@' {
				pos = []int{i, j}
				break
			}
		}
		if len(pos) == 2 {
			break
		}
	}
	for _, r := range regions[1] {
		switch r {
		case '^':
			if move2(board, pos[0], pos[1], -1, 0, false) {
				pos[0]--
			}
		case '<':
			if move2(board, pos[0], pos[1], 0, -1, false) {
				pos[1]--
			}
		case 'v':
			if move2(board, pos[0], pos[1], 1, 0, false) {
				pos[0]++
			}
		case '>':
			if move2(board, pos[0], pos[1], 0, 1, false) {
				pos[1]++
			}
		default:
		}
	}
	sum := 0
	for i, line := range board {
		for j, r := range line {
			fmt.Printf("%c", r)
			if r == '[' {
				sum += 100*i + j
			}
		}
		fmt.Printf("\n")
	}
	return int64(sum)
}

func widen(board [][]rune) [][]rune {
	new_board := make([][]rune, 0, len(board))
	for i, line := range board {
		new_board = append(new_board, make([]rune, 0, 2*len(line)))
		for _, r := range line {
			switch r {
			case '#':
				new_board[i] = append(new_board[i], '#', '#')
			case 'O':
				new_board[i] = append(new_board[i], '[', ']')
			case '.':
				new_board[i] = append(new_board[i], '.', '.')
			case '@':
				new_board[i] = append(new_board[i], '@', '.')
			}
		}
	}
	return new_board
}

func move2(board [][]rune, x, y, dx, dy int, dry bool) bool {
	nx, ny := x+dx, y+dy
	switch true {
	case board[nx][ny] == '#':
		return false
	case board[nx][ny] == '.':
		if !dry {
			board[nx][ny] = board[x][y]
			board[x][y] = '.'
		}
		return true
	case (board[nx][ny] == '[' || board[nx][ny] == ']') && dy != 0:
		ok := move2(board, nx, ny, dx, dy, dry)
		if ok && !dry {
			board[nx][ny] = board[x][y]
			board[x][y] = '.'
		}
		return ok
	case board[nx][ny] == '[' && dx != 0:
		ok := move2(board, nx, ny, dx, dy, true) && move2(board, nx, ny+1, dx, dy, true)
		if ok && !dry {
			move2(board, nx, ny, dx, dy, false)
			move2(board, nx, ny+1, dx, dy, false)
			board[nx][ny] = board[x][y]
			board[x][y] = '.'
		}
		return ok
	case board[nx][ny] == ']' && dx != 0:
		ok := move2(board, nx, ny, dx, dy, true) && move2(board, nx, ny-1, dx, dy, true)
		if ok && !dry {
			move2(board, nx, ny, dx, dy, false)
			move2(board, nx, ny-1, dx, dy, false)
			board[nx][ny] = board[x][y]
			board[x][y] = '.'
		}
		return ok
	default:
		return false
	}
}
