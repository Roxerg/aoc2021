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
					switch char {
					case ")":
						sum += 3
					case "]":
						sum += 57
					case "}":
						sum += 1197
					case ">":
						sum += 25137
					}
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
				linescore = linescore * 5
				switch end {
				case ")":
					linescore += 1
				case "]":
					linescore += 2
				case "}":
					linescore += 3
				case ">":
					linescore += 4
				}
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
