package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

var matrix [][]string

func main() {
	input := utils.MustReadFromFile("input.txt")

	res := part1(input)
	// res := part2(input)

	fmt.Printf("TOTAL FRESH INGRIDIENTS: %d", res)
}

func part1(input []string) int {
	res := make(map[int]struct{})
	var ranges []string
	var fresh []string
	for i, row := range input {
		if row == "" {
			ranges = input[:i]
			fresh = input[i+1:]
			break
		}
	}

	var freshInt []int

	for _, f := range fresh {
		fInt, err := strconv.Atoi(f)
		if err != nil {
			panic(fmt.Sprintf("not valid fresh id: %s", f))
		}
		freshInt = append(freshInt, fInt)
	}

	slices.Sort(freshInt)
	// freshM := make(map[int]struct{})
	// for _, f := range fresh {
	// 	fInt, err := strconv.Atoi(f)
	// 	if err != nil {
	// 		panic(fmt.Sprintf("not valid fresh id: %s", f))
	// 	}
	// 	freshM[fInt] = struct{}{}
	// }

	for _, r := range ranges {
		rr := strings.Split(r, "-")
		start, err := strconv.Atoi(rr[0])
		if err != nil {
			panic(fmt.Sprintf("not valid range start: %s", rr[0]))
		}
		end, err := strconv.Atoi(rr[1])
		if err != nil {
			panic(fmt.Sprintf("not valid range end: %s", rr[1]))
		}

		for _, id := range freshInt {
			if id < start {
				continue
			}
			if id > end {
				break
			}
			res[id] = struct{}{}
		}

		// for i := start; i <= end; i++ {
		// 	if _, ok := freshM[i]; ok {
		// 		res[i] = struct{}{}
		// 	}
		// }
	}
	return len(res)
}

// func part2(input []string) int {
// 	return 0
// }
