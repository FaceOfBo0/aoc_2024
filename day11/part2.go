package day11

import (
	"AoC_24/utils"
	"fmt"
	"math"
	"sync"
)

func splitNumv2(num uint64, base int) (uint64, uint64) {
	exp := math.Pow10(base / 2)
	div := float64(num) / exp
	left, right := math.Modf(div)
	return uint64(left), uint64(right * exp)
}

func stoneNumsRec(num uint64, cnt int) uint64 {
	if cnt == 0 {
		return 1
	}

	if num == 0 {
		return stoneNumsRec(1, cnt-1)
	} else {
		baseRes, _ := math.Modf(math.Log10(float64(num)))
		if int(baseRes+1)%2 == 0 {
			//lft, rht := splitNumv2(num, int(baseRes+1))
			lft, rht := splitNum(num)
			return stoneNumsRec(lft, cnt-1) + stoneNumsRec(rht, cnt-1)
		} else {
			return stoneNumsRec(num*2024, cnt-1)
		}
	}
}

type cacheKey struct {
	num uint64
	cnt int
}

var cache = make(map[cacheKey]uint64)

func stoneNumsRecCached2(num uint64, cnt int) uint64 {
	key := cacheKey{num, cnt}

	// Try to read from cache first
	if val, ok := cache[key]; ok {
		return val
	}

	var result uint64
	if cnt == 0 {
		result = 1
	} else if num == 0 {
		result = stoneNumsRecCached2(1, cnt-1)
	} else {
		baseRes, _ := math.Modf(math.Log10(float64(num)))
		if int(baseRes+1)%2 == 0 {
			lft, rht := splitNum(num)
			result = stoneNumsRecCached2(lft, cnt-1) + stoneNumsRecCached2(rht, cnt-1)
		} else {
			result = stoneNumsRecCached2(num*2024, cnt-1)
		}
	}

	cache[key] = result
	return result
}

func blinkv2seq(stones []uint64, blinkCnt int) uint64 {
	var totalCnt uint64
	totalCnt = 0
	for _, st := range stones {
		totalCnt += stoneNumsRecCached2(st, blinkCnt)
	}
	return totalCnt
}

func blinkv2(stones []uint64, blinkCnt int) uint64 {
	var totalCnt uint64
	totalCnt = 0
	var wg sync.WaitGroup
	totalCntChan := make(chan uint64)

	for _, st := range stones {
		wg.Add(1)
		go func(stone uint64) {
			defer wg.Done()
			totalCntChan <- stoneNumsRecCached2(stone, blinkCnt)
		}(st)
	}

	go func() {
		wg.Wait()
		close(totalCntChan)
	}()

	for elem := range totalCntChan {
		totalCnt += elem
	}
	return totalCnt
}

func PrintPart2() {
	input, _ := utils.ReadLines("day11/input.txt")
	nums := parseNums(input[0])
	fmt.Println("AoC 24 Day 11, Part 2:", blinkv2seq(nums, 25))
}
