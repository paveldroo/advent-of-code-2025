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

	// res := part1(input)
	res := part2(input)

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

func part2(input []string) int {
	res := 0
	var ranges []string
	for i, row := range input {
		if row == "" {
			ranges = input[:i]
			break
		}
	}

	var rangesInt [][]int

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

		rangesInt = append(rangesInt, []int{start, end})
	}

	slices.SortFunc(rangesInt, func(a, b []int) int {
		if a[0] > b[0] {
			return 1
		}
		return -1
	})

	// fmt.Println("rangesInt", rangesInt)

	var rangesAdj [][]int

	for i := range rangesInt {
		start := rangesInt[i][0]
		end := rangesInt[i][1]
		cur := []int{start, end}
		if i > 0 {
			if end <= rangesAdj[len(rangesAdj)-1][1] {
				continue
			}
			if start <= rangesAdj[len(rangesAdj)-1][1] {
				rangesAdj[len(rangesAdj)-1][1] = end
				continue
			}
		}

		if i < len(rangesInt)-1 {
			startNext := rangesInt[i+1][0]
			endNext := rangesInt[i+1][1]
			if startNext <= end && end <= endNext {
				cur[1] = endNext
			}
		}

		rangesAdj = append(rangesAdj, cur)
	}

	// fmt.Println("rangesAdj:", rangesAdj)

	for _, r := range rangesAdj {
		// if r[1] <= r[0] {
		// 	fmt.Println("STRANGE RANGE:", r)
		// }

		// if (r[1] == 523263704114453 || r[0] == 523263704114453) && r[1] != r[0] {
		// 	fmt.Println("CORRUPTED RANGE:", r)
		// }
		cnt := r[1] - r[0] + 1
		res += cnt
	}

	return res
}
