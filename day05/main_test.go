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
			name: "first",
			input: []string{
				"3-5",
				"4-5",
				"10-14",
				"12-18",
				"19-19",
				"16-20",
				"20-22",
				"23-23",
				"23-26",
				"",
				"123",
			},
			want: 3 + 13 + 1 + 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := part2(tc.input)
			require.Equal(t, tc.want, res)
		})
	}
}
