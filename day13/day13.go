package day13

import (
	"fmt"
	"math/rand"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

func firstOne(dots [][]int, folds [][]int) {

	theOneFold := []int{655, 0}
	TooLazyToThinkSoGrid := [1500][1500]int{}

	newDots := [][]int{}
	for _, dot := range dots {
		// ehh just hardcode the X axis
		if theOneFold[0] > 0 && dot[0] > theOneFold[0] {
			newDots = append(newDots, []int{theOneFold[0] - (dot[0] - theOneFold[0]), dot[1]})
			// newDots = append(newDots, []int{dot[0], dot[1] - theOneFold[1] - 1})
		} else if theOneFold[1] > 0 && dot[1] > theOneFold[1] {
			newDots = append(newDots, []int{dot[0], (theOneFold[1] - (dot[1] - theOneFold[1]))})
		} else {
			newDots = append(newDots, []int{dot[0], dot[1]})
		}
	}

	for _, dot := range newDots {
		// if TooLazyToThinkSoGrid[dot[0]][dot[1]] == 1 {
		// 	fmt.Println("overlap!", dot)
		// }
		TooLazyToThinkSoGrid[dot[0]][dot[1]] = 1
	}

	count := 0
	for x, _ := range TooLazyToThinkSoGrid {
		for y, _ := range TooLazyToThinkSoGrid[x] {
			if TooLazyToThinkSoGrid[x][y] == 1 {
				count += 1
			}
		}
	}

	fmt.Println(count)
}

func secondOne(dots [][]int, folds [][]int) {

	// theOneFold := []int{655, 0}
	TooLazyToThinkSoGrid := [1500][1500]int{}

	for _, theOneFold := range folds {
		newDots := [][]int{}
		for _, dot := range dots {
			// ehh just hardcode the X axis
			if theOneFold[0] > 0 && dot[0] > theOneFold[0] {
				newDots = append(newDots, []int{theOneFold[0] - (dot[0] - theOneFold[0]), dot[1]})
				// newDots = append(newDots, []int{dot[0], dot[1] - theOneFold[1] - 1})
			} else if theOneFold[1] > 0 && dot[1] > theOneFold[1] {
				newDots = append(newDots, []int{dot[0], (theOneFold[1] - (dot[1] - theOneFold[1]))})
			} else {
				newDots = append(newDots, []int{dot[0], dot[1]})
			}
		}
		dots = [][]int{}
		for _, dot := range newDots {
			dots = append(dots, []int{dot[0], dot[1]})
		}
	}

	for _, dot := range dots {
		// if TooLazyToThinkSoGrid[dot[1]][dot[0]] == 1 {
		// 	// fmt.Println("overlap!", dot)
		// }
		TooLazyToThinkSoGrid[dot[1]][dot[0]] = 1
	}

	for idx, gridrow := range TooLazyToThinkSoGrid {
		for _, dot := range gridrow[:50] {
			if dot == 1 {
				emoji := []string{"🎄", "🎁", "✨"}
				n := rand.Int() % len(emoji)
				fmt.Print(emoji[n])
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
		if idx > 7 {
			break //REUPUPKR
		}
	}
}

func Run() {
	_, inarr := utils.LoadFile("day13", "\n")

	dots := [][]int{}
	folds := [][]int{}

	var separatorIdx int
	for idx, line := range inarr {
		if len(line) < 3 {
			separatorIdx = idx
			break
		}
		coordsStr := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordsStr[0])
		y, _ := strconv.Atoi(coordsStr[1])
		dots = append(dots, []int{x, y})
	}

	for _, line := range inarr[separatorIdx+1:] {

		command := strings.Split(line, "g")
		axisAndNumber := strings.Split(command[1], "=")
		axis := axisAndNumber[0]
		number, _ := strconv.Atoi(axisAndNumber[1])
		if strings.TrimSpace(axis) == "x" {
			folds = append(folds, []int{number, 0})
		} else {
			folds = append(folds, []int{0, number})
		}

	}

	firstOne(dots, folds)
	secondOne(dots, folds)
}
