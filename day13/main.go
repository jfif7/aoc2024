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
	var cost int64 = 0
	for i := 0; i < len(lines); i += 4 {
		var xa, xb, ya, yb, x, y int
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &xa, &ya)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &xb, &yb)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &x, &y)
		det := xa*yb - xb*ya
		if det != 0 && (yb*x-xb*y)%det == 0 && (ya*x-xa*y)%det == 0 {
			cost += int64(3*(yb*x-xb*y)-ya*x+xa*y) / int64(det)
		}
	}
	return cost
}

func part2(s string) int64 {
	lines := strings.Split(s, "\n")
	var cost int64 = 0
	for i := 0; i < len(lines); i += 4 {
		var xa, xb, ya, yb, x, y int
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &xa, &ya)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &xb, &yb)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &x, &y)
		x += 10000000000000
		y += 10000000000000
		det := xa*yb - xb*ya
		if det != 0 && (yb*x-xb*y)%det == 0 && (ya*x-xa*y)%det == 0 {
			cost += int64(3*(yb*x-xb*y)-ya*x+xa*y) / int64(det)
		}
	}
	return cost
}
