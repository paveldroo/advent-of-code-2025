package main

import (
	"fmt"
	"strconv"

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
	var matrix [][]string
	for _, row := range input {
		// fmt.Println("INOUT ROW:", row)
		var curRow []string
		last := ""
		// fmt.Println("====================ROW:", row)
		for _, letter := range row {
			// fmt.Println("LETTER:", string(letter))
			// fmt.Println("LAST:", last)
			if string(letter) == " " {
				if last != "" {
					curRow = append(curRow, last)
					// fmt.Println("CURROW:", curRow)
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

	// fmt.Println("matrix:", matrix)

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

func digit(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic("not valid integer: " + input)
	}

	return i
}
