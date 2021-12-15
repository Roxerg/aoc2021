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

	fmt.Println(grid[0])
	unvisited := make(map[Coords]Node)
	frontline := make(map[Coords]Node)
	frontline[Coords{0, 0}] = Node{0, 0, true}

	for y := range grid {
		for x := range grid[y] {
			unvisited[Coords{x, y}] = Node{grid[y][x], math.MaxInt, false}
		}
	}

	found := false

	iter := 0

	for len(unvisited) > 0 && !found {
		iter += 1
		//fmt.Println("front", frontline)
		newFrontline := make(map[Coords]Node)
		for currCoords, currNode := range frontline {

			if currCoords.x == len(grid)-1 && currCoords.y == len(grid)-1 {
				found = true
				break
			}

			for _, dir := range []Coords{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nextCoords := addCoords(currCoords, dir)

				nextNode, ok := unvisited[nextCoords]
				if ok {
					// fmt.Println("neigbour coord", nextCoords)
					if !nextNode.visited {

						newTotal := nextNode.totalDistance
						priceToGo := currNode.totalDistance + nextNode.cost

						if newTotal > priceToGo {
							// fmt.Println(newTotal, ">", currNode.totalDistance, "+", nextNode.cost)
							newTotal = priceToGo
						}

						nextNode.totalDistance = priceToGo

						new, okk := newFrontline[nextCoords]
						if okk {
							if newTotal < new.totalDistance {
								newFrontline[nextCoords] = Node{nextNode.cost, priceToGo, false}
							}
						} else {
							newFrontline[nextCoords] = Node{nextNode.cost, priceToGo, false}
						}
						// frontline[nextCoords] = Node{nextNode.cost, priceToGo, false}

					}
				}

			}
			currNode.visited = true
			frontline[currCoords] = Node{currNode.cost, currNode.totalDistance, true}
			delete(unvisited, currCoords)
		}
		if !found {
			frontline = newFrontline
		}
		// if iter == 3 {
		// 	break
		// }
	}
	fmt.Println(frontline)

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
		//fmt.Println("oldgrid", grid)
		//offsetX := I * origX
		offsetY := I * origY
		for y := range grid {
			for x := range grid[y] {
				newgrid[offsetY+y][x] = grid[y][x]
				//fmt.Println("set", offsetY+y, x)

			}
		}

		for y := range grid {
			for x := range grid[y] {
				grid[y][x] = addWithWrap(grid[y][x])
			}
		}
	}

	// fmt.Println("newgrid")
	// for i := range newgrid {
	// 	fmt.Println(newgrid[i])
	// }

	bigRow := make([][]int, len(newgrid))
	for i := range bigRow {
		bigRow[i] = make([]int, origX)
	}

	for i := range bigRow {
		bigRow[i] = append([]int{}, newgrid[i][:origX]...)
	}

	fmt.Println("bigrow")
	for i := range bigRow {
		fmt.Println(bigRow[i])
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

	fmt.Println("final")
	for i := range newgrid {
		fmt.Println(newgrid[i])
	}
	fmt.Println("----")

	firstOne(newgrid)
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

	firstOne(grid)
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
