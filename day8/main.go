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

func add_if_in_bound(n, m, x, y int, antinode_map map[int]bool) bool {
	if x < 0 || x >= n || y < 0 || y >= m {
		return false
	}
	antinode_map[(x<<10)+y] = true
	return true
}

func part1(s string) int {
	antenna_map := map[rune][][]int{}
	antinode_map := map[int]bool{}
	lines := strings.Split(s, "\n")
	n, m := len(lines), len(lines[0])

	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				antenna_map[r] = append(antenna_map[r], []int{i, j})
			}
		}
	}

	for _, antennas := range antenna_map {
		if len(antennas) <= 1 {
			continue
		}
		for i := 1; i < len(antennas); i++ {
			for j := 0; j < i; j++ {
				x1, y1 := antennas[i][0]*2-antennas[j][0], antennas[i][1]*2-antennas[j][1]
				add_if_in_bound(n, m, x1, y1, antinode_map)
				x2, y2 := antennas[j][0]*2-antennas[i][0], antennas[j][1]*2-antennas[i][1]
				add_if_in_bound(n, m, x2, y2, antinode_map)
			}
		}
	}

	return len(antinode_map)
}

func part2(s string) int {
	antenna_map := map[rune][][]int{}
	antinode_map := map[int]bool{}
	lines := strings.Split(s, "\n")
	n, m := len(lines), len(lines[0])

	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				antenna_map[r] = append(antenna_map[r], []int{i, j})
			}
		}
	}

	for _, antennas := range antenna_map {
		if len(antennas) <= 1 {
			continue
		}
		for i := 1; i < len(antennas); i++ {
			for j := 0; j < i; j++ {
				x, y := antennas[i][0], antennas[i][1]
				for add_if_in_bound(n, m, x, y, antinode_map) {
					x += antennas[i][0] - antennas[j][0]
					y += antennas[i][1] - antennas[j][1]
				}
				x, y = antennas[j][0], antennas[j][1]
				for add_if_in_bound(n, m, x, y, antinode_map) {
					x += antennas[j][0] - antennas[i][0]
					y += antennas[j][1] - antennas[i][1]
				}
			}
		}
	}

	return len(antinode_map)
}
