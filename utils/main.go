package utils

import (
	_ "embed"
	"strings"
)

func ReadInputAsBoard(s string) (n, m int, board [][]rune) {
	lines := strings.Split(s, "\n")
	board = make([][]rune, len(lines))
	for i, line := range lines {
		board[i] = make([]rune, len(line))
		for j, r := range line {
			board[i][j] = r
		}
	}
	n, m = len(board), len(board[0])
	return
}

type DisjointSet struct {
	arr []int
}

func (ds *DisjointSet) Init(size int) {
	ds.arr = make([]int, size)
	for i := range ds.arr {
		ds.arr[i] = i
	}
}

func (ds *DisjointSet) Join(a, b int) {
	ra := ds.Root(a)
	rb := ds.Root(b)
	ds.arr[ra] = rb
}

func (ds *DisjointSet) Root(a int) int {
	if ds.arr[a] == a {
		return a
	}
	root := ds.Root(ds.arr[a])
	ds.arr[a] = root
	return root
}

func (ds *DisjointSet) Same(a, b int) bool {
	ra := ds.Root(a)
	rb := ds.Root(b)
	return ra == rb
}
