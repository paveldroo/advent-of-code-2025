package main

import (
	"fmt"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

func main() {
	input := utils.MustReadFromFile("input.txt")
	matrix := [][]string{}
	for _, row := range input {
		splitRow := strings.Split(row, "")
		matrix = append(matrix, splitRow)
	}

	res := part1(matrix)
	// res := part2(matrix)

	fmt.Printf("TOTAL ROLLS: %d", res)
}

func part1(input [][]string) int {
	fmt.Println(input)
	return 0
}

func part2(input [][]string) int {
	return 0
}
