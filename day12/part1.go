package day12

import (
	"AoC_24/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

type Patch struct {
	plantType  string
	areas      []int
	perimeters []int
	overlaps   []int
	borderSets []mapset.Set[Coordinate]
}

type Coordinate struct {
	idx int
	pnt utils.Point
}

func NewPatch(plantType string) *Patch {
	return &Patch{
		plantType:  plantType,
		areas:      make([]int, 0),
		perimeters: make([]int, 0),
		overlaps:   make([]int, 0),
		borderSets: make([]mapset.Set[Coordinate], 0),
	}
}

/*func (p *Patch) updateBordersv2() {
	thisY, nextY := p.points[len(p.points)-2].Y, p.points[len(p.points)-1].Y
	thisX, nextX := p.points[len(p.points)-2].X, p.points[len(p.points)-1].X

	if nextY == thisY && nextX == thisX+1 && !slices.Contains(p.points, utils.Point{X: nextX, Y: nextY - 1}) {
		p.borderMap["N"]++
		p.borderMap["S"]++
		p.perimeter += 2

	} else if nextY == thisY+1 {
		if slices.Contains(p.points, utils.Point{X: nextX, Y: nextY - 1}) {
			p.borderMap["W"]++
			p.borderMap["E"]++
			p.perimeter += 2

		} else {
			p.borderMap["N"]++
			p.borderMap["W"]++
			p.borderMap["E"]++
			p.perimeter += 3

		}
	} else if !slices.Contains(p.points, utils.Point{Y: nextY, X: nextX - 1}) &&
		!slices.Contains(p.points, utils.Point{Y: nextY - 1, X: nextX}) {
		p.borderMap["N"]++
		p.borderMap["W"]++
		p.borderMap["E"]++
		p.borderMap["S"]++
		p.perimeter += 4
	}
}*/

/*func (p *Patch) updateBorders() {
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
*/

/*func (p *Patch) calcBorders() {
	tempList := make([]utils.Point, 0)
	for i, pnt := range p.points {
		if !slices.Contains(tempList, pnt) {
			tempList = append(tempList, pnt)
		}
		for j := i + 1; j < len(p.points); j++ {
			if pnt.X == p.points[j].X || pnt.Y == p.points[j].Y {
				if !slices.Contains(tempList, p.points[j]) {
					tempList = append(tempList, p.points[j])
				}
			}
		}

	}
}*/

func calcCost(garden map[string]*Patch) int {
	totalCost := 0
	for _, patch := range garden {
		for i, area := range patch.areas {
			totalCost += area * (area*4 - (patch.overlaps[i] * 2))
		}
	}
	return totalCost
}

func parsePatches(input []string) map[string]*Patch {
	regionMap := make(map[string]*Patch)
	for y, row := range input {
		for x, elem := range row {
			char := string(elem)
			if key, ok := regionMap[char]; ok {
				isIn := false
				for i, elem := range key.borderSets {
					thisPnt := utils.Point{Y: y, X: x}
					upCoord := Coordinate{idx: i, pnt: utils.Point{Y: y - 1, X: x}}
					backCoord := Coordinate{idx: i, pnt: utils.Point{Y: y, X: x - 1}}
					forwCoord := Coordinate{idx: i, pnt: utils.Point{Y: y, X: x + 1}}

					if elem.Contains(upCoord) || elem.Contains(backCoord) ||
						elem.Contains(forwCoord) {
						elem.Add(Coordinate{idx: i, pnt: thisPnt})
						key.areas[i]++
						// update overlap
						if x <= len(row)-2 {
							if string(row[x+1]) == char {
								elem.Add(forwCoord)
								key.overlaps[i]++
							}
						}
						if y <= len(input)-2 {
							if string(input[y+1][x]) == char {
								elem.Add(Coordinate{idx: i, pnt: utils.Point{Y: y + 1, X: x}})
								if tempX := x - 1; tempX > 0 {
									for string(input[y+1][tempX]) == char {
										elem.Add(Coordinate{idx: i, pnt: utils.Point{Y: y + 1, X: tempX}})
										tempX--
										if tempX < 0 {
											break
										}
									}
								}
								key.overlaps[i]++
							}
						}
						isIn = true
					}
				}
				if !isIn {
					key.borderSets = append(key.borderSets, mapset.NewSet[Coordinate]())
					newIdx := len(key.borderSets) - 1
					key.borderSets[newIdx].Add(Coordinate{idx: newIdx, pnt: utils.Point{Y: y, X: x}})
					key.areas = append(key.areas, 1)
					key.overlaps = append(key.overlaps, 0)
					if x <= len(row)-2 {
						if string(row[x+1]) == char {
							key.borderSets[newIdx].Add(Coordinate{idx: 0, pnt: utils.Point{Y: y, X: x + 1}})
							key.overlaps[newIdx]++
						}
					}
					if y <= len(input)-2 {
						if string(input[y+1][x]) == char {
							key.borderSets[newIdx].Add(Coordinate{idx: 0, pnt: utils.Point{Y: y + 1, X: x}})
							key.overlaps[newIdx]++
						}
					}

				}
				//key.points = append(key.points, utils.Point{Y: y, X: x})
				//key.areas++

			} else {
				regionMap[char] = NewPatch(char)
				p := regionMap[char]
				p.borderSets = append(p.borderSets, mapset.NewSet[Coordinate]())
				p.borderSets[0].Add(Coordinate{idx: 0, pnt: utils.Point{Y: y, X: x}})
				p.areas = append(p.areas, 1)
				p.overlaps = append(p.overlaps, 0)
				// update overlap
				if x <= len(row)-2 {
					if string(row[x+1]) == char {
						p.borderSets[0].Add(Coordinate{idx: 0, pnt: utils.Point{Y: y, X: x + 1}})
						p.overlaps[0]++
					}
				}
				if y <= len(input)-2 {
					if string(input[y+1][x]) == char {
						p.borderSets[0].Add(Coordinate{idx: 0, pnt: utils.Point{Y: y + 1, X: x}})
						p.overlaps[0]++
					}
				}
				//p.points = append(p.points, utils.Point{Y: y, X: x})
				//p.areas++
			}
		}
	}
	return regionMap
}

func PrintPart1() {
	input, _ := utils.ReadLines("day12/test2.txt")
	result := parsePatches(input)
	for k, p := range result {
		fmt.Printf("%v: %v\n", k, p)
	}
	fmt.Println("AoC 24 Day 12, Part 1:", calcCost(result))
}
