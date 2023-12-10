package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type color string

const (
	unkonwn = color("unknown")
	red     = color("red")
	blue    = color("blue")
	green   = color("green")
)

type set map[color]int

type game struct {
	id   int
	sets []set
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

func run() error {

	path := "/Users/yousef/go/src/github.com/alowayed/adventofcode/2023/day_2/input.txt"
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	games, err := parseGames(file)
	if err != nil {
		return err
	}

	// PART 1

	// cubes := map[color]int{
	// 	red:   12,
	// 	green: 13,
	// 	blue:  14,
	// }
	// var sumValidIDs int
	// for _, game := range games {
	// 	valid := validGame(game, cubes)
	// 	log.Printf("game %d is valid: %t", game.id, valid)
	// 	if valid {
	// 		sumValidIDs += game.id
	// 	}
	// }
	// log.Printf("--- sum of valid game IDs: %d", sumValidIDs)

	// PART 2

	sumPowers := 0
	for _, game := range games {
		minS := minSet(game)
		setP := setPower(minS)
		sumPowers += setP
	}
	log.Printf("--- sum of set powers: %d", sumPowers)

	return nil
}

func parseGames(r io.Reader) ([]game, error) {
	var games []game

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		game, err := parseGame(line)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func parseGame(s string) (game, error) {
	g := game{}

	s = strings.TrimPrefix(s, "Game ")
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return g, fmt.Errorf("invalid game: %q", s)
	}
	gameID, err := strconv.Atoi(parts[0])
	if err != nil {
		return g, fmt.Errorf("invalid game id: %q", parts[0])
	}
	setsString := strings.TrimSpace(parts[1])
	setStrings := strings.Split(setsString, "; ")

	var sets []set
	for _, setString := range setStrings {
		set, err := parseSet(setString)
		if err != nil {
			return g, err
		}
		sets = append(sets, set)
	}

	return game{
		id:   gameID,
		sets: sets,
	}, nil
}

// "3 blue, 4 red" ->
func parseSets(s string) ([]set, error) {
	var sets []set
	setsString := strings.Split(s, ", ")
	for _, setString := range setsString {
		set, err := parseSet(setString)
		if err != nil {
			return nil, err
		}
		sets = append(sets, set)
	}
	return sets, nil
}

// 3 blue, 4 red -> blue: 3, red: 4
func parseSet(s string) (set, error) {
	set := set{}

	revealStrings := strings.Split(s, ", ")
	for _, revealString := range revealStrings {
		color, num, err := parseReveal(revealString)
		if err != nil {
			return nil, err
		}
		set[color] = num
	}

	return set, nil
}

// 3 blue -> blue, 3
func parseReveal(s string) (color, int, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return unkonwn, 0, fmt.Errorf("invalid set: %q", s)
	}
	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return unkonwn, 0, fmt.Errorf("invalid set number: %q", parts[0])
	}
	color := color(parts[1])
	return color, num, nil
}

func validGame(g game, cubes map[color]int) bool {
	for _, set := range g.sets {
		if !validSet(set, cubes) {
			return false
		}
	}
	return true
}

func validSet(s set, cubes map[color]int) bool {
	for color, num := range s {
		if cubes[color] < num {
			return false
		}
	}
	return true
}

func minSet(g game) set {
	min := set{}
	for _, set := range g.sets {
		for color, num := range set {
			min[color] = max(min[color], num)
		}
	}
	return min
}

func setPower(s set) int {
	p := 1
	for _, num := range s {
		p *= num
	}
	return p
}
