package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alowayed/adventofcode/2022/day-4/interval"
	mr "github.com/alowayed/adventofcode/mapreduce"
)

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2022/day-4/input.txt"
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// PART 1

	// containedFn := func(s string) (bool, error) {

	// 	intervals, err := interval.New(s)
	// 	if err != nil {
	// 		return false, fmt.Errorf("failed to convert string to interval: %w", err)
	// 	}

	// 	if len(intervals) != 2 {
	// 		return false, fmt.Errorf("line must contain exactly two intervals, found %v", intervals)
	// 	}

	// 	intervalA := intervals[0]
	// 	intervalB := intervals[1]

	// 	subIntervalA, err := interval.SubInterval(intervalA, intervalB)
	// 	if err != nil {
	// 		return false, fmt.Errorf("failed to check if %v is subinterval of %v: %w", intervalA, intervalB, err)
	// 	}

	// 	subIntervalB, err := interval.SubInterval(intervalB, intervalA)
	// 	if err != nil {
	// 		return false, fmt.Errorf("failed to check if %v is subinterval of %v: %w", intervalB, intervalA, err)
	// 	}

	// 	return subIntervalA || subIntervalB, nil
	// }

	// PART 2

	containedFn := func(s string) (bool, error) {

		intervals, err := interval.New(s)
		if err != nil {
			return false, fmt.Errorf("failed to convert string to interval: %w", err)
		}

		if len(intervals) != 2 {
			return false, fmt.Errorf("line must contain exactly two intervals, found %v", intervals)
		}

		intervalA := intervals[0]
		intervalB := intervals[1]

		overlaps, err := interval.Overlap(intervalA, intervalB)
		if err != nil {
			return false, fmt.Errorf("failed to check if %v and %v overlap: %w", intervalA, intervalB, err)
		}

		return overlaps, nil
	}

	linesCh := make(chan string)
	containedCh, errCh := mr.Map(containedFn, linesCh)

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		linesCh <- line
	}
	close(linesCh)

	totalContained := 0
	for contained := range containedCh {
		if contained {
			totalContained += 1
		}
	}

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("error mapping lines to contained: %w", err)
		}
	default:
	}

	log.Printf("Total contained: %v", totalContained)

	return nil
}

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
