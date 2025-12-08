package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string
var res int
var seen map[string][]string

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
	seen = make(map[string][]string)
	result := countBeams()
	return result
}

func countBeams() int {
	for i, row := range matrix {
		for j, val := range row {
			if val == "S" {
				matrix[i][j] = "1"
				continue
			}
			if val == "|" {
				if matrix[i-1][j] != "." {
					matrix[i][j] = matrix[i-1][j]
				}
			}
			if val == "^" && matrix[i-1][j] != "." {
				du := digit(matrix[i-1][j])
				if j > 0 {
					if matrix[i][j-1] == "|" {
						matrix[i][j-1] = str(du)
					} else {
						dl := digit(matrix[i][j-1])
						matrix[i][j-1] = str(du + dl)
					}
				}
				if j < len(row)-1 {
					if matrix[i][j+1] == "|" {
						matrix[i][j+1] = str(du)
					} else {
						dr := digit(matrix[i][j+1])
						matrix[i][j+1] = str(du + dr)
					}
				}
			}
			if val != "." && val != "|" && val != "^" && matrix[i-1][j] != "." {
				d := digit(matrix[i][j])
				matrix[i][j] = str(d + digit(matrix[i-1][j]))
			}
		}
	}

	cnt := 0
	for _, val := range matrix[len(matrix)-1] {
		if val != "." {
			cnt += digit(val)
		}
	}

	return cnt
}

func digit(i string) int {
	val, err := strconv.Atoi(i)
	if err != nil {
		panic("not valid int")
	}
	return val
}

func str(i int) string {
	return strconv.Itoa(i)
}

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
