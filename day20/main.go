package main

import (
	_ "embed"
	"flag"
	"fmt"

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
		ans := part1(input, 100)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input, 100)
		fmt.Println("Output:", ans)
	}
}

func part1(s string, threshold int) int {
	n, m, board := utils.ReadInputAsBoardAddWall(s, '#')
	v := make([][]int, n)
	for i := range v {
		v[i] = make([]int, m)
	}
	q := make([][]int, 0)
	for i, line := range board {
		for j, r := range line {
			if r == 'S' {
				q = append(q, []int{i, j})
				v[i][j] = 1
			}
			if r == 'E' {
				board[i][j] = '.'
			}
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	count := 0
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, dir := range dirs {
			nx, ny := p[0]+dir[0], p[1]+dir[1]
			sx, sy := p[0]+dir[0]*2, p[1]+dir[1]*2
			// gogo
			if board[nx][ny] == '.' && v[nx][ny] == 0 {
				q = append(q, []int{nx, ny})
				v[nx][ny] = v[p[0]][p[1]] + 1
			}
			// find shortcut
			if board[nx][ny] == '#' && v[sx][sy] != 0 {
				if diff := v[p[0]][p[1]] - v[sx][sy]; diff >= threshold+2 {
					count++
				}
			}
		}
	}
	return count
}

func part2(s string, threshold int) int {
	n, m, board := utils.ReadInputAsBoard(s)
	v := make([][]int, n)
	for i := range v {
		v[i] = make([]int, m)
	}
	q := make([][]int, 0)
	for i, line := range board {
		for j, r := range line {
			if r == 'S' {
				q = append(q, []int{i, j})
				v[i][j] = 1
			}
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	count := 0
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, dir := range dirs {
			nx, ny := p[0]+dir[0], p[1]+dir[1]
			// gogo
			if (board[nx][ny] == '.' || board[nx][ny] == 'E') && v[nx][ny] == 0 {
				q = append(q, []int{nx, ny})
				v[nx][ny] = v[p[0]][p[1]] + 1
			}
		}
		// find shortcuts
		count += findShortcuts(board, v, p, threshold)
	}
	return count
}

func findShortcuts(board [][]rune, steps [][]int, ori []int, threshold int) int {
	q := [][]int{ori}
	n, m := len(board), len(board[0])
	v := make([][]int, 41)
	for i := range v {
		v[i] = make([]int, 41)
	}
	v[20][20] = 1
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	count := 0
	ori_step := steps[ori[0]][ori[1]]

	for len(q) > 0 {
		p := q[0]
		x, y := p[0], p[1]
		vv := v[x-ori[0]+20][y-ori[1]+20]
		// fmt.Println(x-ori[0], y-ori[1], vv)
		q = q[1:]
		// check if shortcut
		if step := steps[x][y]; step > 0 {
			if ori_step-step-vv+1 >= threshold {
				count++
			}
		}
		// gogo
		if vv >= 21 {
			continue
		}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || ny < 0 || nx >= n || ny >= m {
				continue
			}
			ix, iy := nx-ori[0]+20, ny-ori[1]+20
			if v[ix][iy] > 0 {
				continue
			}
			v[ix][iy] = vv + 1
			q = append(q, []int{nx, ny})
		}
	}
	return count
}
