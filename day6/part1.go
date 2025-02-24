package day6

import (
	"AoC_24/utils"
	"fmt"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Position struct {
	y int
	x int
}

type Maze struct {
	data      []string
	exits     bool
	height    int
	width     int
	guardPos  Position
	guardDir  Direction
	guardPath []Position
}

func newMaze(input []string) *Maze {
	gPos := Position{}
	for i, v := range input {
		if strings.Contains(v, "^") {
			gPos.y = i
			gPos.x = strings.Index(v, "^")
		}
	}

	return &Maze{height: len(input), width: len(input), guardPath: make([]Position, 0),
		data: input, guardPos: gPos, guardDir: Up, exits: false}
}

func (m *Maze) calcPatrolCnt() int {
	routeCnt := 0
	for _, elem := range m.data {
		routeCnt += strings.Count(elem, "X")
	}
	return routeCnt
}

func (m *Maze) traverse() {
	m.guardPath = append(m.guardPath, m.guardPos)

	for !m.exits {
		m.data[m.guardPos.y] = utils.ReplaceAtIndex(m.data[m.guardPos.y], m.guardPos.x, "X")
		switch m.guardDir {
		case Up:
			if m.guardPos.y-1 < 0 {
				m.exits = true
				break
			}
			if string(m.data[m.guardPos.y-1][m.guardPos.x]) == "#" {
				m.guardDir = Right
				break
			}
			m.guardPos.y -= 1
			m.guardPath = append(m.guardPath, m.guardPos)

		case Down:
			if m.guardPos.y+1 >= m.height {
				m.exits = true
				break
			}
			if string(m.data[m.guardPos.y+1][m.guardPos.x]) == "#" {
				m.guardDir = Left
				break
			}
			m.guardPos.y += 1
			m.guardPath = append(m.guardPath, m.guardPos)

		case Left:
			if m.guardPos.x-1 < 0 {
				m.exits = true
				break
			}
			if string(m.data[m.guardPos.y][m.guardPos.x-1]) == "#" {
				m.guardDir = Up
				break
			}
			m.guardPos.x -= 1
			m.guardPath = append(m.guardPath, m.guardPos)

		case Right:
			if m.guardPos.x+1 >= m.width {
				m.exits = true
				break
			}
			if string(m.data[m.guardPos.y][m.guardPos.x+1]) == "#" {
				m.guardDir = Down
				break
			}
			m.guardPos.x += 1
			m.guardPath = append(m.guardPath, m.guardPos)
		}
	}
}

func (m *Maze) print() {
	for i, v := range m.data {
		fmt.Printf("row %v: %v\n", i, v)
	}
}

func PrintPart1() {
	input, _ := utils.ReadLines("day6/test.txt")
	maze := newMaze(input)
	maze.traverse()
	fmt.Println("AoC 24 Day 6, Part 1:", maze.calcPatrolCnt())
}
