package day10

import (
	"AoC_24/utils"
	"fmt"
	"strconv"

	mapset "github.com/deckarep/golang-set/v2"
)

type Point struct {
	y int
	x int
}

type HikeMap struct {
	topMap       []string
	width        int
	trailCnt     int
	trailCntDist int
	trailEnds    mapset.Set[Point]
}

func NewHikeMap(input []string) HikeMap {
	return HikeMap{
		topMap:       input,
		width:        len(input[0]),
		trailCnt:     0,
		trailCntDist: 0,
		trailEnds:    mapset.NewSet[Point](),
	}
}

func (hm *HikeMap) parseMapv1() {
	for y, row := range hm.topMap {
		for x, col := range row {
			if elem, _ := strconv.Atoi(string(col)); elem == 0 {
				hm.searchTrailsv1(Point{y: y, x: x}, 0)
				hm.trailCnt += hm.trailEnds.Cardinality()
				hm.trailEnds.Clear()
			}
		}
	}
}

func (hm HikeMap) searchTrailsv1(pos Point, height int) {
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

func (hm HikeMap) findNextHeights(pos Point, nextHeight int) []Point {
	positions := make([]Point, 0)
	upBound := pos.y-1 >= 0
	downBound := pos.y+1 <= len(hm.topMap)-1
	leftBound := pos.x-1 >= 0
	rightBound := pos.x+1 <= len(hm.topMap[0])-1

	if rightBound {
		if utils.CharByteToInt(hm.topMap[pos.y][pos.x+1]) == nextHeight {
			positions = append(positions, Point{y: pos.y, x: pos.x + 1})
		}
	}
	if leftBound {
		if utils.CharByteToInt(hm.topMap[pos.y][pos.x-1]) == nextHeight {
			positions = append(positions, Point{y: pos.y, x: pos.x - 1})
		}
	}
	if upBound {
		if utils.CharByteToInt(hm.topMap[pos.y-1][pos.x]) == nextHeight {
			positions = append(positions, Point{y: pos.y - 1, x: pos.x})
		}
	}
	if downBound {
		if utils.CharByteToInt(hm.topMap[pos.y+1][pos.x]) == nextHeight {
			positions = append(positions, Point{y: pos.y + 1, x: pos.x})
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
