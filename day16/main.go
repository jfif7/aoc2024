package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"

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

type Record struct {
	x, y, ori, cost int
}

type RecordWithHistory struct {
	r, hist Record
}

func part1(s string) int64 {
	_, _, board := utils.ReadInputAsBoard(s)
	var sx, sy int
	visited := make([][][]bool, 0)
	for i, line := range board {
		visited = append(visited, make([][]bool, 0))
		for j, r := range line {
			visited[i] = append(visited[i], make([]bool, 4))
			if r == 'S' {
				sx, sy = i, j
			}
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	q := []Record{{sx, sy, 1, 0}}
	for {
		rec := q[0]
		q = q[1:]
		// fmt.Println(rec)
		if visited[rec.x][rec.y][rec.ori] {
			continue
		} else {
			visited[rec.x][rec.y][rec.ori] = true
		}
		if board[rec.x][rec.y] == 'E' {
			return int64(rec.cost)
		}
		for i := -1; i <= 1; i++ {
			ori := (rec.ori + i + 4) % 4
			dx, dy := dirs[ori][0], dirs[ori][1]
			nx, ny := rec.x+dx, rec.y+dy
			if board[nx][ny] != '#' {
				cost := rec.cost + 1
				if i != 0 {
					cost += 1000
				}
				q = append(q, Record{nx, ny, ori, cost})
				slices.SortFunc(q, func(a Record, b Record) int { return a.cost - b.cost })
			}
		}
	}
}

func part2(s string) int64 {
	_, _, board := utils.ReadInputAsBoard(s)
	var sx, sy int
	visited := make([][][]int, 0)
	backtrack := make(map[Record][]Record)
	for i, line := range board {
		visited = append(visited, make([][]int, 0))
		for j, r := range line {
			visited[i] = append(visited[i], make([]int, 4))
			if r == 'S' {
				sx, sy = i, j
			}
		}
	}
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	q := []RecordWithHistory{{Record{sx, sy, 1, 1}, Record{}}}
	var end Record
	for {
		rec := q[0]
		q = q[1:]
		// fmt.Println(rec)
		r := rec.r
		if visited[r.x][r.y][r.ori] > 0 {
			if visited[r.x][r.y][r.ori] == r.cost {
				backtrack[r] = append(backtrack[r], rec.hist)
			}
			continue
		} else {
			visited[r.x][r.y][r.ori] = r.cost
			backtrack[r] = append(backtrack[r], rec.hist)
		}
		if board[r.x][r.y] == 'E' {
			end = r
			break
		}
		for i := -1; i <= 1; i++ {
			ori := (r.ori + i + 4) % 4
			dx, dy := dirs[ori][0], dirs[ori][1]
			nx, ny := r.x+dx, r.y+dy
			if board[nx][ny] != '#' {
				cost := r.cost + 1
				if i != 0 {
					cost += 1000
				}
				q = append(q, RecordWithHistory{Record{nx, ny, ori, cost}, r})
				slices.SortFunc(q, func(a RecordWithHistory, b RecordWithHistory) int { return a.r.cost - b.r.cost })
			}
		}
	}
	backtrack_q := []Record{end}
	backtrack_map := make(map[Record]bool)
	for len(backtrack_q) > 0 {
		// fmt.Println(backtrack_q[0])
		record := backtrack_q[0]
		backtrack_q = backtrack_q[1:]
		if record.x == 0 && record.y == 0 {
			continue
		}
		backtrack_map[Record{record.x, record.y, 0, 0}] = true
		backtrack_q = append(backtrack_q, backtrack[record]...)
	}
	return int64(len(backtrack_map))
}
