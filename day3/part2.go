package day3

import (
	"AoC_24/utils"
	"fmt"
	"strings"
)

func parseProgramP2(code string) int {
	doSplit := strings.Split(code, "do()")
	partRes := 0
	for _, elem := range doSplit {
		doElem := strings.Split(elem, "don't()")[0]
		partRes += parseMull(doElem)
	}
	return partRes
}

func PrintPart2() {
	program := utils.FoldlStr(readInProgram())
	result := parseProgramP2(program)
	fmt.Println("AoC 24 Day 3, Part 2:", result)

}
