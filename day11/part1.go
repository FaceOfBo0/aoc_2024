package day11

import (
	"AoC_24/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseNums(input string) []uint64 {
	nums := utils.MapList(strings.Split(input, " "), func(elem string) uint64 {
		num, _ := strconv.ParseUint(elem, 10, 64)
		return num
	})
	return nums
}

func splitNum(num uint64) (uint64, uint64) {
	numStr := strconv.FormatUint(num, 10)
	leftHalf, _ := strconv.ParseUint(numStr[:len(numStr)/2], 10, 64)
	rightHalf, _ := strconv.ParseUint(numStr[len(numStr)/2:], 10, 64)
	return leftHalf, rightHalf
}

func blink(stones []uint64, blinkCnt int) []uint64 {

	for range blinkCnt {
		newStones := make([]uint64, 0, len(stones)*2)
		for _, num := range stones {
			numStr := strconv.FormatUint(num, 10)
			if num == 0 {
				newStones = append(newStones, 1)
			} else if len(numStr)%2 == 0 {
				lft, rht := splitNum(num)
				newStones = append(newStones, lft, rht)
			} else {
				newStones = append(newStones, num*2024)
			}
		}
		stones = slices.Clone(newStones)
	}

	return stones
}

func PrintPart1() {
	input, _ := utils.ReadLines("day11/input.txt")
	nums := parseNums(input[0])
	fmt.Println("AoC 24 Day 11, Part 1:", len(blink(nums, 25)))
}
