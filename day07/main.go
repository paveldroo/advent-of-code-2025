package main

import (
	"fmt"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string
var res int
var seen map[string]struct{}

func main() {
	input := utils.MustReadFromFile("input.txt")

	// res := part1(input)
	res := part2(input)

	fmt.Printf("TOTAL SPLITS: %d", res)
}

// func part1(input []string) int {
// 	res := 0
// 	matrix := [][]string{}

// 	for _, row := range input {
// 		curRow := strings.Split(row, "")
// 		matrix = append(matrix, curRow)
// 	}

// 	for i, row := range matrix {
// 		for j := range row {
// 			if row[j] == "S" {
// 				matrix[i+1][j] = "|"
// 				continue
// 			}
// 			if row[j] == "|" {
// 				if i < len(matrix)-1 {
// 					if matrix[i+1][j] == "^" {
// 						res++
// 						if j > 0 {
// 							matrix[i+1][j-1] = "|"
// 						}
// 						if j < len(row)-1 {
// 							matrix[i+1][j+1] = "|"
// 						}
// 					} else {
// 						matrix[i+1][j] = "|"
// 					}
// 				}
// 			}
// 		}
// 	}

// 	// fmt.Println("MATRIX:", matrix)
// 	return res
// }

func part2(input []string) int {
	matrix := [][]string{}
	seen = make(map[string]struct{})

	for _, row := range input {
		curRow := strings.Split(row, "")
		matrix = append(matrix, curRow)
	}
	dfs(matrix, 0)

	// fmt.Println("MATRIX:", matrix)
	return res
}

func dfs(matrix [][]string, curRowIdx int) {
	if curRowIdx == len(matrix)-1 {
		res++
		fmt.Println("RES:", res)
		return
	}

	printMatrix(matrix)

	nextRow := matrix[curRowIdx+1]
	row := matrix[curRowIdx]
	for j := range row {
		if row[j] == "S" {
			nextRow[j] = "|"
			dfs(matrix, curRowIdx+1)
			return
		}
		if row[j] == "|" {
			if nextRow[j] == "^" {
				if j > 0 {
					nextRow[j-1] = "|"
					dfs(matrix, curRowIdx+1)
					nextRow[j-1] = "."
				}
				if j < len(row)-1 {
					nextRow[j+1] = "|"
					dfs(matrix, curRowIdx+1)
					nextRow[j+1] = "."
				}
			} else {
				nextRow[j] = "|"
				dfs(matrix, curRowIdx+1)
				nextRow[j+1] = "."
			}
		}
	}
}

func printMatrix(input [][]string) {
	fmt.Println("======NEW MATRIX========")
	for i, row := range input {
		fmt.Println(fmt.Sprintf("%d%v", i, row))
	}
}
