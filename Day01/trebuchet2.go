package main

import (
	"strconv"
	"strings"
)

func extractDigitsAndSpeltStrings(str string) []int {
	spellingMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digits := []int{}
	for i, char := range str {
		if digit, err := strconv.Atoi(string(char)); err == nil {
			digits = append(digits, digit)
		} else {
			for spelling, number := range spellingMap {
				if strings.HasPrefix(str[i:], spelling) {
					digits = append(digits, number)
					break
				}
			}
		}
	}

	return digits
}
