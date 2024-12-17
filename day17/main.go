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

func combo(operand int64, reg []int64) int64 {
	switch {
	case 0 <= operand && operand <= 3:
		return operand
	case 4 <= operand && operand <= 6:
		return reg[operand-4]
	default:
		log.Fatal("combo error")
	}
	return -1
}

func part1(s string) string {
	lines := strings.Split(s, "\n")
	reg := []int64{0, 0, 0}
	var prog_string string
	fmt.Sscanf(lines[0], "Register A: %d", &reg[0])
	fmt.Sscanf(lines[1], "Register B: %d", &reg[1])
	fmt.Sscanf(lines[2], "Register C: %d", &reg[2])
	fmt.Sscanf(lines[4], "Program: %s", &prog_string)
	prog := make([]int64, 0)
	for _, op_str := range strings.Split(prog_string, ",") {
		op, err := strconv.Atoi(op_str)
		if err != nil {
			log.Fatal("parse prog error", err)
		}
		prog = append(prog, int64(op))
	}
	fmt.Println(reg)
	fmt.Println(prog)
	ip := 0
	output := make([]string, 0)
	for {
		if ip >= len(prog) {
			break
		}
		opcode := prog[ip]
		operand := prog[ip+1]
		switch opcode {
		case 0:
			reg[0] = reg[0] >> combo(operand, reg)
		case 1:
			reg[1] ^= operand
		case 2:
			reg[1] = combo(operand, reg) & 7
		case 3:
			if reg[0] != 0 {
				ip = int(operand)
				continue
			}
		case 4:
			reg[1] ^= reg[2]
		case 5:
			output = append(output, strconv.FormatInt(combo(operand, reg)&7, 10))
		case 6:
			reg[1] = reg[0] >> combo(operand, reg)
		case 7:
			reg[2] = reg[0] >> combo(operand, reg)
		}
		ip += 2
	}
	return strings.Join(output, ",")
}

func first_output(reg []int64, prog []int) int64 {
	ip := 0
	for {
		if ip >= len(prog) {
			break
		}
		opcode := prog[ip]
		operand := int64(prog[ip+1])
		switch opcode {
		case 0:
			reg[0] = reg[0] >> combo(operand, reg)
		case 1:
			reg[1] ^= operand
		case 2:
			reg[1] = combo(operand, reg) & 7
		case 3:
			if reg[0] != 0 {
				ip = int(operand)
				continue
			}
		case 4:
			reg[1] ^= reg[2]
		case 5:
			return combo(operand, reg) & 7
		case 6:
			reg[1] = reg[0] >> combo(operand, reg)
		case 7:
			reg[2] = reg[0] >> combo(operand, reg)
		}
		ip += 2
	}
	return -1
}

func dfs(prog []int, index int, a int64) (int64, bool) {
	if index < 0 {
		return a, true
	}
	for i := 0; i <= 7; i++ {
		new_a := (a << 3) + int64(i)
		if first_output([]int64{new_a, 0, 0}, prog) == int64(prog[index]) {
			final_a, ok := dfs(prog, index-1, new_a)
			if ok {
				return final_a, ok
			}
		}
	}
	return -1, false
}

func part2(s string) int64 {
	lines := strings.Split(s, "\n")
	reg := []int64{0, 0, 0}
	var prog_string string
	fmt.Sscanf(lines[0], "Register A: %d", &reg[0])
	fmt.Sscanf(lines[1], "Register B: %d", &reg[1])
	fmt.Sscanf(lines[2], "Register C: %d", &reg[2])
	fmt.Sscanf(lines[4], "Program: %s", &prog_string)
	prog := make([]int, 0)
	for _, op_str := range strings.Split(prog_string, ",") {
		op, err := strconv.Atoi(op_str)
		if err != nil {
			log.Fatal("parse prog error", err)
		}
		prog = append(prog, op)
	}
	fmt.Println(reg)
	fmt.Println(prog)
	a, _ := dfs(prog, len(prog)-1, 0)
	return a
}
