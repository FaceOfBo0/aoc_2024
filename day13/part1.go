package day13

import (
	"AoC_24/utils"
	"fmt"
)

type Button struct {
	incX int
	incY int
}

type BtnConfig struct {
	btnA   Button
	btnB   Button
	prizeX int
	prizeY int
}

func NewBtnConfig(a Button, b Button) BtnConfig {
	return BtnConfig{
		btnA: a,
		btnB: b,
	}
}

func parseBtnConfigs(input []string) {

}
func PrintPart1() {
	input, _ := utils.ReadLines("day13/test.txt")
	for i, v := range input {
		fmt.Printf("%v: %v\n", i, v)
	}
	// fmt.Println("AoC 24 Day 13, Part 1:", 0)
}
