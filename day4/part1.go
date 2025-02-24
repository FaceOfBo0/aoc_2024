package day4

import (
	"AoC_24/utils"
	"fmt"
)

func xmasRight(input string, j int) bool {
	return string(input[j]) == "X" && string(input[j+1]) == "M" &&
		string(input[j+2]) == "A" && string(input[j+3]) == "S"
}

func xmasLeft(input string, j int) bool {
	return string(input[j]) == "X" && string(input[j-1]) == "M" &&
		string(input[j-2]) == "A" && string(input[j-3]) == "S"
}

func xmasDown(input []string, i int, j int) bool {
	return string(input[i][j]) == "X" && string(input[i+1][j]) == "M" &&
		string(input[i+2][j]) == "A" && string(input[i+3][j]) == "S"
}

func xmasUp(input []string, i int, j int) bool {
	return string(input[i][j]) == "X" && string(input[i-1][j]) == "M" &&
		string(input[i-2][j]) == "A" && string(input[i-3][j]) == "S"
}

func xmasRightUp(input []string, i int, j int) bool {
	return string(input[i][j]) == "X" && string(input[i-1][j+1]) == "M" &&
		string(input[i-2][j+2]) == "A" && string(input[i-3][j+3]) == "S"
}

func xmasRightDown(input []string, i int, j int) bool {
	return string(input[i][j]) == "X" && string(input[i+1][j+1]) == "M" &&
		string(input[i+2][j+2]) == "A" && string(input[i+3][j+3]) == "S"
}

func xmasLeftUp(input []string, i int, j int) bool {
	return string(input[i][j]) == "X" && string(input[i-1][j-1]) == "M" &&
		string(input[i-2][j-2]) == "A" && string(input[i-3][j-3]) == "S"
}

func xmasLeftDown(input []string, i int, j int) bool {
	return string(input[i][j]) == "X" && string(input[i+1][j-1]) == "M" &&
		string(input[i+2][j-2]) == "A" && string(input[i+3][j-3]) == "S"
}

func searchXMAS(input []string) int {
	xmasCnt := 0
	for i := range input {
		for j := range input[i] {
			// check horizontal (right) XMAS
			if j+3 <= len(input[i])-1 {
				if xmasRight(input[i], j) {
					xmasCnt++
				}
				// check diagonal (upRight, downRight) XMAS
				if i+3 <= len(input)-1 {
					if xmasRightDown(input, i, j) {
						xmasCnt++
					}
				}
				if i >= 3 {
					if xmasRightUp(input, i, j) {
						xmasCnt++
					}
				}
			}
			// check horizontal (left) XMAS
			if j >= 3 {
				if xmasLeft(input[i], j) {
					xmasCnt++
				}
				// check diagonal (leftDown, leftUp) XMAS
				if i+3 <= len(input)-1 {
					if xmasLeftDown(input, i, j) {
						xmasCnt++
					}
				}
				if i >= 3 {
					if xmasLeftUp(input, i, j) {
						xmasCnt++
					}
				}
			}

			// check vertical (down, up) XMAS
			if i+3 <= len(input)-1 {
				if xmasDown(input, i, j) {
					xmasCnt++
				}
			}

			if i >= 3 {
				if xmasUp(input, i, j) {
					xmasCnt++
				}
			}
		}
	}

	return xmasCnt
}

func PrintPart1() {
	input, _ := utils.ReadLines("day4/input.txt")
	fmt.Println("AoC 24 Day 4, Part 1:", searchXMAS(input))
}
