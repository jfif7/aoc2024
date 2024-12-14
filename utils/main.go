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
