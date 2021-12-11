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

func sum(grid [][]int) int {
	var s int
	for row := range grid {
		for col := range grid[0] {
			s += grid[row][col]
		}
	}
	return s
}

func solve(in []string) {
	var grid [][]int
	for _, l := range in {
		grid = append(grid, aocutil.AtoiS(l, ""))
	}

	var cnt, step int
	for {
		step++

		for row := range grid {
			for col := range grid[0] {
				grid[row][col] += 1
			}
		}

		for row := range grid {
			for col := range grid[0] {
				if grid[row][col] <= 9 {
					continue
				}

				cnt += 1
				grid[row][col] = 0

				var s [][2]int
				s = append(s, [2]int{row, col})
				for {
					if len(s) == 0 {
						break
					}

					r, c := s[len(s)-1][0], s[len(s)-1][1]
					s = s[:len(s)-1]
					for _, n := range aocutil.NeighborsDiag(grid, r, c) {
						nr, nc := n[0], n[1]
						if grid[nr][nc] == 0 {
							continue
						}

						grid[nr][nc] += 1
						if grid[nr][nc] > 9 {
							cnt += 1
							grid[nr][nc] = 0
							s = append(s, n)
						}
					}
				}
			}
		}

		if step == 100 {
			fmt.Println("Part 1:", cnt)
		}

		if sum(grid) == 0 {
			break
		}
	}

	fmt.Println("Part 2:", step)
}
