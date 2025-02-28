package day12

import (
	"AoC_24/utils"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type Region struct {
	plantType  string
	areas      []int
	perimeters []int
	overlaps   []int
	patches    []mapset.Set[Coordinate]
}

type Coordinate struct {
	idx int
	pnt utils.Point
}

type Direction int

const (
	Root Direction = iota
	Right
	Down
	Left
	Up
)

func NewRegion(plantType string) *Region {
	return &Region{
		plantType: plantType,
		areas:     make([]int, 0),
		overlaps:  make([]int, 0),
		patches:   make([]mapset.Set[Coordinate], 0),
	}
}

func (p *Region) calcPerimeters() {
	p.perimeters = make([]int, len(p.areas))
	for i, area := range p.areas {
		p.perimeters[i] = area * (area*4 - (p.overlaps[i] * 2))
	}
}

func (p *Region) findPatchRec(grid []string, char string, setIdx int, yPos, xPos int, dir Direction) {
	p.areas[setIdx]++
	p.patches[setIdx].Add(Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos}})
	leftCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos - 1}}
	rightCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos, X: xPos + 1}}
	downCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos + 1, X: xPos}}
	upCd := Coordinate{idx: setIdx, pnt: utils.Point{Y: yPos - 1, X: xPos}}

	if !p.patches[setIdx].Contains(rightCd) {
		if xPos < len(grid[0])-1 {
			if string(grid[yPos][xPos+1]) == char {
				p.overlaps[setIdx]++
				p.findPatchRec(grid, char, setIdx, yPos, xPos+1, Right)
			}
		}
	} else if !(dir == Left) {
		p.overlaps[setIdx]++
	}

	if !p.patches[setIdx].Contains(downCd) {
		if yPos < len(grid)-1 {
			if string(grid[yPos+1][xPos]) == char {
				p.overlaps[setIdx]++
				p.findPatchRec(grid, char, setIdx, yPos+1, xPos, Down)
			}
		}
	} else if !(dir == Up) {
		p.overlaps[setIdx]++
	}

	if !p.patches[setIdx].Contains(leftCd) {
		if xPos > 0 {
			if string(grid[yPos][xPos-1]) == char {
				p.overlaps[setIdx]++
				p.findPatchRec(grid, char, setIdx, yPos, xPos-1, Left)
			}
		}
	}

	if !p.patches[setIdx].Contains(upCd) {
		if yPos > 0 {
			if string(grid[yPos-1][xPos]) == char {
				p.overlaps[setIdx]++
				p.findPatchRec(grid, char, setIdx, yPos-1, xPos, Up)
			}
		}
	}

	return
}

func parseRegions(input []string) map[string]*Region {
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
					patch.findPatchRec(input, char, newIdx, y, x, Root)
				}
			} else {
				regionMap[char] = NewRegion(char)
				p := regionMap[char]
				p.patches = append(p.patches, mapset.NewSet[Coordinate]())
				p.areas = append(p.areas, 0)
				p.overlaps = append(p.overlaps, 0)
				p.findPatchRec(input, char, 0, y, x, Root)
			}
		}
	}
	return regionMap
}

func calcCost(garden map[string]*Region) int {
	totalCost := 0
	for _, patch := range garden {
		patch.calcPerimeters()
		for _, per := range patch.perimeters {
			totalCost += per
		}
	}
	return totalCost
}

func PrintPart1() {
	input, _ := utils.ReadLines("day12/input.txt")
	result := parseRegions(input)
	fmt.Println("AoC 24 Day 12, Part 1:", calcCost(result))
}
