package day10

import (
	"fmt"
	"roxerg_aoc/utils"
	"sort"
	"strings"
)

func firstOne(inarr []string) {

	sum := 0
	for _, line := range inarr {
		expectedEndingStack := []string{}
		chars := strings.Split(line, "")

		for _, char := range chars {
			if char == "(" || char == "[" || char == "{" || char == "<" {
				if char == "(" {
					expectedEndingStack = Push(expectedEndingStack, ")")
				} else if char == "[" {
					expectedEndingStack = Push(expectedEndingStack, "]")
				} else if char == "{" {
					expectedEndingStack = Push(expectedEndingStack, "}")
				} else if char == "<" {
					expectedEndingStack = Push(expectedEndingStack, ">")
				}

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

	scores := []int{}
	for _, line := range inarr {
		expectedEndingStack := []string{}
		isCorrupted := false
		chars := strings.Split(line, "")

		for _, char := range chars {
			if char == "(" || char == "[" || char == "{" || char == "<" {
				if char == "(" {
					expectedEndingStack = Push(expectedEndingStack, ")")
				} else if char == "[" {
					expectedEndingStack = Push(expectedEndingStack, "]")
				} else if char == "{" {
					expectedEndingStack = Push(expectedEndingStack, "}")
				} else if char == "<" {
					expectedEndingStack = Push(expectedEndingStack, ">")
				}

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
			fmt.Println("-----")
			for idx := len(expectedEndingStack) - 1; idx >= 0; idx-- {
				end := expectedEndingStack[idx]
				fmt.Println(end, linescore)
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
			fmt.Println(strings.Join(expectedEndingStack, ""), " score ", linescore)
			scores = append(scores, linescore)
		}

	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])

}

func Run() {
	_, inarr := utils.LoadFile("day10", "\n")
	// inarr := strings.Split(instr, ",")
	// inarrInt := []int{}
	// for i := range inarr {
	// 	num, _ := strconv.Atoi(strings.TrimSpace(inarr[i]))
	// 	inarrInt = append(inarrInt, num)
	// }

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
