package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

func main() {
	input := utils.MustReadFromFile("input.txt")
	// res := part1(input)
	res := part2(input)

	fmt.Printf("TOTAL JOLTS: %d", res)
}

func part1(input []string) int {
	var res int
	for _, bat := range input {

		q := queuePart1{}
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

type queuePart1 struct {
	store [2]int
}

func (q *queuePart1) push(val int) {
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

func part2(input []string) int {
	var res int
	for _, bat := range input {

		q := queuePart2{}
		for _, val := range bat {
			valInt, err := strconv.Atoi(string(val))
			if err != nil {
				panic(fmt.Errorf("not valid int string(val): %v", val))
			}
			q.push(valInt)
		}

		res += q.digit()
	}

	return res
}

type queuePart2 struct {
	store []int
}

func (q *queuePart2) push(val int) {
	if len(q.store) < 12 {
		q.store = append(q.store, val)
		return
	}

	for i := range q.store {
		if i == len(q.store)-1 {
			break
		}

		if q.store[i] < q.store[i+1] {
			for j := i; j < len(q.store)-1; j++ {
				q.store[j] = q.store[j+1]
			}
			q.store[11] = val
			return
		}
	}

	if q.store[11] <= val {
		q.store[11] = val
	}
}

func (q *queuePart2) digit() int {
	var strRes []string
	for _, val := range q.store {
		strRes = append(strRes, strconv.Itoa(val))
	}

	res, err := strconv.Atoi(strings.Join(strRes, ""))
	if err != nil {
		panic(fmt.Errorf("not valid int strRes: %v", strRes))
	}

	return res
}
