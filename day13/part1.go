package day13

import (
	"fmt"

	linalg "gonum.org/v1/gonum/mat"
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

func parseBtnConfigs(input []string) []BtnConfig {
	btnCfgList := make([]BtnConfig, 0)
	return btnCfgList
}
func PrintPart1() {
	A := linalg.NewDense(2, 2, []float64{94, 22, 34, 67})
	b := linalg.NewVecDense(2, []float64{8400, 5400})
	var result linalg.VecDense
	result.SolveVec(A, b)
	fmt.Printf("result: %v\n", result.RawVector().Data)
	// input, _ := utils.ReadLines("day13/test.txt")
	// for i, v := range input {
	// 	fmt.Printf("%v: %v\n", i, v)
	// }
	// fmt.Println("AoC 24 Day 13, Part 1:", 0)
}
