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

func readLine(s string) []int64 {
	parts := strings.Split(s, ":")
	numbers := strings.Fields(parts[0] + " " + parts[1])
	var slice []int64
	for _, num := range numbers {
		n, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			log.Fatal("Error parsing input", err)
		}
		slice = append(slice, n)
	}
	return slice
}

func canSolve(target int64, operands []int64) bool {
	current_set := map[int64]bool{
		target: true,
	}
	next_set := map[int64]bool{}
	for i := len(operands) - 1; i >= 0; i -= 1 {
		for k := range current_set {
			// can subtract
			if k-operands[i] >= 0 {
				next_set[k-operands[i]] = true
			}
			// can divide
			if k%operands[i] == 0 {
				next_set[k/operands[i]] = true
			}
		}
		current_set = next_set
		next_set = map[int64]bool{}
	}
	_, ok := current_set[0]
	return ok
}

func pow10(exp int) int64 {
	var result int64 = 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}

func canSolveWith3(target int64, operands []int64) bool {
	current_set := map[int64]bool{
		target: true,
	}
	next_set := map[int64]bool{}
	for i := len(operands) - 1; i >= 0; i -= 1 {
		for k := range current_set {
			// can subtract
			if k-operands[i] >= 0 {
				next_set[k-operands[i]] = true
			}
			// can divide
			if k%operands[i] == 0 {
				next_set[k/operands[i]] = true
			}
			// can append
			digit := len(strconv.FormatInt(operands[i], 10))
			if (k-operands[i])%pow10(digit) == 0 {
				next_set[(k-operands[i])/pow10(digit)] = true
			}

		}
		current_set = next_set
		next_set = map[int64]bool{}
	}
	_, ok := current_set[0]
	return ok
}

func part1(s string) int64 {
	var sum int64 = 0
	for _, line := range strings.Split(s, "\n") {
		slice := readLine(line)
		target, operands := slice[0], slice[1:]
		if canSolve(target, operands) {
			sum += target
		}
	}
	return sum
}

func part2(s string) int64 {
	var sum int64 = 0
	for _, line := range strings.Split(s, "\n") {
		slice := readLine(line)
		target, operands := slice[0], slice[1:]
		if canSolveWith3(target, operands) {
			sum += target
		}
	}
	return sum
}
