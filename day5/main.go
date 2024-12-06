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

func check_update(nums []string, rules map[string][]string) int {
	history := make(map[string]bool)
	for _, num := range nums {
		history[num] = true
		for _, back := range rules[num] {
			if _, ok := history[back]; ok {
				return 0
			}
		}
	}
	ret, err := strconv.Atoi(nums[len(nums)/2])
	if err != nil {
		log.Fatalf("Error parsing num")
	}
	return ret
}

func check_update_fix(nums []string, rules map[string][]string) int {
	swapped := false
	history := make(map[string]int)

	for i := 0; i < len(nums); i += 1 {
		num := nums[i]
		history[num] = i
		for _, back := range rules[num] {
			if index, ok := history[back]; ok && index < i {
				// Rewrites history and swap those numbers
				nums[index] = nums[i]
				nums[i] = back
				history[back] = i
				history[nums[index]] = index

				swapped = true
				i = index - 1 // Redo the outer loop from i = index
				break
			}
		}
	}

	if swapped {
		ret, err := strconv.Atoi(nums[len(nums)/2])
		if err != nil {
			log.Fatalf("Error parsing num")
		}
		return ret
	}
	return 0
}

func part1(s string) int {
	rules_map := make(map[string][]string)
	reading_rules := true
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 {
			reading_rules = false
			continue
		}
		if reading_rules {
			rule := strings.Split(line, "|")
			rules_map[rule[0]] = append(rules_map[rule[0]], rule[1])
		} else {
			nums := strings.Split(line, ",")
			sum += check_update(nums, rules_map)
		}
	}
	return sum
}

func part2(s string) int {
	rules_map := make(map[string][]string)
	reading_rules := true
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 {
			reading_rules = false
			continue
		}
		if reading_rules {
			rule := strings.Split(line, "|")
			rules_map[rule[0]] = append(rules_map[rule[0]], rule[1])
		} else {
			nums := strings.Split(line, ",")
			sum += check_update_fix(nums, rules_map)
		}
	}
	return sum
}
