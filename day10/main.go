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

func explore(board, visited [][]int, x, y int) int {
	n, m := len(board), len(board[0])
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	if board[x][y] == 9 {
		return 1
	}
	sum := 0
	for _, dir := range dirs {
		newx, newy := x+dir[0], y+dir[1]
		if newx < 0 || newx >= n || newy < 0 || newy >= m {
			continue
		}
		if visited[newx][newy] == 1 {
			continue
		}
		if board[newx][newy] == board[x][y]+1 {
			visited[newx][newy] += visited[x][y]
			sum += explore(board, visited, newx, newy)
		}
	}
	return sum
}

func reset_visit(v [][]int, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			v[i][j] = 0
		}
	}
}

func part1(s string) int64 {
	lines := strings.Split(s, "\n")
	board := make([][]int, len(lines))
	visited := make([][]int, len(lines))
	for i, line := range lines {
		for _, r := range line {
			board[i] = append(board[i], int(r-'0'))
			visited[i] = append(visited[i], 0)
		}
	}
	n, m := len(board), len(board[0])
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 0 {
				reset_visit(visited, n, m)
				sum += explore(board, visited, i, j)
			}
		}
	}
	return int64(sum)
}

func part2(s string) int64 {
	lines := strings.Split(s, "\n")
	board := make([][]int, len(lines))
	visited := make([][]int, len(lines))
	for i, line := range lines {
		for _, r := range line {
			board[i] = append(board[i], int(r-'0'))
			visited[i] = append(visited[i], 0)
		}
	}
	n, m := len(board), len(board[0])
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 0 {
				reset_visit(visited, n, m)
				sum += explore_bfs(board, visited, i, j)
			}
		}
	}
	return int64(sum)
}

func explore_bfs(board, visited [][]int, ori_x, ori_y int) int {
	n, m := len(board), len(board[0])
	q := make([][]int, 0)
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	q = append(q, []int{ori_x, ori_y})
	visited[ori_x][ori_y] = 1
	sum := 0
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		x, y := v[0], v[1]
		if board[x][y] == 9 {
			sum += visited[x][y]
			continue
		}
		for _, dir := range dirs {
			newx, newy := x+dir[0], y+dir[1]
			if newx < 0 || newx >= n || newy < 0 || newy >= m {
				continue
			}
			if board[newx][newy] == board[x][y]+1 {
				if visited[newx][newy] == 0 {
					q = append(q, []int{newx, newy})
				}
				visited[newx][newy] += visited[x][y]
			}
		}
	}
	return sum
}
