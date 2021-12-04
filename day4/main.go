package main

import (
	"fmt"
	"log"

	"github.com/naisuuuu/aoc2021/conv"
	"github.com/naisuuuu/aoc2021/input"
)

func main() {
	in, err := input.Read("input")
	if err != nil {
		log.Fatal(err)
	}
	solve(in)
}

type board [][]int

func newBoard(s []string) board {
	var bb board
	for _, r := range s {
		bb = append(bb, conv.AtoiS(r, " "))
	}
	return bb
}

func (b board) sumRow(y int) int {
	var s int
	for _, n := range b[y] {
		s += n
	}
	return s
}

func (b board) sumCol(x int) int {
	var s int
	for _, r := range b {
		s += r[x]
	}
	return s
}

func (b board) update(n int) (solved bool) {
	for row := range b {
		for col := range b[0] {
			if b[row][col] != n {
				continue
			}

			b[row][col] = -1

			if b.sumRow(row) == -len(b) {
				// We don't want to exit early, because we need to update the entire board
				// to get the right answer.
				solved = true
			}
			if b.sumCol(col) == -len(b[0]) {
				solved = true
			}
		}
	}
	return
}

func (b board) sum() int {
	var s int
	for row := range b {
		for col := range b[0] {
			if b[row][col] > 0 {
				s += b[row][col]
			}
		}
	}
	return s
}

func solve(in []string) {
	numbers := conv.AtoiS(in[0], ",")

	var boards []board
	for i := 1; i < len(in); i++ {
		if len(in[i]) == 0 {
			boards = append(boards, newBoard(in[i+1:i+6]))
		}
	}

	var partOne bool
	for _, n := range numbers {
		unsolved := make([]board, 0, len(boards))
		for b := range boards {
			solved := boards[b].update(n)
			if !solved {
				unsolved = append(unsolved, boards[b])
			}
			if solved && !partOne {
				fmt.Println("Part 1:", boards[b].sum()*n)
				partOne = true
			}
		}
		if len(boards) == 1 {
			fmt.Println("Part 2:", boards[0].sum()*n)
			return
		}
		boards = unsolved
	}
}
