package main

import (
	"fmt"
	"strconv"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

func main() {
	input := utils.MustReadFromFile("input.txt")
	res := part1(input)
	// res := part2(input)

	fmt.Printf("TOTAL JOLTS: %d", res)
}

func part1(input []string) int {
	var res int
	for _, bat := range input {

		q := queue{}
		for _, val := range bat {
			valInt, err := strconv.Atoi(string(val))
			if err != nil {
				panic(fmt.Errorf("not valid int string(val): %v", val))
			}
			q.push(valInt)
		}
		strRes := strconv.Itoa(q.store[0]) + strconv.Itoa(q.store[1])
		resToAdd, err := strconv.Atoi(strRes)
		if err != nil {
			panic(fmt.Errorf("not valid int strRes: %v", resToAdd))
		}
		res += resToAdd
	}

	return res
}

// func part2(input []string) int {
// 	return 0
// }

type queue struct {
	store [2]int
}

func (q *queue) push(val int) {
	if q.store[0] == 0 {
		q.store[0] = q.store[1]
		q.store[1] = val
		return
	}
	if q.store[1] <= val || q.store[0] < q.store[1] {
		if q.store[0] < q.store[1] {
			q.store[0] = q.store[1]
		}
		q.store[1] = val
	}
}
