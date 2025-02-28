package day12

// Costs:
// test 1.0 = 1206
// test 1.1 = 80
// test 1.2 = 436
// test 2.0 = 236
// test 2.1 = 368

import (
	"AoC_24/utils"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

func parseRegionsP2(input []string) map[string]*Region {
	regionMap := make(map[string]*Region)
	for y, row := range input {
		for x, elem := range row {
			char := string(elem)
			if patch, ok := regionMap[char]; ok {
				isIn := false
				// check if plant is already in a patch
				for i, elem := range patch.patches {
					thisCd := Coordinate{idx: i, pnt: utils.Point{Y: y, X: x}}
					if elem.Contains(thisCd) {
						isIn = true
						break
					}
				}

				if !isIn {
					patch.patches = append(patch.patches, mapset.NewSet[Coordinate]())
					newIdx := len(patch.patches) - 1
					patch.areas = append(patch.areas, 0)
					patch.overlaps = append(patch.overlaps, 0)
					patch.sides = append(patch.sides, 0)
					patch.findPatchRecP2(input, newIdx, y, x, Root)
				}
			} else {
				regionMap[char] = NewRegion(char)
				p := regionMap[char]
				p.patches = append(p.patches, mapset.NewSet[Coordinate]())
				p.areas = append(p.areas, 0)
				p.overlaps = append(p.overlaps, 0)
				p.sides = append(p.sides, 0)
				p.findPatchRecP2(input, 0, y, x, Root)
			}
		}
	}
	return regionMap
}

func (p *Region) calcSidesVert(grid []string, setIdx int, thisCd Coordinate, nextCd Coordinate) {
	if thisCd.pnt.X == 0 {
		p.sides[setIdx]++
		thisRightChar, nextRightChar := string(grid[thisCd.pnt.Y][thisCd.pnt.X+1]), string(grid[nextCd.pnt.Y][nextCd.pnt.X+1])

		if !(thisRightChar == p.plantType || nextRightChar == p.plantType) {
			p.sides[setIdx]++
		}
	} else if thisCd.pnt.X == len(grid[0])-1 {
		p.sides[setIdx]++
		thisLeftChar, nextLeftChar := string(grid[thisCd.pnt.Y][thisCd.pnt.X-1]), string(grid[nextCd.pnt.Y][nextCd.pnt.X-1])

		if !(thisLeftChar == p.plantType || nextLeftChar == p.plantType) {
			p.sides[setIdx]++
		}
	} else {
		thisRightChar, nextRightChar := string(grid[thisCd.pnt.Y][thisCd.pnt.X+1]), string(grid[nextCd.pnt.Y][nextCd.pnt.X+1])
		thisLeftChar, nextLeftChar := string(grid[thisCd.pnt.Y][thisCd.pnt.X-1]), string(grid[nextCd.pnt.Y][nextCd.pnt.X-1])

		if !(thisRightChar == p.plantType || nextRightChar == p.plantType) {
			p.sides[setIdx]++
		}

		if !(thisLeftChar == p.plantType || nextLeftChar == p.plantType) {
			p.sides[setIdx]++
		}
	}
}

func (p *Region) calcSidesHorz(grid []string, setIdx int, thisCd Coordinate, nextCd Coordinate) {
	if thisCd.pnt.Y == 0 {
		p.sides[setIdx]++
		thisDownChar, nextDownChar := string(grid[thisCd.pnt.Y+1][thisCd.pnt.X]), string(grid[nextCd.pnt.Y+1][nextCd.pnt.X])

		if !(thisDownChar == p.plantType || nextDownChar == p.plantType) {
			p.sides[setIdx]++
		}
	} else if thisCd.pnt.Y == len(grid)-1 {
		p.sides[setIdx]++
		thisUpChar, nextUpChar := string(grid[thisCd.pnt.Y-1][thisCd.pnt.X]), string(grid[nextCd.pnt.Y-1][nextCd.pnt.X])

		if !(thisUpChar == p.plantType || nextUpChar == p.plantType) {
			p.sides[setIdx]++
		}
	} else {
		thisDownChar, nextDownChar := string(grid[thisCd.pnt.Y+1][thisCd.pnt.X]), string(grid[nextCd.pnt.Y+1][nextCd.pnt.X])
		thisUpChar, nextUpChar := string(grid[thisCd.pnt.Y-1][thisCd.pnt.X]), string(grid[nextCd.pnt.Y-1][nextCd.pnt.X])

		if !(thisDownChar == p.plantType || nextDownChar == p.plantType) {
			p.sides[setIdx]++
		}

		if !(thisUpChar == p.plantType || nextUpChar == p.plantType) {
			p.sides[setIdx]++
		}
	}
}

func (p *Region) findPatchRecP2(grid []string, setIdx int, yPos, xPos int, dir Direction) {
	p.areas[setIdx]++
	p.patches[setIdx].Add(Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos}})
	thisCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos}}
	leftCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos - 1}}
	rightCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos + 1}}
	downCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos + 1, X: xPos}}
	upCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos - 1, X: xPos}}

	if !p.patches[setIdx].Contains(rightCd) {
		if xPos < len(grid[0])-1 {
			if string(grid[yPos][xPos+1]) == p.plantType {
				p.overlaps[setIdx]++
				p.calcSidesHorz(grid, setIdx, thisCd, rightCd)
				p.findPatchRecP2(grid, setIdx, yPos, xPos+1, Right)
			}
		}
	} else if !(dir == Left) {
		p.overlaps[setIdx]++
		p.calcSidesHorz(grid, setIdx, thisCd, rightCd)
	}

	if !p.patches[setIdx].Contains(downCd) {
		if yPos < len(grid)-1 {
			if string(grid[yPos+1][xPos]) == p.plantType {
				p.overlaps[setIdx]++
				p.calcSidesVert(grid, setIdx, thisCd, downCd)
				p.findPatchRecP2(grid, setIdx, yPos+1, xPos, Down)
			}
		}
	} else if !(dir == Up) {
		p.overlaps[setIdx]++
		p.calcSidesVert(grid, setIdx, thisCd, downCd)
	}

	if !p.patches[setIdx].Contains(leftCd) {
		if xPos > 0 {
			if string(grid[yPos][xPos-1]) == p.plantType {
				p.overlaps[setIdx]++
				p.calcSidesHorz(grid, setIdx, thisCd, leftCd)
				p.findPatchRecP2(grid, setIdx, yPos, xPos-1, Left)
			}
		}
	}

	if !p.patches[setIdx].Contains(upCd) {
		if yPos > 0 {
			if string(grid[yPos-1][xPos]) == p.plantType {
				p.overlaps[setIdx]++
				p.calcSidesVert(grid, setIdx, thisCd, upCd)
				p.findPatchRecP2(grid, setIdx, yPos-1, xPos, Up)
			}
		}
	}

	return
}

func (p *Region) calcSides() {
	p.costs = make([]int, len(p.areas))
	for i, area := range p.areas {
		p.costs[i] = area * (area*4 - (p.overlaps[i] * 2) - p.sides[i])
	}
}

func calcCostP2(garden map[string]*Region) int {
	totalCost := 0
	for _, patch := range garden {
		patch.calcSides()
		for _, per := range patch.costs {
			totalCost += per
		}
	}
	return totalCost
}

func PrintPart2() {
	input, _ := utils.ReadLines("day12/input.txt")
	result := parseRegionsP2(input)
	fmt.Println("AoC 24 Day 12, Part 2:", calcCostP2(result))
}
