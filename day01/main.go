package main

import (
	"fmt"
	"strconv"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

func main() {
	input := utils.MustReadFromFile("input.txt")
	// res := part1(input)
	res := part2(input)

	fmt.Printf("TOTAL ZEROES COUNT: %d", res)
}

func part1(input []string) int {
	cur := 50
	zeroesCnt := 0

	for _, in := range input {
		dir, strVal := in[:1], in[1:]
		val, err := strconv.Atoi(strVal)
		if err != nil {
			panic(fmt.Errorf("strconv.Atoi(strVal): %w", err))
		}

		if dir == "L" {
			next := cur - val
			next = next % 100
			if next < 0 {
				cur = 100 + next
			} else if next == 0 {
				cur = 0
				zeroesCnt++
			} else {
				cur = next
			}
		} else if dir == "R" {
			next := cur + val
			next = next % 100
			if next > 100 {
				cur = next - 99
			} else if next == 100 || next == 0 {
				cur = 0
				zeroesCnt++
			} else {
				cur = next
			}
		}
	}

	return zeroesCnt
}

func part2(input []string) int {
	return 0
}
