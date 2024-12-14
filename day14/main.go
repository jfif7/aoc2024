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
		ans := part1(input, 101, 103)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(s string, w, h int) int64 {
	lines := strings.Split(s, "\n")
	quad := make([]int, 4)
	time := 100
	for _, line := range lines {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		px += vx * time
		py += vy * time
		px = ((px % w) + w) % w
		py = ((py % h) + h) % h
		q := 0
		if px == w/2 || py == h/2 {
			continue
		}
		if px < w/2 {
			q += 1
		}
		if py < h/2 {
			q += 2
		}
		quad[q]++
	}
	return int64(quad[0]) * int64(quad[1]) * int64(quad[2]) * int64(quad[3])
}

func part2(s string) int64 {
	lines := strings.Split(s, "\n")
	wx, wy := 101, 103
	pxs := make([]int, 0)
	pys := make([]int, 0)
	vxs := make([]int, 0)
	vys := make([]int, 0)
	for _, line := range lines {
		var px, py, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		pxs = append(pxs, px)
		pys = append(pys, py)
		vxs = append(vxs, vx)
		vys = append(vys, vy)
	}
	for k := 0; k < 100000; k += 1 {
		pos := make([][]int, wy)
		for i := range pos {
			pos[i] = make([]int, wx)
		}
		for i := range pxs {
			px := pxs[i] + vxs[i]*k
			py := pys[i] + vys[i]*k
			px = ((px % wx) + wx) % wx
			py = ((py % wy) + wy) % wy
			pos[py][px] = 1
		}
		if check(pos) {
			fmt.Printf("%v\n\n\n\n\n\n\n\n\n\n\n\n\n", k)
			for i := range pos {
				for j := range pos[0] {
					if pos[i][j] == 0 {
						fmt.Printf(" ")
					} else {
						fmt.Printf("■")
					}
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}
	return 0
}

func check(pos [][]int) bool {
	count := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 101; j++ {
			if pos[i][j] == 1 && pos[i+1][j] == 1 && pos[i+2][j] == 1 && pos[i+3][j] == 1 {
				count++
			}
		}
	}
	return count >= 2
}
