package day9

import (
	"AoC_24/utils"
	"fmt"
)

func (m *MemoryBlock) createFreeSpaceMap() map[int]freeSpace {

	spaceMap := make(map[int]freeSpace)
	count := 0
	start := -1

	// Iterate through files to find consecutive -1s

	for i, val := range m.files {
		if val == -1 {
			if start == -1 {
				start = i // Mark the start of a new free space block
			}
			count++
		} else if start != -1 {
			// We've reached the end of a free space block
			spaceMap[start] = freeSpace{
				count: count,
				start: start,
				end:   start + count - 1,
			}
			// Reset counters
			count = 0
			start = -1
		}
	}

	// Handle case where free space extends to the end of the slice
	if start != -1 {
		spaceMap[start] = freeSpace{
			count: count,
			start: start,
			end:   start + count - 1,
		}
	}
	return spaceMap
}

func (mb *MemoryBlock) moveBlockP2() {

	// Process files from highest ID to lowest
	for id := mb.files[len(mb.files)-1]; id >= 0; id-- {
		// Find all positions of current file ID
		filePositions := make([]int, 0)
		for i, v := range mb.files {
			if v == id {
				filePositions = append(filePositions, i)
			}
		}

		if len(filePositions) == 0 {
			continue
		}

		// Find the leftmost suitable free space
		fileSize := len(filePositions)
		freeSpaces := mb.createFreeSpaceMap()

		// Find the leftmost free space that can fit the file
		bestStart := len(mb.files)

		for start, space := range freeSpaces {
			if space.count >= fileSize && start < bestStart {
				bestStart = start
			}
		}

		// If we found a suitable space and it's to the left of the file
		if bestStart < filePositions[0] {
			// Move the file to the new position
			for i := 0; i < fileSize; i++ {
				// Mark old position as free
				mb.files[filePositions[i]] = -1
				// Place file in new position
				mb.files[bestStart+i] = id
			}
		}
	}
}

func PrintPart2() {
	input, _ := utils.ReadLines("day9/input.txt")
	mb := parseBlock(input[0])
	mb.moveBlockP2()
	fmt.Println("AoC 24 Day 9, Part 2:", mb.calcCkecksum())
}
