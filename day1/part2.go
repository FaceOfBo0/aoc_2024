package day1

func countFreqs(list []int) map[int]int {
	freqs := make(map[int]int)

	for _, elem := range list {
		freqs[elem]++
	}
	return freqs
}

func calcTotalSim() int {
	numsList := ReadInNums()
	rightfreq := countFreqs(numsList.rightNums)

	totalSimlrty := 0
	for _, elem := range numsList.leftNums {
		totalSimlrty += elem * rightfreq[elem]
	}
	return totalSimlrty
}

func PrintPart2() {
	println("AoC 24 Day 1, Part 2:", calcTotalSim())
}
