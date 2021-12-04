package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/naisuuuu/aoc2021/input"
)

func main() {
	in, err := input.Read("input")
	if err != nil {
		log.Fatal(err)
	}
	if err := solve(in); err != nil {
		log.Fatal(err)
	}
}

func solve(in []string) error {
	var ds []int
	for _, i := range in {
		d, err := strconv.Atoi(i)
		if err != nil {
			return err
		}
		ds = append(ds, d)
	}

	cnt := 0
	for i := 1; i < len(ds); i++ {
		if ds[i] > ds[i-1] {
			cnt++
		}
	}

	fmt.Println("Part one:", cnt)

	cnt = 0
	w := ds[0] + ds[1] + ds[2]
	for i := 3; i < len(ds); i++ {
		nw := w + ds[i] - ds[i-3]
		if nw > w {
			cnt++
		}
		w = nw
	}

	fmt.Println("Part two:", cnt)

	return nil
}
