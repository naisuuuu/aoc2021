package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"

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

func split(s string) (left string, right string) {
	splitIndex := lenParens(s[1:]) + 1
	return s[1:splitIndex], s[splitIndex+1 : len(s)-1]
}

func lenParens(s string) int {
	var stack int
	for i, c := range s {
		switch c {
		case '[':
			stack++
		case ']':
			stack--
		case ',':
			if stack == 0 {
				return i
			}
		}
	}
	return 0
}

// ParseTree parses a string representation of a snailfish number into a tree structure.
func ParseTree(s string) *Node {
	n := &Node{}
	switch s[0] {
	case '[':
		l, r := split(s)
		n.Left, n.Right = ParseTree(l), ParseTree(r)
	default:
		n.Value = aocutil.MustAtoi(s)
	}
	return n
}

// Add performs addition on one or more snailfish numbers.
func Add(left *Node, right ...*Node) *Node {
	for _, r := range right {
		left = &Node{Left: left, Right: r}
		left.reduce()
	}
	return left
}

// Magnitude returns the magnitude of a snailfish number.
func Magnitude(n *Node) int {
	if n == nil {
		return 0
	}
	if n.isRegularNum() {
		return n.Value
	}
	return 3*Magnitude(n.Left) + 2*Magnitude(n.Right)
}

// Node represents a tree node in a snailfish number.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) updateLeft(num int) {
	if num == 0 {
		return
	}
	if n.Left.isRegularNum() {
		n.Left.Value += num
	} else {
		nn := n.Left.Right
		for !nn.isRegularNum() {
			nn = nn.Right
		}
		nn.Value += num
	}
}

func (n *Node) updateRight(num int) {
	if num == 0 {
		return
	}
	if n.Right.isRegularNum() {
		n.Right.Value += num
	} else {
		nn := n.Right.Left
		for !nn.isRegularNum() {
			nn = nn.Left
		}
		nn.Value += num
	}
}

func (n *Node) isRegularNum() bool {
	return n.Left == nil && n.Right == nil
}

func (n *Node) explode(depth int) (exploded bool, left, right int) {
	if n.isRegularNum() {
		return
	}
	switch depth {
	case 0:
		exploded = true
		left = n.Left.Value
		right = n.Right.Value
		n.Left = nil
		n.Right = nil
	default:
		exploded, left, right = n.Left.explode(depth - 1)
		if exploded {
			n.updateRight(right)
			right = 0
			return
		}
		exploded, left, right = n.Right.explode(depth - 1)
		if exploded {
			n.updateLeft(left)
			left = 0
		}
	}
	return
}

func (n *Node) split() bool {
	if n.Left != nil {
		if n.Left.split() {
			return true
		}
		return n.Right.split()
	}
	if n.Value > 9 {
		n.Left = &Node{Value: int(math.Floor(float64(n.Value) / 2))}
		n.Right = &Node{Value: int(math.Ceil(float64(n.Value) / 2))}
		n.Value = 0
		return true
	}
	return false
}

func (n *Node) reduce() {
	exploded, _, _ := n.explode(4)
	for exploded {
		exploded, _, _ = n.explode(4)
	}
	if n.split() {
		n.reduce()
	}
}

func (n *Node) String() string {
	// if Left != nil, Right != nil too.
	if n.Left != nil {
		return fmt.Sprintf("[%s,%s]", n.Left, n.Right)
	}
	return strconv.Itoa(n.Value)
}

func solve(in []string) {
	var ns []*Node
	for _, l := range in {
		ns = append(ns, ParseTree(l))
	}
	got := Add(ns[0], ns[1:]...)
	fmt.Println("Part 1:", Magnitude(got))

	var largest int
	for i, n := range in {
		for j, m := range in {
			if i == j {
				continue
			}
			sum := Add(ParseTree(n), ParseTree(m))
			mag := Magnitude(sum)
			if mag > largest {
				largest = mag
			}
		}
	}
	fmt.Println("Part 2:", largest)
}
