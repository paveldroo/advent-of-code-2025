package main

import (
	"fmt"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string

func main() {
	input := utils.MustReadFromFile("input.txt")

	// res := part1(input)
	res := part2(input)

	fmt.Printf("TOTAL ROLLS: %d", res)
}

func part1(input []string) int {
	steps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	res := 0

	for _, row := range input {
		splitRow := strings.Split(row, "")
		matrix = append(matrix, splitRow)
	}

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == "@" {
				curRolls := 1
				for _, step := range steps {
					r := row + step[0]
					c := col + step[1]
					if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0]) {
						if matrix[r][c] == "@" {
							curRolls++
						}
					}
				}
				if curRolls <= 4 {
					res++
				}
			}
		}
	}

	return res
}

func part2(input []string) int {
	steps := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	res := 0

	for _, row := range input {
		splitRow := strings.Split(row, "")
		matrix = append(matrix, splitRow)
	}
	removed := true

	for removed {
		removed = false
		for row := range matrix {
			for col := range matrix[row] {
				if matrix[row][col] == "@" {
					curRolls := 1
					for _, step := range steps {
						r := row + step[0]
						c := col + step[1]
						if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0]) {
							if matrix[r][c] == "@" {
								curRolls++
							}
						}
					}
					if curRolls <= 4 {
						res++
						matrix[row][col] = "."
						removed = true
					}
				}
			}
		}
	}

	return res
}
