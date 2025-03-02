package day13

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Machine struct {
	ax, ay int // movement from button A
	bx, by int // movement from button B
	px, py int // prize position
}

func PrintPart1Alt() {
	file, _ := os.Open("day13/input.txt")
	scanner := bufio.NewScanner(file)
	var machines []Machine
	var lines []string

	// Read all non-empty lines.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	// Each machine has 3 lines: Button A, Button B, Prize.
	if len(lines)%3 != 0 {
		fmt.Fprintln(os.Stderr, "Input does not have groups of 3 lines")
		os.Exit(1)
	}

	// Process each group of 3 lines.
	for i := 0; i < len(lines); i += 3 {
		lineA := lines[i]
		lineB := lines[i+1]
		lineP := lines[i+2]

		var ax, ay, bx, by, px, py int
		// Parse Button A line, e.g. "Button A: X+94, Y+34"
		_, err := fmt.Sscanf(lineA, "Button A: X+%d, Y+%d", &ax, &ay)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing line:", lineA, err)
			os.Exit(1)
		}
		// Parse Button B line, e.g. "Button B: X+22, Y+67"
		_, err = fmt.Sscanf(lineB, "Button B: X+%d, Y+%d", &bx, &by)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing line:", lineB, err)
			os.Exit(1)
		}
		// Parse Prize line, e.g. "Prize: X=8400, Y=5400"
		_, err = fmt.Sscanf(lineP, "Prize: X=%d, Y=%d", &px, &py)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing line:", lineP, err)
			os.Exit(1)
		}

		machines = append(machines, Machine{ax, ay, bx, by, px, py})
	}

	totalCost := 0
	// For each machine, try all combinations of button presses (0 to 100 each).
	for _, m := range machines {
		bestCost := -1
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				x := m.ax*a + m.bx*b
				y := m.ay*a + m.by*b
				if x == m.px && y == m.py {
					fmt.Printf("a: %v\n", a)
					fmt.Printf("b: %v\n", b)
					cost := 3*a + b
					if bestCost == -1 || cost < bestCost {
						bestCost = cost
					}
				}
			}
		}
		// Only add if a valid combination was found.
		if bestCost != -1 {
			totalCost += bestCost
		}
	}

	fmt.Println(totalCost)
}
