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
			name:  "two invalid ids",
			input: []string{"11-22"},
			want:  33,
		},
		{
			name:  "full test case",
			input: []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224", "1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"},
			want:  1227775554,
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
			name:  "two invalid ids",
			input: []string{"95-115"},
			want:  210,
		},
		{
			name:  "one invalid ids",
			input: []string{"1212121211-1212121212"},
			want:  1212121212,
		},
		{
			name:  "two invalid ids",
			input: []string{"998-1012"},
			want:  999 + 1010,
		},
		{
			name:  "two invalid ids",
			input: []string{"565653-565659"},
			want:  565656,
		},
		{
			name:  "two invalid ids",
			input: []string{"824824821-824824827"},
			want:  824824824,
		},
		{
			name:  "two invalid ids",
			input: []string{"2121212118-2121212124"},
			want:  2121212121,
		},
		{
			name:  "full test case",
			input: []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224", "1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659", "824824821-824824827", "2121212118-2121212124"},
			want:  4174379265,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := part2(tc.input)
			require.Equal(t, tc.want, res)
		})
	}
}
