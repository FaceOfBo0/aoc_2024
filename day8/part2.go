package day8

import (
	"AoC_24/utils"
	"fmt"
)

func (a *AntennaMap) createAntinodesP2() {
	for _, ants := range a.freqAntMap {
		for i := range ants {
			for j := i + 1; j < len(ants); j++ {
				deltaX, deltaY := calcDist(ants[i], ants[j])
				newXPos, newYPos := 0, 0

				if deltaX > 0 {
					deltaY *= -1
					a.antinodes.Add(Antinode{
						xPos: ants[i].xPos,
						yPos: ants[i].yPos,
					})
					newXPos = ants[i].xPos + deltaX
					newYPos = ants[i].yPos - deltaY

					for newXPos < a.width && newYPos >= 0 {
						a.antinodes.Add(Antinode{
							xPos: newXPos,
							yPos: newYPos,
						})
						a.data[newYPos] = utils.ReplaceAtIndex(
							a.data[newYPos], newXPos, "#")
						newXPos += deltaX
						newYPos -= deltaY
					}

					a.antinodes.Add(Antinode{
						xPos: ants[j].xPos,
						yPos: ants[j].yPos,
					})
					newXPos = ants[j].xPos - deltaX
					newYPos = ants[j].yPos + deltaY
					for newXPos >= 0 && newYPos < a.height {
						a.antinodes.Add(Antinode{
							xPos: newXPos,
							yPos: newYPos,
						})
						a.data[newYPos] = utils.ReplaceAtIndex(
							a.data[newYPos], newXPos, "#")
						newXPos -= deltaX
						newYPos += deltaY
					}
				} else {
					deltaY *= -1
					deltaX *= -1
					a.antinodes.Add(Antinode{
						xPos: ants[i].xPos,
						yPos: ants[i].yPos,
					})
					newXPos = ants[i].xPos - deltaX
					newYPos = ants[i].yPos - deltaY

					for newXPos >= 0 && newYPos >= 0 {
						a.antinodes.Add(Antinode{
							xPos: newXPos,
							yPos: newYPos,
						})
						a.data[newYPos] = utils.ReplaceAtIndex(
							a.data[newYPos], newXPos, "#")
						newXPos -= deltaX
						newYPos -= deltaY
					}

					a.antinodes.Add(Antinode{
						xPos: ants[j].xPos,
						yPos: ants[j].yPos,
					})
					newXPos = ants[j].xPos + deltaX
					newYPos = ants[j].yPos + deltaY
					for newXPos < a.width && newYPos < a.height {
						a.antinodes.Add(Antinode{
							xPos: newXPos,
							yPos: newYPos,
						})
						a.data[newYPos] = utils.ReplaceAtIndex(
							a.data[newYPos], newXPos, "#")
						newXPos += deltaX
						newYPos += deltaY
					}
				}
			}
		}
	}
}

func PrintPart2() {
	input, _ := utils.ReadLines("day8/input.txt")
	antMap := newAntennaMap(input)
	antMap.createAntinodesP2()
	fmt.Println("AoC 24 Day 8, Part 2:", antMap.antinodes.Cardinality())
}
