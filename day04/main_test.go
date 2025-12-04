package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input [][]string
		want  int
	}{
		{
			name: "first",
			input: [][]string{
				{"..@@.@@@@."},
				{"@@@.@.@.@@"},
				{"@@@@@.@.@@"},
				{"@.@@@@..@."},
				{"@@.@@@@.@@"},
				{".@@@@@@@.@"},
				{".@.@.@.@@@"},
				{"@.@@@.@@@@"},
				{".@@@@@@@@."},
				{"@.@.@@@.@."},
			},
			want: 13,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := part1(tc.input)
			require.Equal(t, tc.want, res)
		})
	}
}
