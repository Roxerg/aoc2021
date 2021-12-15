package day15

// import (
// 	"fmt"
// 	"math"
// 	"roxerg_aoc/utils"
// 	"strconv"
// 	"strings"
// )

// type Node struct {
// 	cost          int
// 	totalDistance int
// 	visited       bool
// }

// type Coords struct {
// 	x int
// 	y int
// }

// func firstOne(grid [][]int) {

// 	fmt.Println(grid[0])
// 	unvisited := make(map[Coords]Node)
// 	frontline := make(map[Coords]Node)
// 	frontline[Coords{0, 0}] = Node{0, 0, true}

// 	for y := range grid {
// 		for x := range grid[y] {
// 			unvisited[Coords{x, y}] = Node{grid[x][y], math.MaxInt, false}
// 		}
// 	}

// 	found := false

// 	iter := 0

// 	for len(unvisited) > 0 && !found {
// 		iter += 1
// 		//fmt.Println("front", frontline)
// 		newFrontline := make(map[Coords]Node)
// 		for currCoords, currNode := range frontline {

// 			// fmt.Println(currCoords)

// 			if currCoords.x == len(grid)-1 && currCoords.y == len(grid)-1 {
// 				found = true
// 				break
// 			}

// 			for _, dir := range []Coords{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
// 				nextCoords := addCoords(currCoords, dir)

// 				nextNode, ok := unvisited[nextCoords]
// 				if ok {
// 					// fmt.Println("neigbour coord", nextCoords)
// 					if !nextNode.visited {

// 						newTotal := nextNode.totalDistance
// 						priceToGo := currNode.totalDistance + nextNode.cost

// 						if newTotal > priceToGo {
// 							// fmt.Println(newTotal, ">", currNode.totalDistance, "+", nextNode.cost)
// 							newTotal = priceToGo
// 						}

// 						//nextNode.totalDistance = priceToGo

// 						new, okk := newFrontline[nextCoords]
// 						if okk {
// 							if newTotal < new.totalDistance {
// 								newFrontline[nextCoords] = Node{nextNode.cost, newTotal, nextNode.visited}
// 							}
// 						} else {
// 							newFrontline[nextCoords] = Node{nextNode.cost, newTotal, nextNode.visited}
// 						}
// 						//frontline[nextCoords] = Node{nextNode.cost, priceToGo, false}

// 					}
// 				}

// 			}

// 			currNode.visited = true
// 			unvisited[currCoords] = Node{currNode.cost, currNode.totalDistance, true}
// 			//frontline[currCoords] = Node{currNode.cost, currNode.totalDistance, true}
// 		}
// 		if !found {
// 			if len(frontline)%1000 == 0 {
// 				fmt.Println(len(frontline))
// 			}
// 			frontline = newFrontline
// 		}
// 	}

// 	fmt.Println(frontline)

// }

// func dumbHeuristic(current Coords, end Coords) int {
// 	return 0 //(end.x - current.x) + (end.y - current.y)
// }

// func stupidStinkyAStar(grid [][]int) {

// 	start := Coords{0, 0}
// 	end := Coords{len(grid) - 1, len(grid[0]) - 1}
// 	allNodes := make(map[Coords]int)
// 	backupAllNodes := make(map[Coords]int)

// 	frontline := make(map[Coords]int)
// 	frontline[Coords{0, 0}] = 0

// 	// the value coords is the immediate best prev of key coords
// 	previous := make(map[Coords]Coords)

// 	gScore := make(map[Coords]int)

// 	fScore := make(map[Coords]int)

// 	for y := range grid {
// 		fmt.Println(grid[y])
// 		for x := range grid[y] {
// 			allNodes[Coords{y, x}] = grid[y][x]
// 			backupAllNodes[Coords{y, x}] = grid[y][x]
// 			gScore[Coords{y, x}] = math.MaxInt
// 			fScore[Coords{y, x}] = math.MaxInt
// 		}
// 	}

// 	gScore[start] = 0
// 	fScore[start] = dumbHeuristic(start, end)

// 	for len(frontline) > 0 {
// 		var current Coords
// 		var minF int = math.MaxInt
// 		for coords := range frontline {
// 			heur := dumbHeuristic(coords, end)
// 			if minF > heur {
// 				minF = heur
// 				current = Coords{coords.y, coords.x}
// 			}
// 		}

// 		if current == end {
// 			break
// 		}
// 		fmt.Println("current", current)

// 		delete(frontline, current)
// 		fmt.Println(frontline)
// 		delete(allNodes, current)
// 		//fmt.Println(current)
// 		for _, dir := range []Coords{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
// 			neigbourCoords := Coords{current.y + dir.y, current.x + dir.x}
// 			aaa, ok := allNodes[neigbourCoords]
// 			if !ok {
// 				continue
// 			}
// 			fmt.Println("aaa", aaa)
// 			guessedG := gScore[current] + allNodes[neigbourCoords]
// 			// fmt.Println("guessedG", guessedG, neigbourCoords)
// 			if guessedG < gScore[neigbourCoords] {
// 				previous[neigbourCoords] = current
// 				gScore[neigbourCoords] = guessedG
// 				fScore[neigbourCoords] = guessedG + dumbHeuristic(neigbourCoords, end)/2
// 				_, inFront := frontline[neigbourCoords]
// 				if !inFront {
// 					frontline[neigbourCoords] = allNodes[neigbourCoords]
// 				}

// 			}
// 		}

// 	}

// 	display := [][]string{}
// 	for a := range grid {
// 		display = append(display, []string{})
// 		for range grid[a] {
// 			display[a] = append(display[a], "0")
// 		}
// 	}

// 	cost := 0
// 	found := true
// 	next := end
// 	fmt.Println(previous)
// 	display[end.y][end.x] = "E"
// 	for found {
// 		coord, found := previous[next]
// 		fmt.Println(previous)
// 		if found {
// 			display[coord.y][coord.x] = "X"
// 			cost += backupAllNodes[coord]
// 			next = coord
// 			if next == start {
// 				break
// 			}
// 		}
// 	}

// 	for a := range display {
// 		for b := range display[a] {
// 			fmt.Print(display[a][b])
// 		}
// 		fmt.Println("")
// 	}

// 	fmt.Println(cost)

// }

// func secondOne(grid [][]int) {

// 	origY := len(grid)
// 	origX := len(grid[0])

// 	newgrid := make([][]int, origY*5)
// 	for i := range newgrid {
// 		newgrid[i] = make([]int, origX*5)
// 	}

// 	// fmt.Println("oldgrid", grid)

// 	// do one big row first
// 	for I := 0; I < 5; I++ {
// 		//fmt.Println("oldgrid", grid)
// 		//offsetX := I * origX
// 		offsetY := I * origY
// 		for y := range grid {
// 			for x := range grid[y] {
// 				newgrid[offsetY+y][x] = grid[y][x]
// 				//fmt.Println("set", offsetY+y, x)

// 			}
// 		}

// 		for y := range grid {
// 			for x := range grid[y] {
// 				grid[y][x] = addWithWrap(grid[y][x])
// 			}
// 		}
// 	}

// 	// fmt.Println("newgrid")
// 	// for i := range newgrid {
// 	// 	fmt.Println(newgrid[i])
// 	// }

// 	bigRow := make([][]int, len(newgrid))
// 	for i := range bigRow {
// 		bigRow[i] = make([]int, origX)
// 	}

// 	for i := range bigRow {
// 		bigRow[i] = append([]int{}, newgrid[i][:origX]...)
// 	}

// 	// fmt.Println("bigrow")
// 	// for i := range bigRow {
// 	// 	fmt.Println(bigRow[i])
// 	// }

// 	// do rows

// 	for I := 1; I < 5; I++ {
// 		offsetX := I * origX
// 		for y := range bigRow {
// 			for x := range bigRow[y] {
// 				bigRow[y][x] = addWithWrap(bigRow[y][x])
// 			}
// 		}

// 		// offsetY := I * origY
// 		for y := range bigRow {
// 			for x := range bigRow[y] {
// 				newgrid[y][offsetX+x] = bigRow[y][x]
// 			}
// 		}

// 	}

// 	// fmt.Println("final")
// 	// for i := range newgrid {
// 	// 	fmt.Println(newgrid[i])
// 	// }
// 	fmt.Println("----")

// 	stupidStinkyAStar(newgrid)
// }

// func Run() {
// 	_, inarr := utils.LoadFile("day15", "\n")

// 	grid := [][]int{}
// 	for idx, line := range inarr {
// 		grid = append(grid, []int{})
// 		for _, coststr := range strings.Split(line, "") {
// 			cost, _ := strconv.Atoi(coststr)
// 			grid[idx] = append(grid[idx], cost)
// 		}
// 	}

// 	// firstOne(grid)
// 	stupidStinkyAStar(grid)
// 	//secondOne(grid) //2990 is too high
// }

// func addCoords(a Coords, b Coords) Coords {
// 	return Coords{a.x + b.x, a.y + b.y}
// }

// func addWithWrap(x int) int {
// 	x += 1
// 	if x > 9 {
// 		return 1
// 	}
// 	return x
// }
