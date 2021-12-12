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

	fmt.Println("stabilised at: ")
	fmt.Println(allPaths)
}

func secondOne(inarr []string) {
	fmt.Println("")
}

func Run() {
	_, inarr := utils.LoadFile("day12", "\n")

	input := [][]string{}
	for _, line := range inarr {
		input = append(input, strings.Split(line, "-"))
	}

	smallCaves := findSmallCaves(input)
	theMap := buildMap(input)
	fmt.Println(theMap)
	fmt.Println(smallCaves)

	firstOne(theMap)
	//secondOne(inarr)
}

func findSmallCaves(input [][]string) []string {

	smallCaves := []string{}
	for _, path := range input {
		for _, cave := range path {
			if cave == "start" || cave == "end" {
				continue
			}

			if strings.ToLower(cave) != cave {
				fmt.Println(strings.ToLower(cave), "!=", cave)
				continue
			}

			alreadyThere := false
			for _, existing := range smallCaves {
				if cave == existing {
					alreadyThere = true
					break
				}
			}

			if !alreadyThere {
				smallCaves = append(smallCaves, cave)
			}

		}

	}

	return smallCaves
}

func buildMap(input [][]string) map[string][]string {
	theMap := make(map[string][]string)
	for _, path := range input {
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

			fmt.Println(path, "had these options", filteredOptions)
			for _, option := range filteredOptions {

				newPath := append(path, option)
				fmt.Println("new path added: ", newPath)
				newPaths = append(newPaths, newPath)
			}

			if len(filteredOptions) == 0 {
				newPaths = append(newPaths, path)
			}
		} else {
			newPaths = append(newPaths, path)
		}
		//fmt.Println(idx, newPaths)
	}

	fmt.Println("IN ", thePaths)
	fmt.Println("OUT", newPaths)
	fmt.Println("------------")
	return newPaths
}
