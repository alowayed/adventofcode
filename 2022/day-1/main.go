package main

import (
	"bufio"
	"container/heap"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	InvalidArgumentErr = errors.New("invalide argument")
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxCalories(r io.Reader, numElves int) (int, error) {
	if numElves < 1 {
		return 0, InvalidArgumentErr
	}

	currCal := 0
	h := &IntHeap{}
	heap.Init(h)

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()

		// Add onto this Elf's calories.
		if line != "" {
			cal, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}

			currCal += cal
			continue
		}

		heap.Push(h, currCal)
		currCal = 0

		if h.Len() > numElves {
			heap.Pop(h)
		}
	}

	sum := 0
	for h.Len() > 0 {
		cal := heap.Pop(h)
		sum += cal.(int)
	}

	return sum, nil
}

// MaxCalories returns the max sum of calories carried by numElves found in the file at path.
func MaxCalories(path string, numElves int) (int, error) {

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return maxCalories(file, numElves)
}

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2022/Day 1/input.txt"
	maxCal, err := MaxCalories(path, 3)
	if err != nil {
		return err
	}

	log.Printf("Max calories carried is: %v", maxCal)

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
