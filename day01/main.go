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
	zerosCnt := 0

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
				zerosCnt++
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
				zerosCnt++
			} else {
				cur = next
			}
		}
	}

	return zerosCnt
}

func part2(input []string) int {
	cur := 50
	zerosCnt := 0

	for _, in := range input {
		dir, strVal := in[:1], in[1:]
		val, err := strconv.Atoi(strVal)
		if err != nil {
			panic(fmt.Errorf("strconv.Atoi(strVal): %w", err))
		}
		if dir == "L" {
			if cur > 0 {
				cur -= 100
			}
			cur -= val
			clicks := cur / 100
			zerosCnt -= clicks
			fmt.Println("zeroCnt:", zerosCnt)
			cur = cur % 100
			fmt.Println("CUR:", cur)
		} else if dir == "R" {
			if cur < 0 {
				cur += 100
			}
			cur += val
			clicks := cur / 100
			zerosCnt += clicks
			fmt.Println("zeroCnt:", zerosCnt)
			cur = cur % 100
			fmt.Println("CUR:", cur)
		}
	}

	return zerosCnt
}
