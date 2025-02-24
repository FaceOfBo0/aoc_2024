package utils

import (
	"bufio"
	"cmp"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinates interface {
	GetX() int
	GetY() int
}

func CharByteToInt(input byte) int {
	result, _ := strconv.Atoi(string(input))
	return result
}

func ReplaceCoords(grid []string, pos Coordinates, repl string) {
	grid[pos.GetY()] = ReplaceAtIndex(grid[pos.GetY()], pos.GetX(), repl)
}

func ReplaceAtIndex(input string, index int, replacement string) string {
	return strings.Join([]string{input[:index], replacement, input[index+1:]}, "")
}

func SplitListOnce[I comparable](input []I, sep I) ([]I, []I) {
	splitIdx := slices.Index(input, sep)
	return input[:splitIdx], input[splitIdx+1:]
}

func Foldl[I any](list []I, fn func(I, I) I, init I) I {
	result := init
	for _, x := range list {
		result = fn(result, x)
	}
	return result
}

func FoldlStr(list []string) string {
	result := ""
	for _, x := range list {
		result += x
	}
	return result
}

func ReadLines(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return []string{}, err
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func SortKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keysList := make([]K, 0, len(m))
	for k := range m {
		keysList = append(keysList, k)
	}
	slices.SortFunc(keysList, func(a, b K) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	})
	return keysList
}

func MapList[I, O any](list []I, fn func(I) O) []O {
	outList := make([]O, len(list))
	for i, elem := range list {
		outList[i] = fn(elem)
	}
	return outList
}

func RemoveOne[S any](list []S, idx int) []S {
	newSlice := make([]S, 0, len(list)-1)
	for i, v := range list {
		if i != idx {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func RemoveOneMut(slice []int, idx int) []int {
	return slices.Delete(slice, idx, idx+1)
}
