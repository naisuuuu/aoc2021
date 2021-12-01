package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const template = `package main

import (
	"log"

	"github.com/naisuuuu/aoc2021/input"
)

func main() {
	in, err := input.Read()
	if err != nil {
		log.Fatal(err)
	}
	if err := solve(in); err != nil {
		log.Fatal(err)
	}
}

func solve(in []string) error {
	return nil
}
`

func main() {
	day := flag.Int("d", 1, "day of AOC")
	flag.Parse()
	dir := fmt.Sprintf("day%d", *day)
	_, err := os.Stat(dir)
	if err == nil {
		log.Fatalf("%s already exists", dir)
	}
	if !errors.Is(err, fs.ErrNotExist) {
		log.Fatalf("Error stating %s: %v", dir, err)
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("Error creating dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(template), 0644); err != nil {
		log.Fatalf("Error creating main.go: %v", err)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home dir: %v", err)
	}
	dl := filepath.Join(home, "Downloads", "input")
	if _, err := os.Stat(dl); err != nil {
		log.Println("No input file in downloads. Returning early")
		return
	}
	if err := os.Rename(dl, filepath.Join(dir, "input")); err != nil {
		log.Fatalf("Error moving input from downloads: %v", err)
	}
}
