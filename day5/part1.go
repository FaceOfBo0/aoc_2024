package day5

import (
	"AoC_24/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Fst int
	Snd int
}

func (r *Rule) Invert() {
	first := r.Fst
	r.Fst = r.Snd
	r.Snd = first
}

func parseRules(input []string) []Rule {
	rules := make([]Rule, 0, len(input))
	for _, elem := range input {
		first, _ := strconv.Atoi(elem[:2])
		second, _ := strconv.Atoi(elem[3:])
		rules = append(rules, Rule{Fst: first, Snd: second})
	}
	return rules
}

func parseUpdates(input []string) [][]int {
	listUpdates := make([][]int, 0, len(input))
	for _, elem := range input {
		updatesSplit := strings.Split(elem, ",")
		update := utils.MapList(updatesSplit, func(elem string) int { num, _ := strconv.Atoi(elem); return num })
		listUpdates = append(listUpdates, update)
	}
	return listUpdates
}

func getCorrectUpdates(rules []Rule, updates [][]int) [][]int {
	correctUpdates := make([][]int, 0)

	for _, upd := range updates {
		correct := true
		for i := range upd {
			for j := i + 1; j < len(upd); j++ {
				pair := Rule{Fst: upd[i], Snd: upd[j]}
				if !slices.Contains(rules, pair) {
					correct = false
					break
				}
			}
		}
		if correct {
			correctUpdates = append(correctUpdates, upd)
		}
	}

	return correctUpdates
}

func calcMiddleSum(input [][]int) int {
	middleSum := 0
	for _, elem := range input {
		middleSum += elem[(len(elem)-1)/2]
	}
	return middleSum
}

func PrintPart1() {
	input, _ := utils.ReadLines("day5/input.txt")
	rulesRaw, updatesRaw := utils.SplitListOnce(input, "")
	rules := parseRules(rulesRaw)
	updates := parseUpdates(updatesRaw)
	corrUpdates := getCorrectUpdates(rules, updates)
	fmt.Println("AoC 24 Day 5, Part 1:", calcMiddleSum(corrUpdates))

}
