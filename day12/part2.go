package day12

// Costs:
// test 1.0 = 1206
// test 1.1 = 80
// test 1.2 = 436
// test 2.0 = 236
// test 2.1 = 368

import (
	"AoC_24/utils"
	"fmt"
)

func PrintPart2() {
	input, _ := utils.ReadLines("day12/test1.0.txt")
	result := parseRegions(input)
	fmt.Println("AoC 24 Day 12, Part 2:", calcCost(result))
}
