package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+

    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/

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

func toCoorNumeric(r rune) (int, int) {
	switch r {
	case 'A':
		return 3, 2
	case '0':
		return 3, 1
	case '1':
		return 2, 0
	case '2':
		return 2, 1
	case '3':
		return 2, 2
	case '4':
		return 1, 0
	case '5':
		return 1, 1
	case '6':
		return 1, 2
	case '7':
		return 0, 0
	case '8':
		return 0, 1
	case '9':
		return 0, 2
	default:
		log.Fatal("to_coor not found")
	}
	return -1, -1
}

func toCoorDirectional(r rune) (int, int) {
	switch r {
	case 'A':
		return 0, 2
	case '^':
		return 0, 1
	case 'v':
		return 1, 1
	case '<':
		return 1, 0
	case '>':
		return 1, 2
	default:
		log.Fatal("to_coor not found")
	}
	return -1, -1
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func cost_coor(coor []int, depth int, record map[string]int, is_num bool) int {
	if depth == 0 {
		return abs(coor[0] - coor[2]) + abs(coor[1]- coor[3]) + 1
	}
	var horizontal, vertical string
	dy, dx := coor[0] - coor[2], coor[1] - coor[3]
	if dy < 0 {
		vertical = strings.Repeat("^", -dy)
	} else if dy > 0 {
		vertical = strings.Repeat("v", dy)
	}
	if dx < 0 {
		horizontal = strings.Repeat("<", -dx)
	} else if dx > 0 {
		horizontal = strings.Repeat(">", dx)
	}
	if is_num && coor[0] == 3 && coor[3] == 0 {
		return cost_str(strings.Join([]string{horizontal, vertical, "A"}, ""), depth, record)
	}
	if is_num && coor[1] == 0 && coor[2] == 3 {
		return cost_str(strings.Join([]string{vertical, horizontal, "A"}, ""), depth, record)
	}
	if !is_num && coor[0] == 0 && coor[3] == 0 {
		return cost_str(strings.Join([]string{horizontal, vertical, "A"}, ""), depth, record)
	}
	if !is_num && coor[1] == 0 && coor[2] == 0 {
		return cost_str(strings.Join([]string{vertical, horizontal, "A"}, ""), depth, record)
	}
	return min(cost_str(strings.Join([]string{horizontal, vertical, "A"}, ""), depth, record),
		cost_str(strings.Join([]string{vertical, horizontal, "A"}, ""), depth, record))
}

func cost_str(s string, depth int, record map[string]int) int {
	if depth == 0 {
		return len(s)
	}
	key := s+strconv.Itoa(depth)
	cost, ok := record[key]
	if ok {
		return cost
	}
	total := 0
	x, y := toCoorDirectional('A')
	for _, c := range s {
		nx, ny := toCoorDirectional(c)
		coor := []int{nx, ny, x, y}
		x, y = nx, ny
		total += cost_coor(coor, depth-1, record, false)
	}
	record[key] = total
	return total
}

func part1(s string) int {
	record := make(map[string]int)
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		num, err := strconv.Atoi(line[:3])
		cost := 0
		if err != nil {
			log.Fatal(err)
		}
		x, y := toCoorNumeric('A')
		for _, r := range line {
			nx, ny := toCoorNumeric(r)
			cost += cost_coor([]int{nx, ny, x, y}, 2, record, true)
			x, y = nx, ny
		}
		sum += num * cost
	}
	return sum
}

func part2(s string) int {
	record := make(map[string]int)
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		num, err := strconv.Atoi(line[:3])
		cost := 0
		if err != nil {
			log.Fatal(err)
		}
		x, y := toCoorNumeric('A')
		for _, r := range line {
			nx, ny := toCoorNumeric(r)
			cost += cost_coor([]int{nx, ny, x, y}, 25, record, true)
			x, y = nx, ny
		}
		sum += num * cost
	}
	return sum
}
