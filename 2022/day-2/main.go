package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/alowayed/adventofcode/2022/day-2/mapreduce"
)

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2022/day-2/input.txt"
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	score := 0
	mapper := &mapreduce.RPCMapper{}

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		singleScore, err := mapper.Map([]string{line})
		if err != nil {
			return err
		}
		score += singleScore
	}

	log.Printf("score: %v", score)

	return nil
}

func main() {
	start := time.Now()
	log.Println("--- Starting")

	if err := run(); err != nil {
		log.Printf("exit: %v", err)
	}

	end := time.Now()
	duration := end.Sub(start)
	log.Printf("--- Ended after %v\n", duration)
}
