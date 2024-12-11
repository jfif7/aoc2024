package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	next_id := 0
	next_pos := 0
	next_is_block := true
	back_left := int(s[len(s)-1] - '0')
	var sum int64 = 0
	for i, j := 0, len(s)-1; i <= j; i++ {
		size := int(s[i] - '0')
		if next_is_block {
			if i == j {
				size = back_left
			}
			sum += int64(i/2) * int64((next_pos*2+size-1)*(size)/2)
			// fmt.Printf("fill pos [%v, %v] with id %v\n", next_pos, next_pos+size-1, i/2)
			next_pos += size
			next_id += 1
			next_is_block = false
		} else {
			for size > 0 && j > i {
				if back_left > 0 {
					sum += int64(j/2) * int64(next_pos)
					// fmt.Printf("fill pos %v with id %v\n", next_pos, j/2)
					next_pos++
					size--
					back_left--
				} else {
					j -= 2
					back_left = int(s[j] - '0')
				}
			}
			next_is_block = true
		}
	}
	return sum
}

type File struct {
	id   int
	pos  int
	size int
}

func part2(s string) int64 {
	var sum int64 = 0
	blanks := []File{}
	files := []File{}
	next_pos := 0
	for i, r := range s {
		size := int(r - '0')
		if i%2 == 0 { // file
			files = append(files, File{i / 2, next_pos, size})
		} else { // blanks
			blanks = append(blanks, File{0, next_pos, size})
		}
		next_pos += size
	}
	for i := len(files) - 1; i > 0; i-- {
		f := files[i]
		moved := false
		for j, blank := range blanks {
			if blank.pos >= f.pos {
				
				sum += int64(f.id) * int64((blank.pos*2+f.size-1)*(f.size)/2)
				blank.size -= f.size
				blank.pos += f.size
				blanks[j] = blank
				moved = true
				break
			}
		}
		if !moved {
			sum += int64(f.id) * int64((f.pos*2+f.size-1)*(f.size)/2)
		}
	}
	return sum
}
