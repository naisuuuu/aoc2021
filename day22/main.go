package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

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

type Cube struct {
	MinX, MaxX, MinY, MaxY, MinZ, MaxZ int
}

func NewCube(s string) Cube {
	var out Cube
	_, err := fmt.Sscanf(s,
		"x=%d..%d,y=%d..%d,z=%d..%d",
		&out.MinX, &out.MaxX, &out.MinY, &out.MaxY, &out.MinZ, &out.MaxZ)
	if err != nil {
		panic(err)
	}
	return out
}

func NewStep(s string) Step {
	ss := strings.Split(s, " ")
	out := Step{C: NewCube(ss[1])}
	if ss[0] == "on" {
		out.On = true
	}
	return out
}

type Step struct {
	On bool
	C  Cube
}

type State map[Cube]bool

func volume(c Cube) int {
	return (c.MaxX - c.MinX + 1) * (c.MaxY - c.MinY + 1) * (c.MaxZ - c.MinZ + 1)
}

func Count(s State) (count int) {
	for c := range s {
		count += volume(c)
	}
	return
}

func inRange(val, min, max int) bool {
	return val >= min && val <= max
}

// overlaps tests if a overlaps with b
func overlaps(a, b Cube) bool {
	x := inRange(a.MinX, b.MinX, b.MaxX) || inRange(b.MinX, a.MinX, a.MaxX)
	y := inRange(a.MinY, b.MinY, b.MaxY) || inRange(b.MinY, a.MinY, a.MaxY)
	z := inRange(a.MinZ, b.MinZ, b.MaxZ) || inRange(b.MinZ, a.MinZ, a.MaxZ)
	return x && y && z
}

// contains tests if a contains b
func contains(a, b Cube) bool {
	x := inRange(b.MinX, a.MinX, a.MaxX) && inRange(b.MaxX, a.MinX, a.MaxX)
	y := inRange(b.MinY, a.MinY, a.MaxY) && inRange(b.MaxY, a.MinY, a.MaxY)
	z := inRange(b.MinZ, a.MinZ, a.MaxZ) && inRange(b.MaxZ, a.MinZ, a.MaxZ)
	return x && y && z
}

func Apply(state map[Cube]bool, step Step) {
	var overlap []Cube
	for c := range state {
		if overlaps(c, step.C) {
			overlap = append(overlap, c)
		}
	}
	for _, c := range overlap {
		delete(state, c)
		for _, r := range split(c, step.C) {
			state[r] = true
		}
	}
	if step.On {
		state[step.C] = true
	}
}

func ccopy(a Cube) *Cube {
	return &a
}

func splitX(a, b Cube) (left, center, right *Cube) {
	center = ccopy(a)
	if inRange(b.MinX, a.MinX, a.MaxX) {
		left = ccopy(a)
		left.MaxX = b.MinX - 1
		center.MinX = b.MinX
	}
	if inRange(b.MaxX, a.MinX, a.MaxX) {
		right = ccopy(a)
		right.MinX = b.MaxX + 1
		center.MaxX = b.MaxX
	}
	return
}

func splitY(a, b Cube) (left, center, right *Cube) {
	center = ccopy(a)
	if inRange(b.MinY, a.MinY, a.MaxY) {
		left = ccopy(a)
		left.MaxY = b.MinY - 1
		center.MinY = b.MinY
	}
	if inRange(b.MaxY, a.MinY, a.MaxY) {
		right = ccopy(a)
		right.MinY = b.MaxY + 1
		center.MaxY = b.MaxY
	}
	return
}

func splitZ(a, b Cube) (left, center, right *Cube) {
	center = ccopy(a)
	if inRange(b.MinZ, a.MinZ, a.MaxZ) {
		left = ccopy(a)
		left.MaxZ = b.MinZ - 1
		center.MinZ = b.MinZ
	}
	if inRange(b.MaxZ, a.MinZ, a.MaxZ) {
		right = ccopy(a)
		right.MinZ = b.MaxZ + 1
		center.MaxZ = b.MaxZ
	}
	return
}

func isFlat(c Cube) bool {
	return c.MinX > c.MaxX || c.MinY > c.MaxY || c.MinZ > c.MaxZ
}

func split(a, b Cube) (outside []Cube) {
	l, c, r := splitX(a, b)
	if l != nil && !isFlat(*l) {
		outside = append(outside, *l)
	}
	if r != nil && !isFlat(*r) {
		outside = append(outside, *r)
	}
	l, c, r = splitY(*c, b)
	if l != nil && !isFlat(*l) {
		outside = append(outside, *l)
	}
	if r != nil && !isFlat(*r) {
		outside = append(outside, *r)
	}
	l, _, r = splitZ(*c, b)
	if l != nil && !isFlat(*l) {
		outside = append(outside, *l)
	}
	if r != nil && !isFlat(*r) {
		outside = append(outside, *r)
	}
	return
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func tooLarge(s Step, limit int) bool {
	for _, n := range []int{s.C.MinX, s.C.MaxX, s.C.MinY, s.C.MaxY, s.C.MinZ, s.C.MaxZ} {
		if abs(n) > limit {
			return true
		}
	}
	return false
}

func solve(in []string) {
	var steps []Step
	for _, l := range in {
		steps = append(steps, NewStep(l))
	}

	state := make(State)
	for _, step := range steps {
		if tooLarge(step, 50) {
			continue
		}
		Apply(state, step)
	}

	fmt.Println("Part 1:", Count(state))

	state = make(State)
	for _, step := range steps {
		Apply(state, step)
	}
	fmt.Println("Part 2:", Count(state))
}
