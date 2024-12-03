package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
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
func part1(s string) int {
	r, err := regexp.Compile(`mul\((?P<multiplicant>\d{1,3}),(?P<multipler>\d{1,3})\)`)
	if err != nil {
		log.Fatalf("compile regex error")
	}
	matches := r.FindAllStringSubmatch(s, -1)
	sum := 0
	for _, m := range matches {
		multiplicant, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatalf("parse regex error")
		}
		multiplier, err := strconv.Atoi(m[2])
		if err != nil {
			log.Fatalf("parse regex error")
		}
		sum += multiplicant * multiplier
	}
	return sum
}

func part2(s string) int {
	r, err := regexp.Compile(`(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)`)
	if err != nil {
		log.Fatalf("compile regex error")
	}
	matches := r.FindAllStringSubmatch(s, -1)
	sum := 0
	able := true
	for _, m := range matches {
		if m[1] == "mul" && able {
			multiplicant, err := strconv.Atoi(m[2])
			if err != nil {
				log.Fatalf("parse regex error")
			}
			multiplier, err := strconv.Atoi(m[3])
			if err != nil {
				log.Fatalf("parse regex error")
			}
			sum += multiplicant * multiplier
		} else if m[4] == "do" {
			able = true
		} else if m[5] == "don't" {
			able = false
		}
	}
	return sum
}
