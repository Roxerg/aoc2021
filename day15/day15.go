package day15

import (
	"fmt"
	"math"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

type Node struct {
	cost          int
	totalDistance int
	visited       bool
}

type Coords struct {
	x int
	y int
}

func firstOne(grid [][]int) {

	fmt.Println("")
}

func remove(slice [][2]int, s int) [][2]int {
	return append(slice[:s], slice[s+1:]...)
}

func withinBounds(grid [][]int, coords [2]int) bool {
	yy := len(grid)
	xx := len(grid[0])
	return coords[0] >= 0 && coords[0] < yy && coords[1] >= 0 && coords[1] < xx
}

func AFuckingStar(grid [][]int) {
	visited := [][]bool{}
	frontlineCoords := [][2]int{}
	costSoFar := [][]int{}
	goal := [2]int{len(grid) - 1, len(grid[0]) - 1}

	for y := range grid {
		visited = append(visited, []bool{})
		costSoFar = append(costSoFar, []int{})
		for range grid[y] {
			visited[y] = append(visited[y], false)
			costSoFar[y] = append(costSoFar[y], math.MaxInt)
		}
	}
	costSoFar[0][0] = 0
	frontlineCoords = append(frontlineCoords, [2]int{0, 0})

	iter := 0
	for len(frontlineCoords) > 0 {
		iter += 1
		// fmt.Println(frontlineCoords)
		currentCoords := [2]int{}
		var currentCost int
		min := math.MaxInt
		minIdx := 0
		for idx, coord := range frontlineCoords {
			if costSoFar[coord[0]][coord[1]] < min {
				minIdx = idx
				min = costSoFar[coord[0]][coord[1]]
			}
		}

		if currentCoords[0] == goal[0] && currentCoords[1] == goal[1] {
			fmt.Println(currentCost)
		}

		currentCoords = [2]int{frontlineCoords[minIdx][0], frontlineCoords[minIdx][1]}
		currentCost = costSoFar[currentCoords[0]][currentCoords[1]]

		frontlineCoords = remove(frontlineCoords, minIdx)
		visited[currentCoords[0]][currentCoords[1]] = true

		// fmt.Println("---")
		// for _, v := range visited {
		// 	fmt.Println(v)
		// }

		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			neighbourCoord := [2]int{currentCoords[0] + dir[0], currentCoords[1] + dir[1]}

			ok := withinBounds(grid, neighbourCoord)
			if ok {
				skip := visited[neighbourCoord[0]][neighbourCoord[1]]
				if skip {
					continue
				} else {
					alreadyInFrontline := false

					for _, f := range frontlineCoords {
						if neighbourCoord[0] == f[0] && neighbourCoord[1] == f[1] {
							alreadyInFrontline = true
							break
						}
					}

					if !alreadyInFrontline {
						frontlineCoords = append(frontlineCoords, [2]int{neighbourCoord[0], neighbourCoord[1]})
					}

				}
			} else {
				continue
			}

			withCurrentCost := currentCost + grid[neighbourCoord[0]][neighbourCoord[1]]
			if withCurrentCost < costSoFar[neighbourCoord[0]][neighbourCoord[1]] {
				costSoFar[neighbourCoord[0]][neighbourCoord[1]] = withCurrentCost
			}

		}

	}

	// fmt.Println("---")
	// for _, v := range visited {
	// 	fmt.Println(v)
	// }
	fmt.Println(costSoFar[goal[0]][goal[1]])

}

func secondOne(grid [][]int) {

	origY := len(grid)
	origX := len(grid[0])

	newgrid := make([][]int, origY*5)
	for i := range newgrid {
		newgrid[i] = make([]int, origX*5)
	}

	// fmt.Println("oldgrid", grid)

	// do one big row first
	for I := 0; I < 5; I++ {
		offsetY := I * origY
		for y := range grid {
			for x := range grid[y] {
				newgrid[offsetY+y][x] = grid[y][x]
			}
		}

		for y := range grid {
			for x := range grid[y] {
				grid[y][x] = addWithWrap(grid[y][x])
			}
		}
	}

	bigRow := make([][]int, len(newgrid))
	for i := range bigRow {
		bigRow[i] = make([]int, origX)
	}

	for i := range bigRow {
		bigRow[i] = append([]int{}, newgrid[i][:origX]...)
	}

	// do rows

	for I := 1; I < 5; I++ {
		offsetX := I * origX
		for y := range bigRow {
			for x := range bigRow[y] {
				bigRow[y][x] = addWithWrap(bigRow[y][x])
			}
		}

		// offsetY := I * origY
		for y := range bigRow {
			for x := range bigRow[y] {
				newgrid[y][offsetX+x] = bigRow[y][x]
			}
		}

	}

	// fmt.Println("final")
	// for i := range newgrid {
	// 	fmt.Println(newgrid[i])
	// }
	fmt.Println("----")

	AFuckingStar(newgrid)
}

func Run() {
	_, inarr := utils.LoadFile("day15", "\n")

	grid := [][]int{}
	for idx, line := range inarr {
		grid = append(grid, []int{})
		for _, coststr := range strings.Split(line, "") {
			cost, _ := strconv.Atoi(coststr)
			grid[idx] = append(grid[idx], cost)
		}
	}

	//AFuckingStar(grid)
	secondOne(grid)
}

func addCoords(a Coords, b Coords) Coords {
	return Coords{a.x + b.x, a.y + b.y}
}

func addWithWrap(x int) int {
	x += 1
	if x > 9 {
		return 1
	}
	return x
}
