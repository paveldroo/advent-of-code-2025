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
