package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	zeroes := make([]int, len(in[0]))
	for _, i := range in {
		for j, c := range i {
			if c == '0' {
				zeroes[j]++
			} else {
				ones[j]++
			}
		}
	}
	var gamma, eps []string
	for i := range ones {
		if ones[i] > zeroes[i] {
			gamma = append(gamma, "1")
			eps = append(eps, "0")
		} else {
			gamma = append(gamma, "0")
			eps = append(eps, "1")
		}
	}
	gi, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	ei, _ := strconv.ParseInt(strings.Join(eps, ""), 2, 64)
	fmt.Println("Part 1:", gi*ei)
}

func p2(in []string) {
	options := in
	matches := make([]string, 0)
	for i := 0; i < len(in[0]); i++ {
		if len(options) <= 1 {
			break
		}
		var ones, zeroes int
		for _, o := range options {
			if o[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		gv := '1'
		if zeroes > ones {
			gv = '0'
		}
		for _, o := range options {
			if rune(o[i]) == gv {
				matches = append(matches, o)
			}
		}
		options = matches
		matches = make([]string, 0)
	}

	generator := options[0]

	options = in
	matches = make([]string, 0)
	for i := 0; i < len(in[0]); i++ {
		if len(options) <= 1 {
			break
		}
		var ones, zeroes int
		for _, o := range options {
			if o[i] == '0' {
				zeroes++
			} else {
				ones++
			}
		}
		gv := '0'
		if zeroes > ones {
			gv = '1'
		}
		for _, o := range options {
			if rune(o[i]) == gv {
				matches = append(matches, o)
			}
		}
		options = matches
		matches = make([]string, 0)
	}

	scrubber := options[0]

	gi, _ := strconv.ParseInt(generator, 2, 64)
	si, _ := strconv.ParseInt(scrubber, 2, 64)
	fmt.Println("Part 2:", gi*si)
}
