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

func dfs(prevMatrix [][]string, curRowIdx int) {
	if len(prevMatrix) == 1 {
		res++
		return
	}

	// fmt.Println("SEEN:", seen)
	// printMatrix(prevMatrix, curRowIdx)

	matrix := make([][]string, len(prevMatrix))
	for i, row := range prevMatrix {
		newRow := make([]string, len(row))
		copy(newRow, row)
		matrix[i] = newRow
	}
	fmt.Println("RES:", res)
	// printMatrix(matrix, curRowIdx)

	nextRow := make([]string, len(matrix[1]))
	copy(nextRow, matrix[1])

	matrix[1] = nextRow
	row := matrix[0]
	for j := range row {
		if row[j] == "S" {
			nextRow[j] = "|"
			dfs(matrix[1:], curRowIdx+1)
		}
		if row[j] == "|" {
			if nextRow[j] == "^" {
				if j > 0 {
					// key := strconv.Itoa(curRowIdx) + "-" + strconv.Itoa(j) + "-" + strconv.Itoa(curRowIdx+1) + "-" + strconv.Itoa(j-1)
					// if _, ok := seen[key]; !ok {
					// fmt.Println("SET SEEN TO:", key)
					// seen[key] = struct{}{}
					// res++
					// fmt.Println("RES:", res)
					nextRow[j-1] = "|"
					dfs(matrix[1:], curRowIdx+1)
					nextRow[j-1] = "."
					// }
				}
				if j < len(row)-1 {
					// key := strconv.Itoa(curRowIdx) + "-" + strconv.Itoa(j) + "-" + strconv.Itoa(curRowIdx+1) + "-" + strconv.Itoa(j+1)
					// if _, ok := seen[key]; !ok {
					// fmt.Println("SET SEEN TO:", key)
					// seen[key] = struct{}{}
					// res++
					// fmt.Println("RES:", res)
					nextRow[j+1] = "|"
					dfs(matrix[1:], curRowIdx+1)
					nextRow[j+1] = "."
					// }
				}
			} else {
				nextRow[j] = "|"
				dfs(matrix[1:], curRowIdx+1)
			}
		}
	}
}

func printMatrix(input [][]string, curIdx int) {
	fmt.Println("======NEW MATRIX========")
	for i, row := range input {
		fmt.Println(fmt.Sprintf("%d%v", curIdx+i, row))
	}
}
