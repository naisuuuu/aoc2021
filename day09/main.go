package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

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

func neighbors(grid [][]int, row, col int) [][2]int {
	var out [][2]int
	for _, m := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nrow, ncol := row+m[0], col+m[1]
		if nrow < len(grid) && nrow >= 0 && ncol < len(grid[0]) && ncol >= 0 {
			out = append(out, [2]int{nrow, ncol})
		}
	}
	return out
}

func lowPoint(grid [][]int, row, col int) bool {
	for _, n := range neighbors(grid, row, col) {
		if grid[n[0]][n[1]] <= grid[row][col] {
			return false
		}
	}
	return true
}

func solve(in []string) {
	var grid [][]int
	for _, i := range in {
		grid = append(grid, aocutil.AtoiS(i, ""))
	}

	var cnt int
	for row := range grid {
		for col := range grid[0] {
			if lowPoint(grid, row, col) {
				cnt += (1 + grid[row][col])
			}
		}
	}

	fmt.Println("Part 1:", cnt)

	skip := func(i int) bool {
		if i == 9 || i == -1 {
			return true
		}
		return false
	}
	var basins []int
	for row := range grid {
		for col := range grid[0] {
			if skip(grid[row][col]) {
				continue
			}

			l := 1
			grid[row][col] = -1

			q := make([][2]int, 0)
			q = append(q, [2]int{row, col})
			for {
				if len(q) == 0 {
					break
				}
				y, x := q[0][0], q[0][1]
				q = q[1:]
				for _, n := range neighbors(grid, y, x) {
					ny, nx := n[0], n[1]
					if skip(grid[ny][nx]) {
						continue
					}
					grid[ny][nx] = -1
					l++
					q = append(q, [2]int{ny, nx})
				}
			}
			basins = append(basins, l)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	fmt.Println("Part 2:", basins[0]*basins[1]*basins[2])
}
