package utils

import (
	"bufio"
	"fmt"
	"os"
)

func MustReadFromFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("os.Open: %w", err))
	}

	scanner := bufio.NewScanner(f)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
