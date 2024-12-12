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

func part1(s string) int64 {
	lines := strings.Split(s, "\n")
	board := make([][]rune, len(lines))
	visited := make([][]int, len(lines))
	for i, line := range lines {
		board[i] = make([]rune, len(line))
		visited[i] = make([]int, len(line))
		for j, r := range line {
			board[i][j] = r
		}
	}
	// area := make([]intd, 26)
	// perimeter := make([]int, 26)
	n, m := len(board), len(board[0])
	var sum int64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 0 {
				a, p := visit(board, visited, i, j)
				// area[int(board[i][j]-'A')] += a
				// perimeter[int(board[i][j]-'A')] += p
				sum += int64(a * p)
			}
		}
	}
	// for i := 0; i < 26; i++ {
	// 	sum += int64(area[i]) * int64(perimeter[i])
	// }
	return sum
}

func visit(board [][]rune, visited [][]int, x, y int) (int, int) {
	n, m := len(board), len(board[0])
	visited[x][y] = 1
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	area, perimeter := 1, 0
	for _, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		if !in_bound(n, m, nx, ny) {
			perimeter++
			continue
		}
		if board[nx][ny] == board[x][y] && visited[nx][ny] == 0 {
			a, p := visit(board, visited, nx, ny)
			area += a
			perimeter += p
		} else if board[nx][ny] != board[x][y] {
			perimeter++
			continue
		}
	}
	return area, perimeter
}

func in_bound(n, m, x, y int) bool {
	return !(x < 0 || x >= n || y < 0 || y >= m)
}

func visit2(board [][]rune, visited [][]int, x, y int) (int, int) {
	n, m := len(board), len(board[0])
	visited[x][y] = 1
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	area, perimeter := 1, 0
	for i, dir := range dirs {
		nx, ny := x+dir[0], y+dir[1]
		if !in_bound(n, m, nx, ny) {
			nxt_dir := dirs[(i+1)%4]
			nxtx, nxty := x+nxt_dir[0], y+nxt_dir[1]
			if !in_bound(n, m, nxtx, nxty) {
				perimeter++
			} else {
				if board[nxtx][nxty] == board[x][y] {

				} else {
					perimeter++
				}
			}
			continue
		}
		if board[nx][ny] == board[x][y] && visited[nx][ny] == 0 {
			a, p := visit2(board, visited, nx, ny)
			area += a
			perimeter += p
		} else if board[nx][ny] != board[x][y] {
			nxt_dir := dirs[(i+1)%4]
			nxtx, nxty := x+nxt_dir[0], y+nxt_dir[1]
			dx, dy := nx+nxt_dir[0], ny+nxt_dir[1]
			if !in_bound(n, m, nxtx, nxty) || !in_bound(n, m, dx, dy) {
				perimeter++
				continue
			} else {
				if board[nxtx][nxty] == board[x][y] && board[dx][dy] != board[x][y] {

				} else {
					perimeter++
					continue
				}
			}
		}
	}
	return area, perimeter
}

func part2(s string) int64 {
	lines := strings.Split(s, "\n")
	board := make([][]rune, len(lines))
	visited := make([][]int, len(lines))
	for i, line := range lines {
		board[i] = make([]rune, len(line))
		visited[i] = make([]int, len(line))
		for j, r := range line {
			board[i][j] = r
		}
	}
	n, m := len(board), len(board[0])
	var sum int64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == 0 {
				a, p := visit2(board, visited, i, j)
				sum += int64(a * p)
			}
		}
	}
	return sum
}
