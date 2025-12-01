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
