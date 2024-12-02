package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func parseInts(s string) []int {
	// Split the string by spaces
	parts := strings.Fields(s)
	// Create a slice to hold the integers
	var result []int
	// Iterate over the parts and convert them to integers
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			// Handle the error if the conversion fails
			fmt.Println("Error converting string to int:", err)
			continue
		}
		result = append(result, num)
	}
	return result
}

func is_safe(arr []int, inc bool, prev int) int {
	// fmt.Println("is_safe", prev, "+", arr)
	ok := 0
	for i, n := range arr {
		if inc && (n-prev < 1 || n-prev > 3) {
			ok = i + 1
			break
		}
		if !inc && (prev-n < 1 || prev-n > 3) {
			ok = i + 1
			break
		}
		prev = n
	}
	return ok
}

func part1(s string) int {
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		nums := parseInts(line)
		if is_safe(nums[1:], false, nums[0]) == 0 || is_safe(nums[1:], true, nums[0]) == 0 {
			sum += 1
		}
	}
	return sum
}

func part2(s string) int {
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		nums := parseInts(line)
		// fmt.Println(nums)
		ok := is_safe(nums[1:], false, nums[0])
		if ok == 0 {
			sum += 1
			continue
		}
		// fmt.Println(ok)
		if is_safe(nums[ok+1:], false, nums[ok-1]) == 0 {
			sum += 1
			continue
		}
		if ok >= 2 && is_safe(nums[ok:], false, nums[ok-2]) == 0 ||
			ok < 2 && is_safe(nums[ok+1:], false, nums[ok]) == 0 {
			sum += 1
			continue
		}
		ok = is_safe(nums[1:], true, nums[0])
		if ok == 0 {
			sum += 1
			continue
		}
		// fmt.Println(ok)
		if is_safe(nums[ok+1:], true, nums[ok-1]) == 0 {
			sum += 1
			continue
		}
		if ok >= 2 && is_safe(nums[ok:], true, nums[ok-2]) == 0 ||
			ok < 2 && is_safe(nums[ok+1:], true, nums[ok]) == 0 {
			sum += 1
			continue
		}
	}
	return sum
}
