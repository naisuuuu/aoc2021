package input

import (
	"bufio"
	"fmt"
	"os"
)

// Read reads an input file and returns a slice of lines within it.
func Read() ([]string, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scanning input: %w", err)
	}
	return lines, nil
}
