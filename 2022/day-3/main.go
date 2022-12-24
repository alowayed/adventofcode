package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	mr "github.com/alowayed/adventofcode/2022/day-3/mapreduce"
	"github.com/alowayed/adventofcode/2022/day-3/rucksack"
)

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2022/day-3/input.txt"
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	linesCh := make(chan string)
	priorityCh, errCh := mr.Map(rucksack.Priority, linesCh)

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		linesCh <- line
	}
	close(linesCh)

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
