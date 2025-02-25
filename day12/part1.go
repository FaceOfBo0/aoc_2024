package day12

import (
	"AoC_24/utils"
	"fmt"
	"slices"
)

type Patch struct {
	plantType string
	area      int
	perimeter int
	borderMap map[string]int
	coords    []utils.Point
}

func (p *Patch) updateBorder() {
	if len(p.coords) > 1 {
		thisY, nextY := p.coords[len(p.coords)-2].Y, p.coords[len(p.coords)-1].Y
		thisX, nextX := p.coords[len(p.coords)-2].X, p.coords[len(p.coords)-1].X

		if nextY == thisY && nextX == thisX+1 {
			if !slices.Contains(p.coords, utils.Point{X: nextX, Y: nextY - 1}) {
				p.borderMap["N"]++
				p.borderMap["S"]++
			}
		}
	}
}

func NewPatch(plantType string) *Patch {
	bmap := map[string]int{"N": 1, "E": 1, "S": 1, "W": 1}
	return &Patch{
		plantType: plantType,
		area:      0,
		borderMap: bmap,
		perimeter: 0,
		coords:    make([]utils.Point, 0),
	}
}

func parseGarden(input []string) map[string]*Patch {
	regionMap := make(map[string]*Patch)
	for y, row := range input {
		for x, elem := range row {
			char := string(elem)
			if key, ok := regionMap[char]; ok {
				key.coords = append(key.coords, utils.Point{Y: y, X: x})
				key.area++
			} else {
				regionMap[char] = NewPatch(char)
				p := regionMap[char]
				p.coords = append(p.coords, utils.Point{Y: y, X: x})
				p.area++
			}
		}
	}
	return regionMap
}

func PrintPart1() {
	input, _ := utils.ReadLines("day12/test1.txt")
	result := parseGarden(input)
	for k, p := range result {
		fmt.Printf("%v: %v\n", k, p)
	}
	// fmt.Println("AoC 24 Day 12, Part 1:", 0)
}
