package day9

import (
	"AoC_24/utils"
	"fmt"
	"slices"
	"strconv"
)

type freeSpace struct {
	count int
	start int
	end   int
}

type MemoryBlock struct {
	data   []int
	files  []int
	spaces []freeSpace
}

func parseBlock(input string) *MemoryBlock {
	rawData := make([]int, 0)
	for _, ch := range input {
		value, _ := strconv.Atoi(string(ch))
		rawData = append(rawData, value)
	}

	f := make([]int, 0)
	id := 0
	for i, num := range rawData {
		if i%2 == 0 {
			k := 0
			for k < num {
				f = append(f, id)
				k++
			}
			id++
		} else {
			k := 0
			for k < num {
				f = append(f, -1)
				k++
			}
		}
	}

	return &MemoryBlock{data: rawData, files: f, spaces: make([]freeSpace, 0)}
}

func (mb *MemoryBlock) moveBlock() {
	endIdx := len(mb.files) - 1
	freeIdx := slices.Index(mb.files, -1)

	for freeIdx <= endIdx {
		if mb.files[endIdx] != -1 {
			mb.files[freeIdx] = mb.files[endIdx]
		}
		endIdx--
		mb.files = mb.files[:len(mb.files)-1]
		freeIdx = slices.Index(mb.files, -1)
		if freeIdx == -1 {
			break
		}
	}
}

func (mb *MemoryBlock) calcCkecksum() int {
	checksum := 0
	for i, v := range mb.files {
		if v != -1 {
			checksum += i * v
		}

	}
	return checksum
}

func PrintPart1() {
	input, _ := utils.ReadLines("day9/input.txt")
	mb := parseBlock(input[0])

	mb.moveBlock()
	fmt.Println("AoC 24 Day 9, Part 1:", mb.calcCkecksum())
}
