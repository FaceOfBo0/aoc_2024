package day12

import (
	"AoC_24/utils"
	"fmt"
)

func PrintPart1() {
	input, _ := utils.ReadLines("day12/test1.txt")
	for i, v := range input {
		fmt.Printf("%v: %v\n", i, v)
	}
	// fmt.Println("AoC 24 Day 12, Part 1:", 0)
}
