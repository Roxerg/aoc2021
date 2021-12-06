package day5

import (
	"fmt"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

func firstOne(inarr []string) {
	var all_coords = [][]int{}
	for line_idx := range inarr {
		var coords = []int{}
		pair := strings.Split(inarr[line_idx], " -> ")
		for pair_idx := range pair {
			coord := strings.Split(pair[pair_idx], ",")
			x, _ := strconv.Atoi(strings.TrimSpace(coord[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(coord[1]))
			coords = append(coords, x)
			coords = append(coords, y)
		}
		all_coords = append(all_coords, coords)
	}

	var points = [1000][1000]int{}
	for i := range points {
		for j := range points[i] {
			points[i][j] = 0
		}
	}

	for line := range all_coords {
		x1, y1 := all_coords[line][0], all_coords[line][1]
		x2, y2 := all_coords[line][2], all_coords[line][3]
		if x1 == x2 || y1 == y2 {
			// fmt.Println(all_coords[line])
			x, y := x1, y1

			points[x][y] += 1
			for x != x2 || y != y2 {
				if x != x2 {
					if x < x2 {
						x += 1
					} else {
						x -= 1
					}
				}
				if y != y2 {
					if y < y2 {
						y += 1
					} else {
						y -= 1
					}
				}
				points[x][y] += 1
			}

		}
	}

	counter := 0

	for i := range points {
		for j := range points[i] {
			if points[i][j] >= 2 {
				counter += 1
			}
		}
	}

	fmt.Println(counter)

}

func secondOne(inarr []string) {
	var all_coords = [][]int{}
	for line_idx := range inarr {
		var coords = []int{}
		pair := strings.Split(inarr[line_idx], " -> ")
		for pair_idx := range pair {
			coord := strings.Split(pair[pair_idx], ",")
			x, _ := strconv.Atoi(strings.TrimSpace(coord[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(coord[1]))
			coords = append(coords, x)
			coords = append(coords, y)
		}
		all_coords = append(all_coords, coords)
	}

	var points = [1000][1000]int{}
	for i := range points {
		for j := range points[i] {
			points[i][j] = 0
		}
	}

	for line := range all_coords {
		x1, y1 := all_coords[line][0], all_coords[line][1]
		x2, y2 := all_coords[line][2], all_coords[line][3]
		if true {
			// fmt.Println(all_coords[line])
			x, y := x1, y1

			points[x][y] += 1
			for x != x2 || y != y2 {
				if x != x2 {
					if x < x2 {
						x += 1
					} else {
						x -= 1
					}
				}
				if y != y2 {
					if y < y2 {
						y += 1
					} else {
						y -= 1
					}
				}
				points[x][y] += 1
			}

		}
	}

	counter := 0

	for i := range points {
		for j := range points[i] {
			if points[i][j] >= 2 {
				counter += 1
			}
		}
	}

	fmt.Println(counter)

}

func Run() {
	_, inarr := utils.LoadFile("day5", "\n")
	firstOne(inarr)
	secondOne(inarr)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
