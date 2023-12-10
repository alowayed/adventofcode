package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alowayed/adventofcode/2022/day-3/rucksack"
	mr "github.com/alowayed/adventofcode/mapreduce"
)

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2022/day-3/input.txt"
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// PART 1

	// linesCh := make(chan string)
	// priorityCh, errCh := mr.Map(rucksack.Priority, linesCh)

	// s := bufio.NewScanner(file)
	// for s.Scan() {
	// 	line := s.Text()
	// 	linesCh <- line
	// }
	// close(linesCh)

	// PART 2

	rucksacksCh := make(chan []string)
	priorityCh, errCh := mr.Map(rucksack.Badge, rucksacksCh)

	s := bufio.NewScanner(file)
	rucksacks := []string{}
	numRucksacks := 0
	for s.Scan() {
		line := s.Text()
		rucksacks = append(rucksacks, line)
		numRucksacks += 1

		if numRucksacks == 3 {
			rucksacksCh <- rucksacks
			rucksacks = []string{}
			numRucksacks = 0
		}
	}
	close(rucksacksCh)

	totalPriority := 0
	for priority := range priorityCh {
		totalPriority += priority
	}

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("error mapping lines to priorities: %w", err)
		}
	default:
	}

	log.Printf("Total priority: %v", totalPriority)

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
