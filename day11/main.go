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
	stones := strings.Fields(s)
	for i := 0; i < 25; i++ {
		stones_new := make([]string, 0)
		if len(stones) < 8 {
			fmt.Println(stones)
		} else {
			fmt.Println(stones[:8])
		}
		for j := 0; j < len(stones); j++ {
			stone := stones[j]
			switch true {
			case stone == "0":
				stones_new = append(stones_new, "1")
			case len(stone)%2 == 0:
				left, right := stone[:len(stone)/2], stone[len(stone)/2:]
				left = string_to_string(left, 1)
				right = string_to_string(right, 1)
				stones_new = append(stones_new, left)
				stones_new = append(stones_new, right)
			default:
				stones_new = append(stones_new, string_to_string(stone, 2024))
			}
		}
		stones = stones_new
	}
	return int64(len(stones))
}

func string_to_string(s string, m int) string {
	left_n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal()
	}
	return strconv.Itoa(left_n * m)
}

func part2(s string) int64 {
	stones := make(map[string]int)
	for _, stone := range strings.Fields(s) {
		stones[stone]++
	}
	for i := 0; i < 75; i++ {
		stones_new := make(map[string]int)
		for stone, val := range stones {
			switch true {
			case stone == "0":
				stones_new["1"] += val
			case len(stone)%2 == 0:
				left, right := stone[:len(stone)/2], stone[len(stone)/2:]
				left = string_to_string(left, 1)
				right = string_to_string(right, 1)
				stones_new[left] += val
				stones_new[right] += val
			default:
				stones_new[string_to_string(stone, 2024)] += val
			}
		}
		stones = stones_new
	}
	var sum int64
	for _, v := range stones {
		sum += int64(v)
	}
	return sum
}
