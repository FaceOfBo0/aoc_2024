package day13

import (
	"AoC_24/utils"
	"fmt"
	"math"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Button struct {
	incX float64
	incY float64
}

type BtnConfig struct {
	btnA   Button
	btnB   Button
	prizeX float64
	prizeY float64
}

func NewBtnConfig(a Button, b Button, prX, prY float64) BtnConfig {
	return BtnConfig{
		btnA:   a,
		btnB:   b,
		prizeX: prX,
		prizeY: prY,
	}
}

func parseBtnConfigs(input []string) []BtnConfig {
	btnCfgList := make([]BtnConfig, 0)
	for i, v := range input {
		if strings.HasPrefix(v, "Button A") {
			elemsA, elemsB, prizes := strings.Split(v, ","), strings.Split(input[i+1], ","), strings.Split(input[i+2], ",")
			aX, _ := strconv.ParseFloat(elemsA[0][len(elemsA[0])-2:], 64)
			aY, _ := strconv.ParseFloat(elemsA[1][len(elemsA[1])-2:], 64)
			bX, _ := strconv.ParseFloat(elemsB[0][len(elemsB[0])-2:], 64)
			bY, _ := strconv.ParseFloat(elemsB[1][len(elemsB[1])-2:], 64)
			prX, _ := strconv.ParseFloat(prizes[0][len(prizes[0])-4:], 64)
			prY, _ := strconv.ParseFloat(prizes[1][len(prizes[1])-4:], 64)
			btnA := Button{incX: aX, incY: aY}
			btnB := Button{incX: bX, incY: bY}
			btnCfg := NewBtnConfig(btnA, btnB, prX, prY)
			btnCfgList = append(btnCfgList, btnCfg)
		}
	}
	return btnCfgList
}

func calculateCost(btnCfgs []BtnConfig) int {
	totalCost := 0
	for _, btncfg := range btnCfgs {
		// Aalt := utils.NewSqrMatrix(2, []float64{btncfg.btnA.incX, btncfg.btnB.incX, btncfg.btnA.incY, btncfg.btnB.incY})
		// balt := utils.NewVector(2, []float64{btncfg.prizeX, btncfg.prizeY})
		// res, _ := utils.SolveLinEq(Aalt, balt)
		//fmt.Printf("res: %v\n", res)
		A := mat.NewDense(2, 2, []float64{btncfg.btnA.incX, btncfg.btnB.incX, btncfg.btnA.incY, btncfg.btnB.incY})
		b := mat.NewVecDense(2, []float64{btncfg.prizeX, btncfg.prizeY})
		b.SolveVec(A, b)

		resA := b.RawVector().Data[0]
		resB := b.RawVector().Data[1]
		fmt.Printf("resA: %v\n", resA)
		fmt.Printf("resB: %v\n", resB)
		intA, fracA := math.Modf(resA)

		fracA, _ = math.Modf(fracA * 100000)
		if fracA == 99999 {
			intA++
			fracA = 0
		}

		intB, fracB := math.Modf(resB)
		fracB, _ = math.Modf(fracB * 100000)
		if fracB == 99999 {
			intB++
			fracB = 0
		}

		if (fracA == 0) &&
			(fracB == 0) &&
			(intA <= 100) &&
			(intB <= 100) &&
			(intA > 0) &&
			(intB > 0) {
			fmt.Printf("intA: %v\n", intA)
			fmt.Printf("intB: %v\n", intB)
			totalCost += int(intA*3 + intB)
		}
	}
	return totalCost
}
func PrintPart1() {
	input, _ := utils.ReadLines("day13/input.txt")
	btnCfgs := parseBtnConfigs(input)
	fmt.Println("AoC 24 Day 13, Part 1:", calculateCost(btnCfgs))
	a := 8
	b := 4
	f := a / b
	fmt.Println(f)
}
