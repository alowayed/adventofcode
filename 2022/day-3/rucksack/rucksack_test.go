package rucksack

import (
	"testing"
)

func TestPriority(t *testing.T) {
	type args struct {
		rucksack string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{{
		name: "valid rucksack 1",
		args: args{
			rucksack: "vJrwpWtwJgWrhcsFMMfFFhFp",
		},
		want: 16,
	}, {
		name: "valid rucksack 2",
		args: args{
			rucksack: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		},
		want: 38,
	}, {
		name: "valid rucksack 3",
		args: args{
			rucksack: "PmmdzqPrVvPwwTWBwg",
		},
		want: 42,
	}, {
		name: "no common item",
		args: args{
			rucksack: "vJrwxWtwJgWrhcsFMMfFFhFp",
		},
		wantErr: true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Priority(tt.args.rucksack)
			if (err != nil) != tt.wantErr {
				t.Errorf("Priority() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBadge(t *testing.T) {
	type args struct {
		rucksacks []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{{
		name: "valid rucksacks 1",
		args: args{
			rucksacks: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
		},
		want:    18,
		wantErr: false,
	}, {
		name: "valid rucksacks 1",
		args: args{
			rucksacks: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
		},
		want:    52,
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Badge(tt.args.rucksacks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Badge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Badge() = %v, want %v", got, tt.want)
			}
		})
	}
}
