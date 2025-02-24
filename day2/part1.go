package day2

import (
	"AoC_24/utils"
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func readInReports() []Report {
	file, _ := os.Open("day2/input.txt")
	scanner := bufio.NewScanner(file)
	repList := make([]Report, 0)
	for scanner.Scan() {
		numStr := strings.Split(scanner.Text(), " ")
		numList := utils.MapList(numStr, func(elem string) int { num, _ := strconv.Atoi(elem); return num })
		repList = append(repList, Report{levels: numList})
	}
	return repList
}

func allIncr(level []int) bool {
	incrCnt := 0
	comparisons := len(level) - 1
	for i := 0; i < comparisons; i++ {
		if level[i] < level[i+1] {
			incrCnt++
		}
	}
	return incrCnt == comparisons
}

func allDecr(level []int) bool {
	decrCnt := 0
	comparisons := len(level) - 1
	for i := 0; i < comparisons; i++ {
		if level[i] > level[i+1] {
			decrCnt++
		}
	}
	return decrCnt == comparisons
}

func checkConditions(level []int) int {
	comparisons := len(level) - 1
	incrFlag := allIncr(level)
	decrFlag := allDecr(level)

	if !(incrFlag || decrFlag) {
		return 0
	}
	for i := 0; i < comparisons; i++ {
		diff := math.Abs(float64(level[i]) - float64(level[i+1]))
		if !(diff == 1 || diff == 2 || diff == 3) {
			return 0
		}
	}
	return 1
}

func PrintPart1() {
	reportList := readInReports()
	safeReps := 0
	for _, rep := range reportList {
		safeReps += checkConditions(rep.levels)
	}

	println("AoC 24 Day 2, Part 1:", safeReps)
}
