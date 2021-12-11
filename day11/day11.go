package day11

import (
	"fmt"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

func firstOne(inputt [][]int) {
	input := make([][]int, len(inputt))
	for i := range inputt {
		input[i] = make([]int, len(inputt[i]))
		copy(input[i], inputt[i])
	}

	totalFlashes := 0
	// do this for 10 years

	for iteration := 0; iteration < 100; iteration++ {

		// step 1: add 1
		for y := range input {
			for x := range input[y] {
				input[y][x] += 1
			}
		}

		// octo flashes only once per step, this keeps coordinates that flashed
		detonationsThisStep := [][]int{}

		// step 2: chain reaction of DoooooooooooooooooooM

		detonatedAtLeastOnce := true
		for detonatedAtLeastOnce {
			detonatedAtLeastOnce = false
			detonationsBefore := len(detonationsThisStep)
			for y := range input {
				for x := range input[y] {
					if input[y][x] > 9 {
						previouslyDetonated := false
						for _, det := range detonationsThisStep {
							if det[0] == y && det[1] == x {
								previouslyDetonated = true
							}
						}
						if previouslyDetonated {
							continue
						}
						detonationsThisStep = append(detonationsThisStep, []int{y, x})
						input = FrigginExplode([]int{y, x}, input)
						totalFlashes += 1
					}
				}
			}
			if len(detonationsThisStep) > detonationsBefore {
				detonatedAtLeastOnce = true
			}
		}

		// step 3: reset the bangers
		for _, detonee := range detonationsThisStep {
			input[detonee[0]][detonee[1]] = 0
		}

		fmt.Println("iteration  ", iteration+1)
		fmt.Println("")
		for y := range input {
			for x := range input[y] {
				fmt.Print(input[y][x])
			}
			fmt.Println("")
		}
		fmt.Println("   ")
	}

	fmt.Println("total flashes: ", totalFlashes)

}

func secondOne(inputt [][]int) {
	input := make([][]int, len(inputt))
	for i := range inputt {
		input[i] = make([]int, len(inputt[i]))
		copy(input[i], inputt[i])
	}

	// do this for 10 years

	iteration := 0
	for true {
		iteration += 1

		// step 1: add 1
		for y := range input {
			for x := range input[y] {
				input[y][x] += 1
			}
		}

		// octo flashes only once per step, this keeps coordinates that flashed
		detonationsThisStep := [][]int{}

		// step 2: chain reaction of DoooooooooooooooooooM

		detonatedAtLeastOnce := true
		for detonatedAtLeastOnce {
			detonatedAtLeastOnce = false
			detonationsBefore := len(detonationsThisStep)
			for y := range input {
				for x := range input[y] {
					if input[y][x] > 9 {
						previouslyDetonated := false
						for _, det := range detonationsThisStep {
							if det[0] == y && det[1] == x {
								previouslyDetonated = true
							}
						}
						if previouslyDetonated {
							continue
						}
						detonationsThisStep = append(detonationsThisStep, []int{y, x})
						input = FrigginExplode([]int{y, x}, input)
					}
				}
			}
			if len(detonationsThisStep) > detonationsBefore {
				detonatedAtLeastOnce = true
			}
		}

		// step 3: reset the bangers
		for _, detonee := range detonationsThisStep {
			input[detonee[0]][detonee[1]] = 0
		}

		if len(detonationsThisStep) == 100 {
			fmt.Println("sync at:  ", iteration)
			for y := range input {
				for x := range input[y] {
					fmt.Print(input[y][x])
				}
				fmt.Println("")
			}
			return
		}

	}

}

func Run() {
	_, inarr := utils.LoadFile("day11", "\n")

	ingrid := [][]int{}

	for _, line := range inarr {
		gridline := []int{}
		for _, octopusStr := range strings.Split(line, "") {
			octopusInt, _ := strconv.Atoi(octopusStr)
			gridline = append(gridline, octopusInt)
		}
		ingrid = append(ingrid, gridline)
	}

	firstOne(ingrid)
	secondOne(ingrid)
}

func FrigginExplode(explodeePos []int, theWholeGang [][]int) [][]int {
	neigbours := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}}
	for _, coord := range neigbours {
		neigbourPos := []int{explodeePos[0] + coord[0], explodeePos[1] + coord[1]}
		if neigbourPos[0] >= 0 && neigbourPos[0] <= len(theWholeGang)-1 && neigbourPos[1] >= 0 && neigbourPos[1] <= len(theWholeGang[0])-1 {
			theWholeGang[neigbourPos[0]][neigbourPos[1]] += 1
		}
	}
	// theWholeGang[explodeePos[0]][explodeePos[1]] = 0
	return theWholeGang
}

/*
6594254334 6594254334
3856965822 3856965822
6375667284 6375667284
7252447257 7252447257
7468496589 7468496589
5278635756 5278635756
3287952832 3287952832
7993992245 7993992245
5957959665 5957959665
6394862637 6394862637



*/
