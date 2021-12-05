package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	aocURL   = "https://adventofcode.com/%d/day/%d"
	template = `package main

import (
	"flag"
	"fmt"
	"log"

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

func solve(in []string) {
	fmt.Println("Part 1:", in)
}
`
)

func main() {
	day := flag.Int("d", 1, "day of AOC")
	year := flag.Int("y", 2021, "year of AOC")
	flag.Parse()

	if err := gen(*day, *year); err != nil {
		log.Fatalln("Error", err)
	}
}

func gen(day, year int) error {
	dir := fmt.Sprintf("day%d", day)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("creating dir: %w", err)
	}

	if err := writeTemplate(filepath.Join(dir, "main.go")); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	if err := downloadExample(filepath.Join(dir, "example"), year, day); err != nil {
		log.Printf("Error downloading example: %v", err)
	}

	cookie, err := os.ReadFile(".cookie")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Println(".cookie not found. Input will not be downloaded.")
			return nil
		}
		return fmt.Errorf("reading .cookie: %v", err)
	}

	if err := downloadInput(
		strings.Trim(string(cookie), " \n"),
		filepath.Join(dir, "input"),
		year, day); err != nil {
		return fmt.Errorf("downloading input: %v", err)
	}

	return nil
}

func downloadInput(cookie, fpath string, year, day int) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(aocURL, year, day)+"/input", nil)
	if err != nil {
		return fmt.Errorf("preparing request: %w", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("input download non-200 status code: %d", res.StatusCode)
	}

	f, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, res.Body); err != nil {
		return fmt.Errorf("copying body: %w", err)
	}

	return nil
}

func downloadExample(fpath string, year, day int) error {
	res, err := http.Get(fmt.Sprintf(aocURL, year, day))
	if err != nil {
		return fmt.Errorf("getting puzzle body: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("non-200 status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("parsing puzzle body: %w", err)
	}

	example := doc.FindMatcher(goquery.Single(".day-desc pre code")).Text()
	if len(strings.Trim(example, " \n")) == 0 {
		return fmt.Errorf("zero length example: '%s'", example)
	}

	f, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(example); err != nil {
		return fmt.Errorf("writing example: %w", err)
	}

	return nil
}

func writeTemplate(fpath string) error {
	_, err := os.Stat(fpath)
	if err == nil {
		log.Printf("%s already exists. Template will not be written", fpath)
		return nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("stat %s: %w", fpath, err)
	}

	if err := os.WriteFile(fpath, []byte(template), 0644); err != nil {
		return fmt.Errorf("creating main.go: %w", err)
	}

	return nil
}
