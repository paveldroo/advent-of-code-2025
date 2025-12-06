package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string

func main() {
	input := utils.MustReadFromFile("input.txt")

	// res := part1(input)
	res := part2(input)

	fmt.Printf("TOTAL MATH: %d", res)
}

func part1(input []string) int {
	var matrix [][]string
	for _, row := range input {
		var curRow []string
		last := ""
		for _, letter := range row {
			if string(letter) == " " {
				if last != "" {
					curRow = append(curRow, last)
					last = ""
				}
				continue
			}
			last += string(letter)
		}
		if last != "" {
			curRow = append(curRow, last)
		}
		matrix = append(matrix, curRow)
	}

	res := 0

	for i := range matrix[0] {
		op := matrix[len(matrix)-1][i]
		curRes := digit(matrix[0][i])
		for j := 1; j < len(matrix)-1; j++ {
			if op == "*" {
				curRes *= digit(matrix[j][i])
			} else if op == "+" {
				curRes += digit(matrix[j][i])
			} else {
				panic("not valid operator: " + op)
			}
		}
		res += curRes
	}

	return res
}

func part2(input []string) int {
	var matrix [][]string

	indexes := [][]int{}
	curLen := 0
	for i := len(input[0]) - 1; i >= 0; i-- {
		curLen++
		if string(input[len(input)-1][i]) != " " {
			indexes = append(indexes, []int{i, curLen})
			curLen = 0
			i--
		}
	}

	slices.Reverse(indexes)

	fmt.Println(indexes)

	for _, row := range input {
		curRow := []string{}
		curDigitStr := ""
		for _, coords := range indexes {
			i, l := coords[0], coords[1]
			for j := i; j < i+l; j++ {
				if string(row[j]) == " " {
					curDigitStr += "0"
				} else if string(row[j]) == "*" || string(row[j]) == "+" {
					curDigitStr = string(row[j])
					break
				} else {
					curDigitStr += string(row[j])
				}
			}
			curRow = append(curRow, curDigitStr)
			curDigitStr = ""
		}
		matrix = append(matrix, curRow)
	}

	fmt.Println("matrix:", matrix)

	res := 0

	for i := range matrix[0] {
		op := matrix[len(matrix)-1][i]

		curRes := 0
		for j := range matrix[0][i] {
			digitStr := ""
			for k, row := range matrix {
				if k == len(matrix)-1 {
					continue
				}

				if string(row[i][j]) == "0" {
					continue
				}

				digitStr += string(row[i][j])
			}

			if curRes == 0 {
				curRes = digit(digitStr)
			} else if op == "*" {
				curRes *= digit(digitStr)
			} else if op == "+" {
				curRes += digit(digitStr)
			} else {
				panic("not valid operator: " + op)
			}
		}
		res += curRes
	}

	return res
}

func digit(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic("not valid integer: " + input)
	}

	return i
}
