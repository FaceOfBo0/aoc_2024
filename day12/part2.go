package day12

import (
	"AoC_24/utils"
	"fmt"
)

func PrintPart2() {
	input, _ := utils.ReadLines("day12/input.txt")
	result := parseRegions(input)
	fmt.Println("AoC 24 Day 12, Part 2:", result)
}
