package main

import (
	"flag"
	"fmt"
	"log"

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

func fishSim(fish string, days int) (count int) {
	var ff [9]int
	for _, n := range aocutil.AtoiS(fish, ",") {
		ff[n] += 1
	}

	for d := 0; d < days; d++ {
		ff[(d+7)%9] += ff[d%9]
	}

	for _, n := range ff {
		count += n
	}
	return
}

func solve(in []string) {
	fmt.Println("Part 1:", fishSim(in[0], 80))
	fmt.Println("Part 2:", fishSim(in[0], 256))
}
