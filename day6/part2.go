package day6

import "AoC_24/utils"

func checkLoop(m *Maze) {
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
		}
	}
}

func (m *Maze) traverseP2() {
	for !m.exits {
		m.data[m.guardPos.y] = utils.ReplaceAtIndex(m.data[m.guardPos.y], m.guardPos.x, "X")

		tempData := make([]string, 0, len(m.data))
		copy(tempData, m.data)

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
		}
	}
}

func PrintPart2() {

}
