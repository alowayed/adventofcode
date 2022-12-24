package rucksack

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArguments = errors.New("invalid arguments")

	itemToPriority = map[string]int{}
)

func init() {
	items := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",

		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	for priority, item := range items {
		itemToPriority[item] = priority + 1
	}
}

func Priority(rucksack string) (int, error) {

	// -- Validate inputs.

	if len(rucksack)%2 == 1 {
		return 0, fmt.Errorf("rucksack must have even length, found odd %q: %w", rucksack, ErrInvalidArguments)
	}

	comp1 := rucksack[:len(rucksack)/2]
	comp2 := rucksack[len(rucksack)/2:]

	comp1Items := map[string]bool{}
	comp2Items := map[string]bool{}

	for _, item := range comp1 {
		comp1Items[string(item)] = true
	}
	for _, item := range comp2 {
		comp2Items[string(item)] = true
	}

	// -- Check for shared items between compartiments.

	sharedItems := map[string]bool{}
	for item := range comp1Items {
		if _, ok := comp2Items[item]; ok {
			sharedItems[item] = true
		}
	}

	// -- Validate outputs.

	if len(sharedItems) > 1 {
		return 0, fmt.Errorf("rucksack containes multiple shared components between compartments %q: %w", rucksack, ErrInvalidArguments)
	}
	if len(sharedItems) == 0 {
		return 0, fmt.Errorf("rucksack containes no shared components between compartments %q: %w", rucksack, ErrInvalidArguments)
	}

	priority := 0
	ok := false
	for item := range sharedItems {
		priority, ok = itemToPriority[item]
		if !ok {
			return 0, fmt.Errorf("rucksack containes unrecognized item %q: %w", item, ErrInvalidArguments)
		}
	}

	return priority, nil
}
