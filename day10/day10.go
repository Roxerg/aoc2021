package day10

import (
	"fmt"
	"roxerg_aoc/utils"
	"sort"
	"strings"
)

func firstOne(inarr []string) {

	bracketPairs := make(map[string]string)

	bracketPairs["("] = ")"
	bracketPairs["{"] = "}"
	bracketPairs["["] = "]"
	bracketPairs["<"] = ">"
	charValues := make(map[string]int)
	charValues[")"] = 3
	charValues["]"] = 57
	charValues["}"] = 1197
	charValues[">"] = 25137

	sum := 0
	for _, line := range inarr {
		expectedEndingStack := []string{}

		for _, char := range strings.Split(line, "") {
			bracketEnd, exists := bracketPairs[char]
			if exists {
				expectedEndingStack = Push(expectedEndingStack, bracketEnd)

			} else {
				ending := expectedEndingStack[len(expectedEndingStack)-1]
				if char != ending {
					sum += charValues[char]
					break
				} else if char == ending {
					expectedEndingStack, _ = Pop(expectedEndingStack)
				}

			}
		}

	}

	fmt.Println(sum)

}

func secondOne(inarr []string) {

	bracketPairs := make(map[string]string)
	bracketPairs["("] = ")"
	bracketPairs["{"] = "}"
	bracketPairs["["] = "]"
	bracketPairs["<"] = ">"
	charValues := make(map[string]int)
	charValues[")"] = 1
	charValues["]"] = 2
	charValues["}"] = 3
	charValues[">"] = 4

	scores := []int{}

	for _, line := range inarr {
		expectedEndingStack := []string{}
		isCorrupted := false

		for _, char := range strings.Split(line, "") {
			bracketEnd, exists := bracketPairs[char]
			if exists {
				expectedEndingStack = Push(expectedEndingStack, bracketEnd)

			} else {
				ending := expectedEndingStack[len(expectedEndingStack)-1]
				if char != ending {
					isCorrupted = true
					break
				} else if char == ending {
					expectedEndingStack, ending = Pop(expectedEndingStack)
				}

			}
		}

		linescore := 0
		if !isCorrupted {
			for idx := len(expectedEndingStack) - 1; idx >= 0; idx-- {
				end := expectedEndingStack[idx]
				linescore = linescore*5 + charValues[end]
			}
			// fmt.Println(strings.Join(expectedEndingStack, ""), " score ", linescore)
			scores = append(scores, linescore)
		}

	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])

}

func Run() {
	_, inarr := utils.LoadFile("day10", "\n")

	firstOne(inarr)
	secondOne(inarr)
}

func Push(s []string, i string) []string {
	return append(s, i)
}

func Pop(s []string) ([]string, string) {
	l := len(s)
	if l == 0 {
		return []string{}, ""
	}
	return s[:l-1], s[l-1]
}
