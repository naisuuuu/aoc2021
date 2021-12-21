package main

import (
	"flag"
	"fmt"
	"log"
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

type Image [][]byte

func (i Image) Count() (ones int) {
	for row := range i {
		for col := range i[0] {
			ones += int(i[row][col])
		}
	}
	return
}

func (i Image) String() string {
	var b strings.Builder
	for j, col := range i {
		for _, c := range col {
			switch c {
			case 1:
				b.WriteRune('#')
			case 0:
				b.WriteRune('.')
			}
		}
		if j < len(i)-1 {
			b.WriteRune('\n')
		}
	}
	return b.String()
}

func (i Image) enhancePixel(row, col int, infinite bool) (index uint) {
	for r := -1; r <= 1; r++ {
		nr := row + r
		if nr >= len(i) || nr < 0 {
			index <<= 3
			if infinite {
				index += 7
			}
			continue
		}
		for c := -1; c <= 1; c++ {
			index <<= 1
			nc := col + c
			if nc >= len(i[0]) || nc < 0 {
				if infinite {
					index++
				}
				continue
			}
			if i[nr][nc] == 1 {
				index++
			}
		}
	}
	return
}

func (i Image) enhance(lookup []byte, infinite bool) Image {
	var n Image
	for row := -1; row <= len(i); row++ {
		ll := make([]byte, len(i[0])+2)
		for col := -1; col <= len(i[0]); col++ {
			ll[col+1] = lookup[i.enhancePixel(row, col, infinite)]
		}
		n = append(n, ll)
	}
	return n
}

func (i Image) EnhanceTimes(lookup []byte, times int) Image {
	var alternating, infinite bool
	if lookup[0] == 1 && lookup[511] == 0 {
		alternating = true
		fmt.Println("alternating")
	}
	for j := 0; j < times; j++ {
		i = i.enhance(lookup, infinite)
		if alternating {
			infinite = !infinite
		}
	}
	return i
}

func translate(c byte) byte {
	switch c {
	case '#':
		return 1
	default:
		return 0
	}
}

func solve(in []string) {
	var lookup []byte
	i := 0
	for i < len(in) {
		if len(in[i]) == 0 {
			break
		}
		for _, c := range in[i] {
			switch c {
			case '.':
				lookup = append(lookup, 0)
			case '#':
				lookup = append(lookup, 1)
			}
		}
		i++
	}

	var image Image
	for _, l := range in[i+1:] {
		ll := make([]byte, len(l))
		for i := range l {
			ll[i] = translate(l[i])
		}
		image = append(image, ll)
	}

	fmt.Println("Part 1:", image.EnhanceTimes(lookup, 2).Count())
	fmt.Println("Part 2:", image.EnhanceTimes(lookup, 50).Count())
}
