package main

import (
	"fmt"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string

func main() {
	input := utils.MustReadFromFile("input.txt")

	res := part1(input)
	// res := part2(input)

	fmt.Printf("TOTAL MATH: %d", res)
}

func part1(input []string) int {
	return 0
}
