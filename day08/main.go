package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/naisuuuu/aoc2021/aocutil"
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

type entry struct {
	inputs  []string
	outputs []string
}

func makeEntry(s string) entry {
	ss := strings.Split(s, " | ")
	return entry{
		inputs:  strings.Fields(ss[0]),
		outputs: strings.Fields(ss[1]),
	}
}

func countMissingRunes(in, from string) (cnt int, last rune) {
	for _, c := range from {
		if !strings.ContainsRune(in, c) {
			cnt++
			last = c
		}
	}
	return
}

func solveEntry(e entry) int {
	nums := make(map[int]string)
	lens := make(map[int][]string)
	for _, o := range e.inputs {
		switch len(o) {
		case 2:
			nums[1] = o
		case 3:
			nums[7] = o
		case 4:
			nums[4] = o
		case 7:
			nums[8] = o
		default:
			lens[len(o)] = append(lens[len(o)], o)
		}
	}

	for _, d := range lens[6] {
		// Numbers of len 6 are 0, 6 and 9. When comparing to 4, only 9 is missing 2, instead of 3 edges.
		if missing, _ := countMissingRunes(nums[4], d); missing == 2 {
			nums[9] = d
			break
		}
	}

	for _, d := range lens[5] {
		missing, missed := countMissingRunes(d, nums[9])
		// Numbers of len 5: 2, 3, 5.
		switch missing {
		// When comparing to 9, 2 has 2 missing edges.
		case 2:
			nums[2] = d
		case 1:
			i := 3
			// Both 5 and 3 have 1 missing edge comparing to 9, but only 5's missing edge is present in 1.
			if strings.ContainsRune(nums[1], missed) {
				i = 5
			}
			nums[i] = d
		}
	}

	for _, d := range lens[6] {
		// We already know what 9 is.
		if d == nums[9] {
			continue
		}
		missing, _ := countMissingRunes(nums[5], d)
		switch missing {
		case 2:
			nums[0] = d
		case 1:
			nums[6] = d
		}
	}

	decode := make(map[string]string)
	for k, v := range nums {
		decode[aocutil.SortString(v)] = strconv.Itoa(k)
	}

	out := ""
	for _, n := range e.outputs {
		out += decode[aocutil.SortString(n)]
	}

	o, _ := strconv.Atoi(out)

	return o
}

func solve(in []string) {
	var es []entry
	for _, i := range in {
		es = append(es, makeEntry(i))
	}

	var cnt int
	for _, e := range es {
		for _, o := range e.outputs {
			switch len(o) {
			case 2, 4, 3, 7:
				cnt++
			}
		}
	}

	fmt.Println("Part 1:", cnt)

	var out []int
	for _, e := range es {
		out = append(out, solveEntry(e))
	}

	cnt = 0
	for _, o := range out {
		cnt += o
	}

	fmt.Println("Part 2:", cnt)
}
