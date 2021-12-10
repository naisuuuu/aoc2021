package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

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

var points1 = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var points2 = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var openers = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func parseLine(line string) (illegal rune, remaining []rune) {
	for _, c := range line {
		switch c {
		case ')', ']', '}', '>':
			n := len(remaining) - 1
			if remaining[n] != openers[c] {
				return c, nil
			}
			remaining = remaining[:n]
		default:
			remaining = append(remaining, c)
		}
	}
	return
}

func solve(in []string) {
	var p1 int
	var p2 []int

	for _, line := range in {
		illegal, remaining := parseLine(line)
		switch len(remaining) {
		case 0:
			p1 += points1[illegal]
		default:
			var score int
			for i := len(remaining) - 1; i >= 0; i-- {
				score *= 5
				score += points2[remaining[i]]
			}
			p2 = append(p2, score)
		}
	}

	fmt.Println("Part 1:", p1)
	sort.Ints(p2)
	fmt.Println("Part 2:", p2[len(p2)/2])
}
