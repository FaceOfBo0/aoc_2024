package day10

import (
	"AoC_24/utils"
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"
)

type HikeMap struct {
	topMap       []string
	width        int
	trailCnt     int
	trailCntDist int
	trailEnds    mapset.Set[utils.Point]
}

func NewHikeMap(input []string) HikeMap {
	return HikeMap{
		topMap:       input,
		width:        len(input[0]),
		trailCnt:     0,
		trailCntDist: 0,
		trailEnds:    mapset.NewSet[utils.Point](),
	}
}

func (hm *HikeMap) parseMapv1() {
	for y, row := range hm.topMap {
		for x, col := range row {
			if elem, _ := strconv.Atoi(string(col)); elem == 0 {
				hm.searchTrailsv1(utils.Point{Y: y, X: x}, 0)
				hm.trailCnt += hm.trailEnds.Cardinality()
				hm.trailEnds.Clear()
			}
		}
	}
}

func (hm HikeMap) searchTrailsv1(pos utils.Point, height int) {
	nextHts := hm.findNextHeights(pos, height+1)
	if height+1 == 9 {
		for _, elem := range nextHts {
			hm.trailEnds.Add(elem)
		}
		return
	}
	for _, point := range nextHts {
		hm.searchTrailsv1(point, height+1)
	}
}

func (hm HikeMap) findNextHeights(pos utils.Point, nextHeight int) []utils.Point {
	positions := make([]utils.Point, 0)
	upBound := pos.Y-1 >= 0
	downBound := pos.Y+1 <= len(hm.topMap)-1
	leftBound := pos.X-1 >= 0
	rightBound := pos.X+1 <= len(hm.topMap[0])-1

	if rightBound {
		if utils.CharByteToInt(hm.topMap[pos.Y][pos.X+1]) == nextHeight {
			positions = append(positions, utils.Point{Y: pos.Y, X: pos.X + 1})
		}
	}
	if leftBound {
		if utils.CharByteToInt(hm.topMap[pos.Y][pos.X-1]) == nextHeight {
			positions = append(positions, utils.Point{Y: pos.Y, X: pos.X - 1})
		}
	}
	if upBound {
		if utils.CharByteToInt(hm.topMap[pos.Y-1][pos.X]) == nextHeight {
			positions = append(positions, utils.Point{Y: pos.Y - 1, X: pos.X})
		}
	}
	if downBound {
		if utils.CharByteToInt(hm.topMap[pos.Y+1][pos.X]) == nextHeight {
			positions = append(positions, utils.Point{Y: pos.Y + 1, X: pos.X})
		}
	}

	return positions
}

func PrintPart1() {
	input, _ := utils.ReadLines("day10/input.txt")
	hm := NewHikeMap(input)
	hm.parseMapv1()
	fmt.Println("AoC 24 Day 10, Part 1:", hm.trailCnt)
}
