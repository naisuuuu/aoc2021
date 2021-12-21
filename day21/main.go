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

func getPos(s string) int {
	return aocutil.MustAtoi(strings.Split(s, ": ")[1])
}

func game(p1, p2, limit int) (rolls, s1, s2 int) {
	p1--
	p2--
	for s1 < limit && s2 < limit {
		player := &p2
		score := &s2
		if rolls%2 == 0 {
			player = &p1
			score = &s1
		}
		var steps [3]int
		for i := range steps {
			steps[i] = rolls%100 + 1
			rolls++
		}
		npos := (*player + steps[0] + steps[1] + steps[2]) % 10
		*player = npos
		*score += *player + 1
		// p := 1
		// if rolls%2 == 0 {
		// 	p = 2
		// }
		// fmt.Printf(
		// 	"Player %d rolls %d+%d+%d and moves to space %d for a total score of %d.\n",
		// 	p,
		// 	steps[0], steps[1], steps[2],
		// 	*player,
		// 	*score,
		// )
	}
	return
}

// Combinations of 3-sided die (10 total, 7 unique) / no. permutations (27 total):
// 1, 1, 1 = 3 / 1
// 1, 1, 2 = 4 / 3
// 1, 1, 3 = 5 / 3
// 1, 2, 2 = 5 / 3
// 1, 2, 3 = 6 / 6
// 2, 2, 2 = 6 / 1
// 1, 3, 3 = 7 / 3
// 2, 2, 3 = 7 / 3
// 2, 3, 3 = 8 / 3
// 3, 3, 3 = 9 / 1

var outcomes = [7][2]int{
	{3, 1},
	{4, 3},
	{5, 6},
	{6, 7},
	{7, 6},
	{8, 3},
	{9, 1},
}

func quantumGame(p1, p2, limit int) (s1, s2 int) {
	var game func(int, int, int, int, int, bool)
	game = func(pl1, pl2, sc1, sc2, universes int, p1turn bool) {
		if sc1 >= limit {
			s1 += universes
			return
		}
		if sc2 >= limit {
			s2 += universes
			return
		}
		for _, o := range outcomes {
			switch p1turn {
			case true:
				np := (pl1 + o[0]) % 10
				game(np, pl2, sc1+np+1, sc2, universes*o[1], false)
			default:
				np := (pl2 + o[0]) % 10
				game(pl1, np, sc1, sc2+np+1, universes*o[1], true)
			}
		}
	}
	game(p1-1, p2-1, 0, 0, 1, true)
	return
}

func solve(in []string) {
	player1, player2 := getPos(in[0]), getPos(in[1])
	rolls, s1, s2 := game(player1, player2, 1000)
	fmt.Println("Part 1:", rolls*aocutil.Min(s1, s2))
	fmt.Println("Part 2:", aocutil.Max(quantumGame(player1, player2, 21)))
}
