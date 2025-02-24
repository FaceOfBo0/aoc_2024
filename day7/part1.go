package day7

import (
	"AoC_24/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Equation struct {
	result   int
	operands []int
}

func parseEquations(input []string) []Equation {
	equations := make([]Equation, 0, len(input))

	for _, row := range input {
		rowSplit := strings.Split(row, ":")
		res, _ := strconv.Atoi(rowSplit[0])
		equations = append(equations, Equation{
			result: res,
			operands: utils.MapList(strings.Split(strings.TrimSpace(rowSplit[1]), " "),
				func(elem string) int {
					op, _ := strconv.Atoi(elem)
					return op
				}),
		})
	}

	return equations
}

func createRepTable(length int, base int) []string {
	table := make([]string, 0)
	for i := range int(math.Pow(float64(base), float64(length))) {
		bin := strconv.FormatInt(int64(i), base)
		for len(bin) < length {
			bin = "0" + bin
		}
		table = append(table, bin)
	}
	return table
}

func checkEquations(equations []Equation) int {
	calibSum := 0
	for _, eq := range equations {
		binTable := createRepTable(len(eq.operands)-1, 2)
		for _, rep := range binTable {
			calc := eq.operands[0]
			for i := 0; i < len(rep); i++ {
				if string(rep[i]) == "0" {
					calc += eq.operands[i+1]
				} else {
					calc *= eq.operands[i+1]
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

func PrintPart1() {
	input, _ := utils.ReadLines("day7/input.txt")
	eqs := parseEquations(input)
	fmt.Println("AoC 24 Day 7, Part 1:", checkEquations(eqs))

}
