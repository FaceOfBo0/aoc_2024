package day8

import (
	"AoC_24/utils"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type Antenna struct {
	xPos int
	yPos int
}

type Antinode struct {
	xPos int
	yPos int
}

type AntennaMap struct {
	data          []string
	height, width int
	freqAntMap    map[string][]Antenna
	antinodes     mapset.Set[Antinode]
}

func newAntennaMap(input []string) *AntennaMap {
	nodes := mapset.NewSet[Antinode]()
	freqAntsMap := make(map[string][]Antenna)
	for i, row := range input {
		for j := 0; j < len(row); j++ {
			if string(row[j]) != "." {
				freqAntsMap[string(row[j])] = append(freqAntsMap[string(row[j])],
					Antenna{xPos: j, yPos: i})
			}
		}
	}
	return &AntennaMap{data: input, height: len(input), width: len(input[0]),
		freqAntMap: freqAntsMap, antinodes: nodes}
}

func calcDist(a Antenna, b Antenna) (int, int) {
	return a.xPos - b.xPos, a.yPos - b.yPos
}

func (a *AntennaMap) createAntinodes() {
	for _, ants := range a.freqAntMap {
		for i := range ants {
			for j := i + 1; j < len(ants); j++ {
				deltaX, deltaY := calcDist(ants[i], ants[j])
				newXPos, newYPos := 0, 0

				if deltaX > 0 {
					deltaY *= -1
					newXPos = ants[i].xPos + deltaX
					newYPos = ants[i].yPos - deltaY

					if newXPos < a.width && newYPos >= 0 {
						a.antinodes.Add(Antinode{xPos: newXPos, yPos: newYPos})
					}

					newXPos = ants[j].xPos - deltaX
					newYPos = ants[j].yPos + deltaY
					if newXPos >= 0 && newYPos < a.height {
						a.antinodes.Add(Antinode{xPos: newXPos, yPos: newYPos})
					}
				} else {
					deltaY *= -1
					deltaX *= -1
					newXPos = ants[i].xPos - deltaX
					newYPos = ants[i].yPos - deltaY

					if newXPos >= 0 && newYPos >= 0 {
						a.antinodes.Add(Antinode{xPos: newXPos, yPos: newYPos})
					}

					newXPos = ants[j].xPos + deltaX
					newYPos = ants[j].yPos + deltaY
					if newXPos < a.width && newYPos < a.height {
						a.antinodes.Add(Antinode{xPos: newXPos, yPos: newYPos})
					}
				}
			}
		}
	}
}

func PrintPart1() {
	input, _ := utils.ReadLines("day8/input.txt")
	aMap := newAntennaMap(input)
	aMap.createAntinodes()
	fmt.Println("AoC 24 Day 8, Part 1:", aMap.antinodes.Cardinality())
}
