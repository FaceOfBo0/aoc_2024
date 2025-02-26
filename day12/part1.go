package day12

import (
	"AoC_24/utils"
	"fmt"
	"slices"
)

type Border map[string]int
type Coordinates []utils.Point

type Patch struct {
	plantType string
	area      int
	perimeter int
	borderMap Border
	coords    []Coordinates
	points    Coordinates
}

func (p *Patch) getPerimeter() int {
	result := 0
	for _, val := range p.borderMap {
		result += val
	}

	return result
}

func (p *Patch) updateBorders() {
	for _, coord := range p.coords {
		thisY, nextY := coord[len(p.coords)-2].Y, coord[len(p.coords)-1].Y
		thisX, nextX := coord[len(p.coords)-2].X, coord[len(p.coords)-1].X

		if nextY == thisY && nextX == thisX+1 && !slices.Contains(coord, utils.Point{X: nextX, Y: nextY - 1}) {
			p.borderMap["N"]++
			p.borderMap["S"]++
			p.perimeter += 2

		} else if nextY == thisY+1 {

			if slices.Contains(coord, utils.Point{X: nextX, Y: nextY - 1}) {
				p.borderMap["W"]++
				p.borderMap["E"]++
				p.perimeter += 2
			} else {
				p.borderMap["N"]++
				p.borderMap["W"]++
				p.borderMap["E"]++
				p.perimeter += 3
			}
		}
	}
}

func NewPatch(plantType string) *Patch {
	bmap := map[string]int{"N": 1, "E": 1, "S": 1, "W": 1}
	return &Patch{
		plantType: plantType,
		area:      0,
		perimeter: 4,
		borderMap: bmap,
		coords:    make([]Coordinates, 0),
		points:    make(Coordinates, 0),
	}
}

func calcCost(garden map[string]*Patch) int {
	totalCost := 0
	for _, patch := range garden {
		totalCost += patch.area * patch.perimeter
	}
	return totalCost
}

func calcBorders() {

}

func parsePatches(input []string) map[string]*Patch {
	regionMap := make(map[string]*Patch)
	for y, row := range input {
		for x, elem := range row {
			char := string(elem)
			if key, ok := regionMap[char]; ok {
				key.points = append(key.points, utils.Point{Y: y, X: x})
				key.area++
			} else {
				regionMap[char] = NewPatch(char)
				p := regionMap[char]
				p.points = append(p.points, utils.Point{Y: y, X: x})
				p.area++
			}
		}
	}
	return regionMap
}

func PrintPart1() {
	input, _ := utils.ReadLines("day12/test1.txt")
	result := parsePatches(input)
	for k, p := range result {
		fmt.Printf("%v: %v\n", k, p)
	}
	fmt.Println("AoC 24 Day 12, Part 1:", calcCost(result))
}
