package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(stringData []string) int {
	sum := 0
	desiredCountMap := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}
	for i, game := range stringData {
		gameId := i + 1
		flag := 1
		revealSets := strings.Split(strings.Split(game, ":")[1], ";")

		for _, revealSet := range revealSets {
			colorCountMap := map[string]int{
				"blue":  0,
				"red":   0,
				"green": 0,
			}
			individualColors := strings.Split(revealSet, ",")
			for _, individualColor := range individualColors {
				countColor := strings.Split(strings.TrimSpace(individualColor), " ")
				count, err := strconv.Atoi(strings.TrimSpace(countColor[0]))
				if err != nil {
					fmt.Println("Got invalid string", err)
				}
				color := strings.TrimSpace(countColor[1])
				colorCountMap[color] = count
			}
			for color, count := range colorCountMap {
				if desiredCountMap[color] < count {
					flag = 0
					break
				}
			}
			if flag == 0 {
				break
			}
		}

		if flag == 1 {
			sum += gameId
		}
	}
	return sum
}

func part2(stringData []string) int {
	result := 0
	for _, game := range stringData {
		revealSets := strings.Split(strings.Split(game, ":")[1], ";")
		colorCountMap := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}

		for _, revealSet := range revealSets {
			individualColors := strings.Split(revealSet, ",")
			for _, individualColor := range individualColors {
				countColor := strings.Split(strings.TrimSpace(individualColor), " ")
				count, err := strconv.Atoi(strings.TrimSpace(countColor[0]))
				if err != nil {
					fmt.Println("Got invalid string", err)
					return 0
				}
				color := strings.TrimSpace(countColor[1])
				if colorCountMap[color] < count {
					colorCountMap[color] = count
				}
			}
		}
		product := colorCountMap["red"] * colorCountMap["blue"] * colorCountMap["green"]
		result += product
	}

	return result
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	stringData := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Day 2 >> Puzzle 1 >> ", part1(stringData))
	fmt.Println("Day 2 >> Puzzle 2 >> ", part2(stringData))

}
