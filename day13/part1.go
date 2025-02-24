package day13

import (
	"AoC_24/utils"
	"fmt"
)

func PrintPart1() {
	input, _ := utils.ReadLines("day13/input.txt")
	for i, v := range input {
		fmt.Printf("%v: %v\n", i, v)
	}
	// fmt.Println("AoC 24 Day 13, Part 1:", 0)
}
