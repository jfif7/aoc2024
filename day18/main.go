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
		ans := part1(input, 70, 1024)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input, 70)
		fmt.Println("Output:", ans)
	}
}

func create_board(dim int) [][]int {
	w := dim + 3
	ret := make([][]int, w)
	for i := range ret {
		ret[i] = make([]int, w)
	}
	for i, line := range ret {
		for j := range line {
			if i == 0 || j == 0 || i == w-1 || j == w-1 {
				ret[i][j] = 1
			}
		}
	}
	return ret
}

func part1(s string, dim, count int) int {
	board := create_board(dim)
	for i, line := range strings.Split(s, "\n") {
		if i >= count {
			break
		}
		var y, x int
		fmt.Sscanf(line, "%d,%d", &y, &x)
		board[x+1][y+1] = 1
	}
	for i, line := range board {
		for j := range line {
			fmt.Printf("%c", ".#O"[board[i][j]])
		}
		fmt.Printf("\n")
	}
	q := []int{0}
	b := dim + 3
	for {
		coor := q[0]
		q = q[1:]
		c, x, y := coor/b/b, (coor/b)%b, coor%b
		fmt.Println(c, x, y)
		if x == dim && y == dim {
			return c
		}
		dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if board[nx+1][ny+1] != 0 {
				continue
			}
			board[nx+1][ny+1] |= 2
			q = append(q, (c+1)*b*b+nx*b+ny)
		}
	}
}

func part2(s string, dim int) string {
	w := dim + 3
	board := create_board(dim)
	ds := utils.DisjointSet{}
	ds.Init(w * w)
	for i := 1; i < w-1; i++ {
		ds.Join(w, i*w)
		ds.Join(w, (w-1)*w+i)
		ds.Join(1, i)
		ds.Join(1, i*w+w-1)
	}
	dirs := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for _, line := range strings.Split(s, "\n") {
		var y, x int
		fmt.Sscanf(line, "%d,%d", &y, &x)
		board[x+1][y+1] = 1
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if board[nx+1][ny+1] == 1 {
				// fmt.Println(x, y, nx, ny)
				ds.Join(x*w+y+w+1, nx*w+ny+w+1)
				if ds.Same(1, w) {
					// fmt.Println(x, y)
					return fmt.Sprintf("%d,%d", y, x)
				}
			}
		}
	}
	for i, line := range board {
		for j := range line {
			// fmt.Printf("%c", ".#O"[board[i][j]])
			fmt.Printf("%2d ", (ds.Root(i*w+j) % 100))
		}
		fmt.Printf("\n")
	}
	return ""
}
