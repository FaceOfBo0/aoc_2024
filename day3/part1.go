package day3

import (
	"AoC_24/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseMull(code string) int {
	multPattern := regexp.MustCompile(`mul\((\d{1,3},\d{1,3})\)`)
	multSubmatches := multPattern.FindAllStringSubmatch(code, -1)
	partRes := 0

	for _, elem := range multSubmatches {
		partRes += parseMult(elem)
	}

	return partRes
}

func parseMult(text []string) int {
	nums := strings.Split(text[1], ",")
	numA, _ := strconv.Atoi(nums[0])
	numB, _ := strconv.Atoi(nums[1])
	return numA * numB
}

func readInProgram() []string {
	file, _ := os.Open("day3/input.txt")
	scanner := bufio.NewScanner(file)
	linesProgram := make([]string, 0)

	for scanner.Scan() {
		linesProgram = append(linesProgram, scanner.Text())
	}

	return linesProgram
}

func PrintPart1() {
	program := utils.FoldlStr(readInProgram())
	result := parseMull(program)
	fmt.Println("AoC 24 Day 3, Part 1:", result)
}
