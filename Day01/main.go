package main

import (
	"fmt"
	"os"
	"strings"
)

type extractFunc func(string) []int

// Common function that calculates sum of all the results once the digits are extracted
func getSumOfAllResults(lines []string, extractDigitFunction extractFunc) int {
	answer := 0
	for _, line := range lines {
		digits := extractDigitFunction(line)
		if len(digits) > 0 {
			first, last := digits[0], digits[len(digits)-1]
			answer += first*10 + last
		} else {
			fmt.Println("No digits found in the string")
		}
	}
	return answer
}

func main() {
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	// First Challenge
	fmt.Println("First Challenge Answer >>", getSumOfAllResults(lines, extractDigits))
	// Second Challenge
	fmt.Println("Second Challenge Answer >>", getSumOfAllResults(lines, extractDigitsAndSpeltStrings))
}
