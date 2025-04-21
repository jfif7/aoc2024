package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"log"
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

func has_edge(adj map[string][]string, u, v string) bool {
	for _, node := range adj[u] {
		if node == v {
			return true
		}
	}
	return false
}

func part1(s string) int {
	total := 0
	adj := make(map[string][]string)
	for _, line := range strings.Split(s, "\n") {
		var to, from string
		_, ok := fmt.Sscanf(line, "%2s-%2s", &to, &from)
		if ok != nil {
			log.Fatal("parsing error")
		}
		adj[to] = append(adj[to], from)
		adj[from] = append(adj[from], to)
	}
	for i:=0; i<26; i++ {
		key := "t" + string(rune(int('a')+i))
		for j, u := range adj[key] {
			if u[0] == 't' && u[1] < key[1] {
				continue
			}
			for k:=j+1;k < len(adj[key]);k++ {
				v := adj[key][k]
				if v[0] == 't' && v[1] < key[1] {
					continue
				}
				if has_edge(adj, u, v) {
					total++
				}
			}
		}
	}
	return total
}

func has_edge2(adj [][]int, u, v int) bool {
	for _, node := range adj[u] {
		if node == v {
			return true
		}
	}
	return false
}

func dfs(adj [][]int, idx int, cur []int, ans *[]int) {
	if len(cur) > len(*ans) {
		*ans = make([]int, len(cur))
		copy(*ans, cur)
	}
	if idx >= 26*26 {
		return
	}
	for i := idx; i< 26*26; i++ {
		ok := true
		for _, cur_node := range cur {
			if !has_edge2(adj, i, cur_node) {
				ok = false
				break
			}
		}
		if ok {
			dfs(adj, i + 1, append(cur, i), ans)
		}
	}
}

func part2(s string) string {
	adj := make([][]int, 26*26)
	for _, line := range strings.Split(s, "\n") {
		var to, from string
		_, ok := fmt.Sscanf(line, "%2s-%2s", &to, &from)
		if ok != nil {
			log.Fatal("parsing error")
		}
		to_idx := int(int(to[0] - 'a')*26 + int(to[1] - 'a'))
		from_idx := int(int(from[0] - 'a')*26 + int(from[1] - 'a'))
		adj[to_idx] = append(adj[to_idx], from_idx)
		adj[from_idx] = append(adj[from_idx], to_idx)
	}
	var ans []int
	dfs(adj, 0, []int{}, &ans)
	var buf bytes.Buffer
	for _, node := range ans {
		buf.WriteRune(rune(node/26 + 'a'))
		buf.WriteRune(rune(node%26 + 'a'))
		buf.WriteRune(',')
	}
	ans_string := buf.String()
	return ans_string[:len(ans_string)-1]
}
