package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	digitRE = "zero|one|two|three|four|five|six|seven|eight|nine|\\d"
)

var (
	// firstPattern matches the first numeric or word digit in a string.
	firstPattern = fmt.Sprintf("(%s)", digitRE)
	// lastPattern matches the last numeric or word digit in a string.
	lastPattern = fmt.Sprintf(".*(%s)", digitRE)

	firstRE = regexp.MustCompile(firstPattern)
	lastRE  = regexp.MustCompile(lastPattern)

	numToNumeric = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func main() {
	start := time.Now()
	log.Println("--- Starting")

	defer func() {
		end := time.Now()
		duration := end.Sub(start)
		log.Printf("--- Ended after %v\n", duration)
	}()

	if err := run(); err != nil {
		log.Panicf("ERROR: %v", err)
	}
}

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2023/day_1/input.txt"
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// PART 1
		// num, err := extractNumerics(line)

		// PART 2
		num, err := extractNumbers(line)
		if err != nil {
			return fmt.Errorf("error extracting numerics from %q: %w", line, err)
		}
		total += num
	}
	log.Printf("total: %v", total)

	return nil
}

func extractNumerics(line string) (int, error) {
	var first int
	var last int
	for _, char := range line {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		last = num
		if first == 0 {
			first = num
		}
	}
	numS := fmt.Sprintf("%d%d", first, last)
	num, err := strconv.Atoi(numS)
	if err != nil {
		return 0, fmt.Errorf("error converting %v to int: %w", numS, err)
	}
	return num, nil
}

func extractNumbers(line string) (int, error) {

	firstMatches := firstRE.FindStringSubmatch(line)
	lastMatches := lastRE.FindStringSubmatch(line)

	firstMatch := firstMatches[1]
	lastMatch := lastMatches[1]

	first, err := toNumeric(firstMatch)
	if err != nil {
		return 0, fmt.Errorf("error converting %q to numeric: %w", firstMatch, err)
	}
	last, err := toNumeric(lastMatch)
	if err != nil {
		return 0, fmt.Errorf("error converting %q to numeric: %w", lastMatch, err)
	}

	numS := fmt.Sprintf("%d%d", first, last)
	num, err := strconv.Atoi(numS)
	if err != nil {
		return 0, fmt.Errorf("error converting %v to int: %w", numS, err)
	}
	return num, nil
}

func toNumeric(s string) (int, error) {
	if num, ok := numToNumeric[s]; ok {
		return num, nil
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error converting %q to int: %w", s, err)
	}
	return num, nil
}
