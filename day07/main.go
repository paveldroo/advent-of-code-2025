package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string
var res int
var seen map[string]struct{}

func main() {
	input := utils.MustReadFromFile("input.txt")

	// result := part1(input)
	result := part2(input)

	fmt.Printf("TOTAL SPLITS: %d", result)
}

func part1(input []string) int {
	res := 0

	for _, row := range input {
		curRow := strings.Split(row, "")
		matrix = append(matrix, curRow)
	}

	for i, row := range matrix {
		for j := range row {
			if row[j] == "S" {
				matrix[i+1][j] = "|"
				continue
			}
			if row[j] == "|" {
				if i < len(matrix)-1 {
					if matrix[i+1][j] == "^" {
						res++
						if j > 0 {
							matrix[i+1][j-1] = "|"
						}
						if j < len(row)-1 {
							matrix[i+1][j+1] = "|"
						}
					} else {
						matrix[i+1][j] = "|"
					}
				}
			}
		}
	}

	// printMatrix(matrix)
	return res
}

func part2(input []string) int {
	_ = part1(input)
	seen = make(map[string]struct{})
	result := countBeams()
	// matrix := [][]string{}

	// for _, row := range input {
	// 	curRow := strings.Split(row, "")
	// 	matrix = append(matrix, curRow)
	// }

	// dfs(matrix, 0, "")

	// fmt.Println("MATRIX:", matrix)
	return result
}

func countBeams() int {
	for i, row := range matrix {
		fmt.Println("beams:", len(seen))
		for j, val := range row {
			if val == "S" {
				seen[dot(i, j)] = struct{}{}
				matrix[i][j] = "|"
				continue
			}
			if val == "|" {
				if matrix[i-1][j] == "|" {
					for key, _ := range seenCopy(seen) {
						if strings.HasSuffix(key, dot(i-1, j)) {
							// fmt.Println("DELETE BEAM:", key)
							delete(seen, key)
							newKey := beam(key, dot(i, j))
							// fmt.Println("ADD BEAM:", newKey)
							seen[newKey] = struct{}{}
						}
					}
				}
				matrix[i-1][j] = "."
			}
			if val == "^" {
				if matrix[i-1][j] == "|" {
					for key, _ := range seenCopy(seen) {
						if strings.HasSuffix(key, dot(i-1, j)) {
							// fmt.Println("DELETE BEAM:", key)
							delete(seen, key)
							if j > 0 {
								newKey := beam(key, dot(i, j-1))
								// fmt.Println("ADD BEAM:", newKey)
								seen[newKey] = struct{}{}
							}
							if j < len(row)-1 {
								newKey := beam(key, dot(i, j+1))
								// fmt.Println("ADD BEAM:", newKey)
								seen[newKey] = struct{}{}
							}
						}
					}
				}
				matrix[i-1][j] = "."
			}
		}
	}

	printMatrix(matrix)
	return len(seen) - 1
}

func dot(i, j int) string {
	return strconv.Itoa(i) + strconv.Itoa(j)
}

func beam(b string, point string) string {
	return b + "-" + point
}

func seenCopy(seen map[string]struct{}) map[string]struct{} {
	newSeen := make(map[string]struct{})
	for key, _ := range seen {
		newSeen[key] = struct{}{}
	}
	return newSeen
}

// func dfs(matrix [][]string, curRowIdx int, curRoute string) {
// 	if curRowIdx == len(matrix)-1 {
// 		res++
// 		fmt.Println("RES:", res)
// 		return
// 	}

// 	printMatrix(matrix)

// 	row := matrix[curRowIdx]
// 	for j := range row {
// 		if row[j] == "S" {
// 			newMatrix := matrixCopy(matrix)
// 			newMatrix[curRowIdx+1][j] = "|"
// 			// curRoute += fmt.Sprintf("%d-%d", curRowIdx+1, j)
// 			// seen[curRoute] = struct{}{}
// 			dfs(newMatrix, curRowIdx+1, curRoute)
// 			return
// 		}
// 		if row[j] == "|" {
// 			if matrix[curRowIdx+1][j] == "^" {
// 				if j > 0 {
// 					newMatrix := matrixCopy(matrix)
// 					newMatrix[curRowIdx+1][j-1] = "|"
// 					// curRoute += fmt.Sprintf("%d-%d", curRowIdx+1, j)
// 					// if _, ok := seen[curRoute]; !ok {
// 					// 	seen[curRoute] = struct{}{}
// 					dfs(newMatrix, curRowIdx+1, curRoute)
// 					// }
// 					// nextRow[j-1] = "."
// 				}
// 				if j < len(row)-1 {
// 					newMatrix := matrixCopy(matrix)
// 					newMatrix[curRowIdx+1][j+1] = "|"
// 					// curRoute += fmt.Sprintf("%d-%d", curRowIdx+1, j)

// 					// if _, ok := seen[curRoute]; !ok {
// 					// 	seen[curRoute] = struct{}{}
// 					dfs(newMatrix, curRowIdx+1, curRoute)
// 					// }
// 					// nextRow[j+1] = "."
// 				}
// 			} else {
// 				newMatrix := matrixCopy(matrix)
// 				newMatrix[curRowIdx+1][j] = "|"
// 				// if curRowIdx == 4 && j == 7 {
// 				// 	fmt.Println("IM HERE")
// 				// 	printMatrix(newMatrix)
// 				// }
// 				// curRoute += fmt.Sprintf("%d-%d", curRowIdx+1, j)

// 				// if _, ok := seen[curRoute]; !ok {
// 				// 	seen[curRoute] = struct{}{}
// 				dfs(newMatrix, curRowIdx+1, curRoute)
// 				// }
// 				// nextRow[j+1] = "."
// 			}
// 		}
// 	}
// }

func printMatrix(input [][]string) {
	fmt.Println("======NEW MATRIX========")
	for i, row := range input {
		idx := strconv.Itoa(i)
		if len(idx) == 1 {
			idx = "0" + idx
		}
		fmt.Println(fmt.Sprintf("%s%v", idx, row))
	}
}

// func matrixCopy(input [][]string) [][]string {
// 	matrix := make([][]string, len(input))
// 	copy(matrix, input)
// 	for i, row := range input {
// 		newRow := make([]string, len(row))
// 		copy(newRow, row)
// 		matrix[i] = newRow
// 	}

// 	return matrix
// }
