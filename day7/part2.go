package day7

import (
	"AoC_24/utils"
	"fmt"
	"strconv"
)

func checkEquationsP2(equations []Equation) int {
	calibSum := 0
	for _, eq := range equations {
		binTable := createRepTable(len(eq.operands)-1, 3)
		for _, rep := range binTable {
			calc := eq.operands[0]
			for i := 0; i < len(rep); i++ {
				if string(rep[i]) == "0" {
					calc += eq.operands[i+1]
				} else if string(rep[i]) == "1" {
					calc *= eq.operands[i+1]
				} else {
					leftStr := strconv.Itoa(calc)
					rightStr := strconv.Itoa(eq.operands[i+1])
					calc, _ = strconv.Atoi(leftStr + rightStr)
				}
			}
			if eq.result == calc {
				calibSum += calc
				break
			}
		}
	}
	return calibSum
}

func PrintPart2() {
	input, _ := utils.ReadLines("day7/input.txt")
	eqs := parseEquations(input)
	fmt.Println("AoC 24 Day 7, Part 2:", checkEquationsP2(eqs))

}
