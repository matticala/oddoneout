package main

import (
	"errors"
	"fmt"
)

func Solve(input []int, maxCount int) (int, error) {
	buffer := make(map[int]int, len(input))

	for _, v := range input {
		buffer[v] += 1
		if buffer[v] == maxCount {
			delete(buffer, v)
		}
	}

	if len(buffer) > 1 {
		return -1, fmt.Errorf("Input had more than a unique number. [%v]", buffer)
	}

	for k, _ := range buffer {
		return k, nil
	}

	return -1, errors.New("No unique elements in the input.")
}
