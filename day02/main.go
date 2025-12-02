package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paveldroo/advent-of-code-2025/utils"
)

func main() {
	input := utils.MustReadFromFile("input.txt")
	input = strings.Split(input[0], ",")
	// res := part1(input)
	res := part2(input)

	fmt.Printf("INVALID IDS SUM: %d", res)
}

func part1(input []string) int {
	sum := 0

	for _, rng := range input {
		start, end := startEnd(rng)
		for id := start; id <= end; id++ {
			strID := strconv.Itoa(id)
			middle := len(strID) / 2
			first, second := strID[:middle], strID[middle:]
			if first == second {
				sum += id
			}
		}
	}

	return sum
}

func part2(input []string) int {
	sum := 0

	for _, rng := range input {
		start, end := startEnd(rng)
		for id := start; id <= end; id++ {
			strID := strconv.Itoa(id)
			var times []int

			for i := 2; i <= len(strID); i++ {
				if len(strID)%i == 0 {
					times = append(times, i)
				}
			}

			if len(times) == 0 {
				continue
			}

			for _, time := range times {
				step := len(strID) / time
				seen := map[string]int{}
				for i := range time {
					startIdx, endIdx := step*i, step*i+step
					part := strID[startIdx:endIdx]
					if _, ok := seen[part]; !ok {
						seen[part] = 1
					} else {
						seen[part]++
					}
				}
				if len(seen) == 1 {
					sum += id
					break
				}

			}
		}
	}

	return sum
}

func startEnd(rng string) (int, int) {
	spl := strings.Split(rng, "-")
	if len(spl) != 2 {
		panic(fmt.Errorf("range split failed, spl val:%s", spl))
	}
	startStr, endStr := spl[0], spl[1]
	start, err := strconv.Atoi(startStr)
	if err != nil {
		panic(fmt.Errorf("strconv.Atoi(startStr): %w", err))
	}
	end, err := strconv.Atoi(endStr)
	if err != nil {
		panic(fmt.Errorf("strconv.Atoi(endStr): %w", err))
	}

	return start, end
}
