package day2

import (
	"AoC_24/utils"
	"math"
)

func checkConditionsP2(level []int) int {
	incrFlag := allIncr(level)
	decrFlag := allDecr(level)
	newLvl := []int(nil)

	if !(incrFlag || decrFlag) {
		for i := range level {
			newLvl = utils.RemoveOne(level, i)
			incrFlag = allIncr(newLvl)
			decrFlag = allDecr(newLvl)
			if incrFlag || decrFlag {
				if diffFlag := allDiff(newLvl); diffFlag {
					return 1
				}
			}
		}
	} else {
		if diffFlag := allDiff(level); diffFlag {
			return 1
		} else {
			for i := range level {
				newLvl = utils.RemoveOne(level, i)
				if diffFlag = allDiff(newLvl); diffFlag {
					return 1
				}
			}
		}
	}

	return 0
}

func allDiff(level []int) bool {
	diffCnt := 0
	comparisons := len(level) - 1

	for i := 0; i < comparisons; i++ {
		diff := math.Abs(float64(level[i]) - float64(level[i+1]))
		if diff == 1 || diff == 2 || diff == 3 {
			diffCnt++
		}
	}

	return diffCnt == comparisons
}

func PrintPart2() {
	reportList := readInReports()
	safeReps := 0
	for _, rep := range reportList {
		safeReps += checkConditionsP2(rep.levels)
	}

	println("AoC 24 Day 2, Part 2:", safeReps)

}
