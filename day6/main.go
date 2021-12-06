package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/naisuuuu/aoc2021/conv"
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
	var ff, nf [9]int

	for _, n := range conv.AtoiS(fish, ",") {
		ff[n] += 1
	}

	for d := 0; d < days; d++ {
		for i := 8; i >= 0; i-- {
			switch i {
			case 0:
				nf[8] = ff[i]
				nf[6] += ff[i]
			default:
				nf[i-1] = ff[i]
			}
		}
		ff, nf = nf, ff
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
