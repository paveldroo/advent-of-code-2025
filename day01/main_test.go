package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "one zero",
			input: []string{"L50", "R70", "R100"},
			want:  1,
		},
		{
			name:  "two zero",
			input: []string{"L50", "L1", "R1"},
			want:  2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := part1(tc.input)
			require.Equal(t, tc.want, res)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "two zero",
			input: []string{"L50", "R70", "R100"},
			want:  2,
		},
		{
			name:  "three zeros",
			input: []string{"L51", "R1", "R120"},
			want:  3,
		},
		{
			name:  "four zeros",
			input: []string{"L51", "R1", "R220"},
			want:  4,
		},
		{
			name:  "full test case from adventofcode",
			input: []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
			want:  6,
		},
		{
			name:  "full test case from adventofcode plus hundreds",
			input: []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82", "R668", "L965", "L89"},
			want:  23,
		},
		{
			name:  "minus 161",
			input: []string{"L161"},
			want:  2,
		},
		{
			name:  "more zeros with hundreds",
			input: []string{"L50", "R600", "R100", "R10", "L110", "L500"},
			want:  15,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := part2(tc.input)
			require.Equal(t, tc.want, res)
		})
	}
}
