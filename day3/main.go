package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/naisuuuu/aoc2021/input"
)

func main() {
	in, err := input.Read()
	if err != nil {
		log.Fatal(err)
	}
	p1(in)
	p2(in)
}

func p1(in []string) {
	ones := make([]int, len(in[0]))
	for _, i := range in {
		for j, c := range i {
			if c == '1' {
				ones[j]++
			}
		}
	}
	var gamma uint
	for i, cnt := range ones {
		if cnt > len(in)/2 {
			gamma += 1 << (len(ones) - 1 - i)
		}
	}
	fmt.Println("Part 1:", gamma*(^gamma&(1<<len(ones)-1)))
}

func sieve(in []string, a, b byte) int64 {
	options, matches := in, []string{}
	for i := 0; i < len(in[0]); i++ {
		if len(options) <= 1 {
			break
		}
		var zeroes int
		for _, o := range options {
			if o[i] == '0' {
				zeroes++
			}
		}
		f := a
		if zeroes > len(options)/2 {
			f = b
		}
		for _, o := range options {
			if o[i] == f {
				matches = append(matches, o)
			}
		}
		options, matches = matches, []string{}
	}
	out, _ := strconv.ParseInt(options[0], 2, 64)
	return out
}

func p2(in []string) {
	g := sieve(in, '1', '0')
	s := sieve(in, '0', '1')
	fmt.Println("Part 2:", g*s)
}
