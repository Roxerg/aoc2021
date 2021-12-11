package day09

import (
	"fmt"
	"roxerg_aoc/utils"
	"sort"
	"strconv"
	"strings"
)

func firstOne(inarr []string) {
	grid := [][]int{}

	for _, line := range inarr {

		splitLine := strings.Split(line, "")
		ints := []int{}
		for _, v := range splitLine {
			vInt, _ := strconv.Atoi(v)
			ints = append(ints, vInt)
		}
		grid = append(grid, ints)
	}

	totalRiskLevel := 0
	for y := range grid {
		for x := range grid[y] {
			current := grid[y][x]
			noLower := true
			neighbours := []int{}
			if x > 0 {
				neighbours = append(neighbours, grid[y][x-1])
			}
			if x < len(grid[y])-1 {
				neighbours = append(neighbours, grid[y][x+1])
			}
			if y > 0 {
				neighbours = append(neighbours, grid[y-1][x])
			}
			if y < len(grid)-1 {
				neighbours = append(neighbours, grid[y+1][x])
			}
			for _, n := range neighbours {
				if n <= current {
					noLower = false
					break
				}
			}
			if noLower {
				totalRiskLevel += 1 + current
			}
		}
	}

	fmt.Println(totalRiskLevel)

}

func secondOne(inarr []string) {
	grid := [][]int{}

	for _, line := range inarr {

		splitLine := strings.Split(line, "")
		ints := []int{}
		for _, v := range splitLine {
			vInt, _ := strconv.Atoi(v)
			ints = append(ints, vInt)
		}
		grid = append(grid, ints)
	}

	basinSizes := []int{}
	usedCoords := [][]int{}
	for y := range grid {
		for x := range grid[y] {
			current := grid[y][x]
			noLower := true
			neighbours := []int{}
			if x > 0 {
				neighbours = append(neighbours, grid[y][x-1])
			}
			if x < len(grid[y])-1 {
				neighbours = append(neighbours, grid[y][x+1])
			}
			if y > 0 {
				neighbours = append(neighbours, grid[y-1][x])
			}
			if y < len(grid)-1 {
				neighbours = append(neighbours, grid[y+1][x])
			}
			for _, n := range neighbours {
				if n <= current {
					noLower = false
					break
				}
			}
			if noLower {

				partOfBasinAlready := false
				for _, coord := range usedCoords {
					if coord[0] == y && coord[1] == x {
						partOfBasinAlready = true
						break
					}
				}

				if partOfBasinAlready {
					continue
				}

				moves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

				frontline := [][]int{}
				frontline = append(frontline, []int{y, x})
				usedCoords = append(usedCoords, []int{y, x})
				lastFrontline := [][]int{}
				for len(frontline) != len(lastFrontline) {
					lastFrontline = frontline
					for _, coords := range frontline {
						for _, move := range moves {
							newCoords := []int{coords[0] + move[0], coords[1] + move[1]}
							if newCoords[0] >= 0 && newCoords[0] <= len(grid)-1 && newCoords[1] >= 0 && newCoords[1] <= len(grid[0])-1 {

								if grid[newCoords[0]][newCoords[1]] != 9 {
									exists := false
									for _, pastCoord := range frontline {
										if pastCoord[0] == newCoords[0] && pastCoord[1] == newCoords[1] {
											exists = true
											break
										}
									}
									if !exists {
										frontline = append(frontline, newCoords)
										usedCoords = append(usedCoords, newCoords)
									}
								}
							} else {
								// fmt.Println("new coord filtered", newCoords)
							}
						}
					}
				}
				fmt.Println("basin is ", frontline)
				basinSizes = append(basinSizes, len(frontline))
			}
		}
	}

	sort.Ints(basinSizes)
	res := 1
	for _, n := range basinSizes[len(basinSizes)-3:] {
		res *= n
	}
	fmt.Println(res)

}

func Run() {
	_, inarr := utils.LoadFile("day09", "\n")
	// inarr := strings.Split(instr, ",")
	// inarrInt := []int{}
	// for i := range inarr {
	// 	num, _ := strconv.Atoi(strings.TrimSpace(inarr[i]))
	// 	inarrInt = append(inarrInt, num)
	// }

	firstOne(inarr)
	secondOne(inarr)
}

func Max(s []int) int {
	m := 0
	for i, e := range s {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

func Intersection(a string, b string) string {
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")
	sort.Strings(arrA)
	sort.Strings(arrB)
	output := ""
	for ai := range arrA {
		for ab := range arrB {
			if arrA[ai] == arrB[ab] {
				output += arrA[ai]
			}
		}
	}
	return output
}

func LeaveOnlyUnique(a string, b string) string {
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")
	sort.Strings(arrA)
	sort.Strings(arrB)

	// newA := ""
	newB := ""
	for _, bchar := range arrB {
		shared := false
		for _, achar := range arrA {
			if achar == bchar {
				shared = true
			}
		}
		if !shared {
			newB += bchar
		}
	}
	if newB == "" {
		newB = b
	}

	return newB
}
