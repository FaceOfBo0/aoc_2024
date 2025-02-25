package day10

import (
	"AoC_24/utils"
	"fmt"
	"strconv"
)

func (hm *HikeMap) parseMapv2() {
	for y, row := range hm.topMap {
		for x, col := range row {
			if elem, _ := strconv.Atoi(string(col)); elem == 0 {
				hm.searchTrailsv2(utils.Point{Y: y, X: x}, 0)
			}
		}
	}
}

func (hm *HikeMap) searchTrailsv2(pos utils.Point, height int) {
	nextHts := hm.findNextHeights(pos, height+1)
	if height+1 == 9 {
		hm.trailCntDist += len(nextHts)
		return
	}
	for _, point := range nextHts {
		hm.searchTrailsv2(point, height+1)
	}
}

func PrintPart2() {
	input, _ := utils.ReadLines("day10/input.txt")
	hm := NewHikeMap(input)
	hm.parseMapv2()
	fmt.Println("AoC 24 Day 10, Part 2:", hm.trailCntDist)
}
