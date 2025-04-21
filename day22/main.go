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

func next(secret uint32) uint32 {
	var mod_base uint32 = (1 << 24) - 1
	// var result uint32 = secret
	secret = (secret ^ (secret << 6)) & mod_base
	secret = (secret ^ (secret >> 5)) & mod_base
	secret = (secret ^ (secret << 11)) & mod_base
	return secret
}

func part1(s string) int {
	total := 0
	for _, line := range strings.Split(s, "\n") {
		secret, ok := strconv.Atoi(line)
		var usecret uint32 = uint32(secret)
		if ok != nil {
			log.Fatal("input parsing error")
		}
		for i:=0;i<2000;i++ {
			usecret = next(usecret)
		}
		total += int(usecret)
	}
	return total
}

func part2(s string) int {
	total_map := make(map[uint32]int)
	best := 0

	for _, line := range strings.Split(s, "\n") {
		secret, ok := strconv.Atoi(line)
		if ok != nil {
			log.Fatal("input parsing error")
		}
		var usecret uint32 = uint32(secret)
		visited := make(map[uint32]bool)
		var key uint32 = 0
		for i := 0; i < 2000; i++ {
			new_secret := next(usecret)
			diff := int(new_secret % 10) - int(usecret % 10)
			usecret = new_secret
			banana := int(new_secret % 10)
			key = ((key << 5) | uint32(diff+9))& ((1 << 20) - 1)
			_, ok := visited[key]
			if ok {
				continue
			} else if i >= 3 {
				visited[key] = true
				total_map[key] += banana
				if total_map[key] > best {
					best = total_map[key]
				}
			}
		}
	}
	return best
}
