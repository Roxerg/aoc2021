package day1

import (
	"fmt"
	"roxerg_aoc/utils"
	"strconv"
)

func firstOne(inarr []string) {
	last_depth := 0
	count := -1
	for _, d := range inarr {
		int_depth, _ := strconv.Atoi(d)
		if int_depth > last_depth {
			count += 1
		}

		last_depth = int_depth
	}

	fmt.Println(count)
}

func secondOne(inarr []string) {

	last_depths := []int{}
	count := 0

	for _, d := range inarr[:3] {
		int_depth, _ := strconv.Atoi(d)
		last_depths = append(last_depths, int_depth)
	}

	for _, d := range inarr[3:] {
		int_depth, _ := strconv.Atoi(d)
		new_depths := last_depths
		_, new_depths = new_depths[0], append(new_depths[1:], int_depth)

		if sum(new_depths) > sum(last_depths) {
			count += 1
		}

		last_depths = new_depths
	}

	fmt.Println(count)

}

func Run() {
	_, inarr := utils.LoadFile("day1", "\n")
	firstOne(inarr)
	secondOne(inarr)
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
