package day5

import (
	"AoC_24/utils"
	"fmt"
	"slices"
)

func SwitchElements(list []int, pair Rule) {
	fstIdx := slices.Index(list, pair.Fst)
	sndIdx := slices.Index(list, pair.Snd)
	list[sndIdx] = pair.Fst
	list[fstIdx] = pair.Snd
}

func getIncorrectUpdates(rules []Rule, updates [][]int) [][]int {
	incorrectUpdates := make([][]int, 0)

	for _, upd := range updates {
		correct := true
		for i := range upd {
			for j := i + 1; j < len(upd); j++ {
				pair := Rule{Fst: upd[i], Snd: upd[j]}
				if !slices.Contains(rules, pair) {
					correct = false
					SwitchElements(upd, pair)
				}
			}
		}
		if !correct {
			incorrectUpdates = append(incorrectUpdates, upd)
		}
	}

	return incorrectUpdates
}

func PrintPart2() {
	input, _ := utils.ReadLines("day5/input.txt")
	rulesRaw, updatesRaw := utils.SplitListOnce(input, "")
	rules := parseRules(rulesRaw)
	updates := parseUpdates(updatesRaw)
	incorrUpdates := getIncorrectUpdates(rules, updates)
	fmt.Println("AoC 24 Day 5, Part 2:", calcMiddleSum(incorrUpdates))
}
