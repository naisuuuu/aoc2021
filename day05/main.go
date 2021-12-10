package main

import (
	"flag"
	"fmt"
	"log"
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

type coord [4]int

func newCoord(s string) coord {
	cs := strings.Split(s, " -> ")
	a := aocutil.AtoiS(cs[0], ",")
	b := aocutil.AtoiS(cs[1], ",")
	return coord{a[0], a[1], b[0], b[1]}
}

type grid [][]int

func newGrid(side int) grid {
	grid := make([][]int, side+1)
	for i := range grid {
		grid[i] = make([]int, side+1)
	}
	return grid
}

func (g grid) cnt() int {
	var cnt int
	for row := range g {
		for col := range g[0] {
			if g[row][col] > 1 {
				cnt += 1
			}
		}
	}
	return cnt
}

func solve(in []string) {
	var side int
	var coords []coord
	for _, l := range in {
		c := newCoord(l)
		coords = append(coords, c)
		for _, co := range c {
			if co > side {
				side = co
			}
		}
	}

	g := newGrid(side)

	for _, c := range coords {
		if c[0] != c[2] && c[1] != c[3] {
			// Ignore diagonal lines for p1.
			continue
		}

		minx, maxx := c[0], c[2]
		if c[0] > c[2] {
			minx, maxx = c[2], c[0]
		}

		miny, maxy := c[1], c[3]
		if c[1] > c[3] {
			miny, maxy = c[3], c[1]
		}

		for x := minx; x <= maxx; x++ {
			for y := miny; y <= maxy; y++ {
				g[y][x]++
			}
		}
	}

	fmt.Println("Part 1:", g.cnt())

	for _, c := range coords {
		if c[0] == c[2] || c[1] == c[3] {
			// Straight lines have already been applied.
			continue
		}

		xinc := 1
		if c[0] > c[2] {
			xinc = -1
		}

		yinc := 1
		if c[1] > c[3] {
			yinc = -1
		}

		x, y := c[0], c[1]
		for {
			g[y][x]++
			if x == c[2] && y == c[3] {
				break
			}
			x += xinc
			y += yinc
		}
	}

	fmt.Println("Part 2:", g.cnt())
}
