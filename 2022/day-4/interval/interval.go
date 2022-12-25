package interval

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidInterval  = errors.New("invalid Interval")
	ErrInvalidArguments = errors.New("invalid arguments")
)

type Interval struct {
	Start int
	End   int
}

func Valid(i Interval) error {

	if i.Start > i.End {
		return fmt.Errorf("start must be the smaller or equal to end but found start=%v, end=%v: %w", i.Start, i.End, ErrInvalidInterval)
	}

	return nil
}

func SubInterval(inner, outter Interval) (bool, error) {

	for intervalName, interval := range map[string]Interval{"outter": outter, "inner": inner} {
		if err := Valid(interval); err != nil {
			return false, fmt.Errorf("invalid %v Interval %+v: %w", intervalName, interval, err)
		}
	}

	if outter.Start > inner.Start || outter.End < inner.End {
		return false, nil
	}

	return true, nil
}

func Overlap(a, b Interval) (bool, error) {

	if a.End < b.Start {
		return false, nil
	}
	if b.End < a.Start {
		return false, nil
	}

	return true, nil
}

// New converts a string of form: "START_1-END_1,START_2-END_2,...,START_N,END_N"into N intervals.
func New(s string) ([]Interval, error) {

	internalStrings := strings.Split(s, ",")
	intervals := []Interval{}
	for _, intervalString := range internalStrings {
		edges := strings.Split(intervalString, "-")
		if len(edges) != 2 {
			return nil, fmt.Errorf("interval can only have two edges, found %v: %w", edges, ErrInvalidArguments)
		}

		start, err := strconv.Atoi(edges[0])
		if err != nil {
			return nil, fmt.Errorf("interval edges must be ints, found %v: %w", edges[0], ErrInvalidArguments)
		}
		end, err := strconv.Atoi(edges[1])
		if err != nil {
			return nil, fmt.Errorf("interval edges must be ints, found %v: %w", edges[1], ErrInvalidArguments)
		}

		interval := Interval{Start: start, End: end}
		if err := Valid(interval); err != nil {
			return nil, fmt.Errorf("invalid interval while parsing start %v and end %v: %w", start, end, err)
		}

		intervals = append(intervals, Interval{Start: start, End: end})
	}

	return intervals, nil
}
