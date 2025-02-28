package day12

<<<<<<< HEAD
=======
// Costs:
// test 1.0 = 1206
// test 1.1 = 80
// test 1.2 = 436
// test 2.0 = 236
// test 2.1 = 368

>>>>>>> b78744ea6719ea1c9eb448e13eb0a21dfa2a9581
import (
	"AoC_24/utils"
	"fmt"
)

func PrintPart2() {
<<<<<<< HEAD
	input, _ := utils.ReadLines("day12/input.txt")
	result := parseRegions(input)
	fmt.Println("AoC 24 Day 12, Part 2:", result)
=======
	input, _ := utils.ReadLines("day12/test1.0.txt")
	result := parseRegions(input)
	fmt.Println("AoC 24 Day 12, Part 2:", calcCost(result))
>>>>>>> b78744ea6719ea1c9eb448e13eb0a21dfa2a9581
}
