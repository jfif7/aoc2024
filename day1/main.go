package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strings"
	"log"
)

//go:embed input.txt
var input string

func main() {	var part int
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

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func part1(s string) int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for _, line := range strings.Split(s, "\n") {
		var a, b int
		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			log.Fatalf("input parsing error")
		}
		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}
	slices.Sort(arr1)
	slices.Sort(arr2)
	sum := 0
	for i := range arr1 {
		sum += abs(arr1[i] - arr2[i])
	}

	return sum
}

func part2(s string) int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for _, line := range strings.Split(s, "\n") {
		var a, b int
		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			log.Fatalf("input parsing error")
		}
		map1[a] += 1
		map2[b] += 1
	}
	sum := 0
	for k, v := range map1 {
		sum += k * v * map2[k]
	}

	return sum
}