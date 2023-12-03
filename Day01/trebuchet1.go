package main

import (
	"strconv"
)

func extractDigits(str string) []int {
	digits := []int{}
	for _, char := range str {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		}
	}
	return digits
}
