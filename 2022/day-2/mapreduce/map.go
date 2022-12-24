package mapreduce

import (
	"errors"
	"fmt"
	"strings"
)

type StrMapper interface {
	Map([][]string) (int, error)
}

type round struct {
	play1 string
	play2 string
}

var (
	ErrUnimplemented     = errors.New("unimplemented")
	ErrUnrecognizedRound = errors.New("round must container two letters")
	ErrUnrecognizedShape = errors.New("unrecognized shape")
	ErrUnrecognizedBonus = errors.New("unrecognized bonus")
)

type RPCMapper struct{}

func (m *RPCMapper) Map(input []string) (int, error) {

	output := 0

	for _, line := range input {

		shapes := strings.Split(line, " ")

		if len(shapes) != 2 {
			return -1, fmt.Errorf("%v: %v", ErrUnrecognizedRound, shapes)
		}

		r := round{shapes[0], shapes[1]}

		shapeBonus, err := shapeBonus(r)
		if err != nil {
			return -1, err
		}

		roundBonus, err := roundBonus(r)
		if err != nil {
			return -1, err
		}

		output += shapeBonus + roundBonus
	}

	return output, nil
}

func shapeBonus(r round) (int, error) {

	shapeToBonus := map[round]int{
		{"A", "X"}: 3, // Scissors loses to rock.
		{"B", "X"}: 1, // Rock loses to paper.
		{"C", "X"}: 2, // Paper loses to scissors.

		{"A", "Y"}: 1, // Rock draws with rock.
		{"B", "Y"}: 2, // Paper draws with paper.
		{"C", "Y"}: 3, // Scissors draws with scissors.

		{"A", "Z"}: 2, // Paper beats rock.
		{"B", "Z"}: 3, // Scissors beats paper.
		{"C", "Z"}: 1, // Rock beats scissors.
	}

	bonus, ok := shapeToBonus[r]
	if !ok {
		return bonus, fmt.Errorf("round %v: %w", r, ErrUnrecognizedShape)
	}

	return bonus, nil
}

func roundBonus(r round) (int, error) {

	roundToBonus := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	bonus, ok := roundToBonus[r.play2]
	if !ok {
		return bonus, fmt.Errorf("play %v: %w", r.play2, ErrUnrecognizedBonus)
	}

	return bonus, nil
}
