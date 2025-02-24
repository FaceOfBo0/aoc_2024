package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NumPair struct {
	leftNums  []int
	rightNums []int
}

func ReadInNums() NumPair {
	file, _ := os.Open("day1/input.txt")
	scanner := bufio.NewScanner(file)

	numList := NumPair{leftNums: make([]int, 0), rightNums: make([]int, 0)}
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		leftNum, _ := strconv.Atoi(nums[0])
		rightNum, _ := strconv.Atoi(nums[1])
		numList.leftNums = append(numList.leftNums, leftNum)
		numList.rightNums = append(numList.rightNums, rightNum)
	}

	return numList
}

func calcTotalDiff() int {
	numLists := ReadInNums()
	sort.Ints(numLists.leftNums)
	sort.Ints(numLists.rightNums)

	totalDiff := 0
	for i, elmL := range numLists.leftNums {
		if temp := elmL - numLists.rightNums[i]; temp < 0 {
			temp *= -1
			totalDiff += temp
		} else {
			totalDiff += temp
		}
	}
	return totalDiff
}

func PrintPart1() {
	println("AoC 24 Day 1, Part 1:", calcTotalDiff())
}
