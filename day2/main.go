package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	type dir struct {
		dir string
		amt int
	}
	var ds []dir
	for _, i := range in {
		ss := strings.Split(i, " ")
		si, err := strconv.Atoi(ss[1])
		if err != nil {
			return err
		}
		ds = append(ds, dir{ss[0], si})
	}

	var (
		up, down, forward int
	)
	for _, d := range ds {
		switch d.dir {
		case "forward":
			forward += d.amt
		case "down":
			down += d.amt
		case "up":
			up += d.amt
		}
	}
	fmt.Println("Part one:", forward*(down-up))

	var (
		aim, hor, depth int
	)
	for _, d := range ds {
		switch d.dir {
		case "forward":
			hor += d.amt
			depth += (aim * d.amt)
		case "down":
			aim += d.amt
		case "up":
			aim -= d.amt
		}
	}
	fmt.Println("Part two:", hor*depth)

	return nil
}
