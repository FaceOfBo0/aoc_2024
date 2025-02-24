package day4

import (
	"AoC_24/utils"
	"fmt"
)

func masMas(input []string, i int, j int) bool {
	return string(input[i][j]) == "M" && string(input[i+1][j+1]) == "A" && string(input[i+2][j+2]) == "S" &&
		string(input[i+2][j]) == "M" && string(input[i][j+2]) == "S"
}

func masSam(input []string, i int, j int) bool {
	return string(input[i][j]) == "M" && string(input[i+1][j+1]) == "A" && string(input[i+2][j+2]) == "S" &&
		string(input[i+2][j]) == "S" && string(input[i][j+2]) == "M"
}

func samMas(input []string, i int, j int) bool {
	return string(input[i][j]) == "S" && string(input[i+1][j+1]) == "A" && string(input[i+2][j+2]) == "M" &&
		string(input[i+2][j]) == "M" && string(input[i][j+2]) == "S"
}

func samSam(input []string, i int, j int) bool {
	return string(input[i][j]) == "S" && string(input[i+1][j+1]) == "A" && string(input[i+2][j+2]) == "M" &&
		string(input[i+2][j]) == "S" && string(input[i][j+2]) == "M"
}

func searchMAS(input []string) int {
	masCnt := 0

	for i := range input {
		for j := range input[i] {
			if i+2 <= len(input)-1 && j+2 <= len(input[i])-1 &&
				(masMas(input, i, j) || masSam(input, i, j) || samMas(input, i, j) || samSam(input, i, j)) {
				masCnt++
			}
		}
	}
	return masCnt
}

func PrintPart2() {
	input, _ := utils.ReadLines("day4/input.txt")
	fmt.Println("AoC 24 Day 4, Part 2:", searchMAS(input))
}
