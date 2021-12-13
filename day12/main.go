package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/naisuuuu/aoc2021/input"
)

func main() {
	runInput := flag.Bool("i", false, "run on real input (instead of example)")
	flag.Parse()

	in, err := input.Read("example")
	if *runInput {
		in, err = input.Read("input")
	}
	if err != nil {
		log.Fatal(err)
	}

	solve(in)
}

func solve(in []string) {
	paths := make(map[string][]string, 0)
	for _, l := range in {
		ll := strings.Split(l, "-")
		paths[ll[0]] = append(paths[ll[0]], ll[1])
		paths[ll[1]] = append(paths[ll[1]], ll[0])
	}

	visited := make(map[string]bool)
	for k := range paths {
		visited[k] = false
	}

	seen := make(map[string]bool)
	var out []string

	var backtrack func(string, string, bool)
	backtrack = func(node, path string, visitUsed bool) {
		if node == "end" {
			if _, ok := seen[path]; ok {
				return
			}
			out = append(out, path)
			seen[path] = true
			return
		}
		for _, n := range paths[node] {
			if n == "start" || visited[n] {
				continue
			}
			if unicode.IsLower(rune(n[0])) {
				if !visitUsed {
					backtrack(n, fmt.Sprintf("%s,%s", path, n), true)
				}
				visited[n] = true
			}
			backtrack(n, fmt.Sprintf("%s,%s", path, n), visitUsed)
			visited[n] = false
		}
	}

	backtrack("start", "start", true)
	fmt.Println("Part 1:", len(out))

	out = []string{}
	seen = make(map[string]bool)

	backtrack("start", "start", false)
	fmt.Println("Part 2:", len(out))
}
