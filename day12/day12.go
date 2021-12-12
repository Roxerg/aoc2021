package day12

import (
	"fmt"
	"roxerg_aoc/utils"
	"strings"
)

func firstOne(theMap map[string][]string) {

	allPaths := [][]string{{"start"}}
	beforePathsLen := 0

	for beforePathsLen != len(allPaths) {
		beforePathsLen = len(allPaths)
		allPaths = forkOut(allPaths, theMap)
	}

	//fmt.Println("stabilised at: ")
	//fmt.Println(allPaths)

	goodPathCount := 0
	for _, path := range allPaths {
		if path[len(path)-1] == "end" {
			goodPathCount += 1
			//fmt.Println(path)
		}
	}

	fmt.Println(goodPathCount)
}

func secondOne(theMap map[string][]string) {
	allPaths := [][]string{{"start"}}
	beforePathsLen := 0

	for beforePathsLen != len(allPaths) {
		beforePathsLen = len(allPaths)
		allPaths = forkOutTwoElectricBoogaloo(allPaths, theMap)
	}

	goodPathCount := 0
	for _, path := range allPaths {
		if path[len(path)-1] == "end" {
			goodPathCount += 1
			//fmt.Println(path)
		}
	}

	fmt.Println(goodPathCount)
}

func Run() {
	_, inarr := utils.LoadFile("day12", "\n")

	input := [][]string{}
	for _, line := range inarr {
		input = append(input, strings.Split(line, "-"))
	}
	theMap := buildMap(input)
	fmt.Println(theMap)

	//fmt.Println("start", theMap["start"])

	firstOne(theMap)
	secondOne(theMap)
}

func buildMap(input [][]string) map[string][]string {
	theMap := make(map[string][]string)
	for _, path := range input {
		if path[1] != "start" && path[0] != "end" {
			if _, ok := theMap[path[0]]; ok {
				alreadyThere := false
				for _, existingVal := range theMap[path[0]] {
					if existingVal == path[1] {
						alreadyThere = true
						break
					}
				}
				if !alreadyThere {
					theMap[path[0]] = append(theMap[path[0]], path[1])
				}
			} else {
				theMap[path[0]] = []string{path[1]}
			}
		}

		// flip the bitch
		if path[0] != "start" && path[1] != "end" {
			if _, ok := theMap[path[1]]; ok {
				alreadyThere := false
				for _, existingVal := range theMap[path[1]] {
					if existingVal == path[0] {
						alreadyThere = true
						break
					}
				}
				if !alreadyThere {
					theMap[path[1]] = append(theMap[path[1]], path[0])
				}
			} else {
				theMap[path[1]] = []string{path[0]}
			}
		}

	}

	return theMap
}

func forkOut(thePaths [][]string, theMap map[string][]string) [][]string {

	idx := 0
	newPaths := [][]string{}
	for _, path := range thePaths {
		idx += 1
		last := path[len(path)-1]
		options, ok := theMap[last]
		if ok {
			filteredOptions := []string{}

			for _, opt := range options {
				isSmol := strings.ToLower(opt) == opt
				if isSmol {

					inPath := false
					for _, cave := range path {
						if cave == opt {
							inPath = true
							break
						}
					}

					if !inPath {
						filteredOptions = append(filteredOptions, opt)
					}
				} else {
					filteredOptions = append(filteredOptions, opt)
				}
			}

			// fmt.Println(path, "had these options", filteredOptions)
			for _, option := range filteredOptions {

				newPath := append(make([]string, 0, len(path)), append(path, option)...)
				newPaths = append(newPaths, newPath)
			}

			if len(filteredOptions) == 0 {
				newPath := append(make([]string, 0, len(path)), path...)
				newPaths = append(newPaths, newPath)
			}
		} else {
			newPath := append(make([]string, 0, len(path)), path...)
			newPaths = append(newPaths, newPath)
		}
		//fmt.Println(idx, newPaths)
	}

	// fmt.Println("IN ", thePaths)
	//fmt.Println("OUT", newPaths)
	//fmt.Println("------------")
	return newPaths
}

func forkOutTwoElectricBoogaloo(thePaths [][]string, theMap map[string][]string) [][]string {
	newPaths := [][]string{}

	for _, path := range thePaths {

		littleBoyeCounter := make(map[string]int)
		twiceDone := false

		for _, cave := range path {
			if strings.ToLower(cave) == cave && cave != "start" && cave != "end" {
				_, ok := littleBoyeCounter[cave]
				if ok {
					littleBoyeCounter[cave] += 1
					if littleBoyeCounter[cave] > 1 {
						twiceDone = true
					}
					if littleBoyeCounter[cave] > 2 {
						fmt.Println("I FUCKED UP")
					}
				} else {
					littleBoyeCounter[cave] = 1
				}
			}
		}

		last := path[len(path)-1]
		options, ok := theMap[last]
		if ok {
			filteredOptions := []string{}

			for _, opt := range options {
				isSmol := strings.ToLower(opt) == opt && opt != "start" && opt != "end"
				if isSmol {

					inPath := false
					for _, cave := range path {
						if cave == opt {
							inPath = true
							break
						}
					}

					if !inPath || !twiceDone {

						filteredOptions = append(filteredOptions, opt)
					}
				} else {
					filteredOptions = append(filteredOptions, opt)
				}
			}

			// fmt.Println(path, "had these options", filteredOptions)
			for _, option := range filteredOptions {

				newPath := append(make([]string, 0, len(path)), append(path, option)...)
				newPaths = append(newPaths, newPath)
			}

			if len(filteredOptions) == 0 {
				newPath := append(make([]string, 0, len(path)), path...)
				newPaths = append(newPaths, newPath)
			}
		} else {
			newPath := append(make([]string, 0, len(path)), path...)
			newPaths = append(newPaths, newPath)
		}
		//fmt.Println(idx, newPaths)
	}

	// fmt.Println("IN ", thePaths)
	//fmt.Println("OUT", newPaths)
	//fmt.Println("------------")
	return newPaths
}

// func findSmallCaves(input [][]string) []string {

// 	smallCaves := []string{}
// 	for _, path := range input {
// 		for _, cave := range path {
// 			if cave == "start" || cave == "end" {
// 				continue
// 			}

// 			if strings.ToLower(cave) != cave {
// 				fmt.Println(strings.ToLower(cave), "!=", cave)
// 				continue
// 			}

// 			alreadyThere := false
// 			for _, existing := range smallCaves {
// 				if cave == existing {
// 					alreadyThere = true
// 					break
// 				}
// 			}

// 			if !alreadyThere {
// 				smallCaves = append(smallCaves, cave)
// 			}

// 		}

// 	}

// 	return smallCaves
// }
