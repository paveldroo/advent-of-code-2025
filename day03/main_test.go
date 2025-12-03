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
			name:  "first",
			input: []string{"987654321111111"},
			want:  98,
		},
		{
			name:  "second",
			input: []string{"987654321111191"},
			want:  99,
		},
		{
			name:  "third",
			input: []string{"234234234234278"},
			want:  78,
		},
		{
			name:  "forth",
			input: []string{"818181911112111"},
			want:  92,
		},
		{
			name:  "full test",
			input: []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"},
			want:  357,
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
			name:  "first",
			input: []string{"987654321111111"},
			want:  987654321111,
		},
		{
			name:  "second",
			input: []string{"811111111111119"},
			want:  811111111119,
		},
		{
			name:  "my second",
			input: []string{"811111111111919"},
			want:  811111111919,
		},
		{
			name:  "third",
			input: []string{"234234234234278"},
			want:  434234234278,
		},
		{
			name:  "forth",
			input: []string{"818181911112111"},
			want:  888911112111,
		},
		{
			name:  "my",
			input: []string{"888181911112111"},
			want:  888911112111,
		},
		{
			name:  "full test",
			input: []string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"},
			want:  3121910778619,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := part2(tc.input)
			require.Equal(t, tc.want, res)
		})
	}
}
