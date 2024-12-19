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

func canDo(towels map[string]bool, query string) bool {
	if len(query) == 0 {
		return true
	}
	t, ok := towels[query]
	if ok {
		return t
	}
	for i := 1; i < len(query); i++ {
		if canDo(towels, query[:i]) && canDo(towels, query[i:]) {
			towels[query] = true
			return true
		}
	}
	towels[query] = false
	return false
}

func part1(s string) int {
	regions := strings.Split(s, "\n\n")
	towels := make(map[string]bool)
	for _, towel := range strings.Split(regions[0], ", ") {
		towels[towel] = true
	}
	sum := 0
	for _, query := range strings.Split(regions[1], "\n") {
		if canDo(towels, query) {
			sum++
		}
	}
	return sum
}

func manyCanDo(towels map[string]bool, more_towels map[string]int, query string) int {
	if len(query) == 0 {
		return 1
	}
	t, ok := more_towels[query]
	if ok {
		return t
	}
	sum := 0
	for i := 1; i <= len(query); i++ {
		if towels[query[:i]] {
			many := manyCanDo(towels, more_towels, query[i:])
			sum += many
		}
	}
	more_towels[query] = sum
	return sum
}

func part2(s string) int {
	regions := strings.Split(s, "\n\n")
	towels := make(map[string]bool)
	more_towels := make(map[string]int)
	for _, towel := range strings.Split(regions[0], ", ") {
		towels[towel] = true
	}
	sum := 0
	for _, query := range strings.Split(regions[1], "\n") {
		if many := manyCanDo(towels, more_towels, query); many != 0 {
			sum += many
		}
	}
	return sum
}
