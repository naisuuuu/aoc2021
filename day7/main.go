package main

import (
	"flag"
	"fmt"
	"log"
	"math"

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

func solve(in []string) {
	hs := aocutil.AtoiS(in[0], ",")

	mx := aocutil.Max(hs...)
	mn := aocutil.Min(hs...)

	lowest := math.MaxInt
	for i := mn; i <= mx; i++ {
		var cnt int
		for _, h := range hs {
			cnt += int(math.Abs(float64(h - i)))
		}
		if cnt < lowest {
			lowest = cnt
		}
	}

	fmt.Println("Part 1:", lowest)

	lowest = math.MaxInt
	for i := mn; i <= mx; i++ {
		var cnt int
		for _, h := range hs {
			c := int(math.Abs(float64(h - i)))
			for j := 1; j <= c; j++ {
				cnt += j
			}
		}
		if cnt < lowest {
			lowest = cnt
		}
	}

	fmt.Println("Part 2:", lowest)
}
