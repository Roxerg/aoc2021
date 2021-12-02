package day2

import (
	"fmt"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

func firstOne(inarr []string) {
	horizontal := 0
	depth := 0

	for idx := range inarr {
		splitten := strings.Split(inarr[idx], " ")
		c := splitten[0]
		n, _ := strconv.Atoi(splitten[1])
		switch c {
		case "up":
			depth -= n
		case "down":
			depth += n
		case "forward":
			horizontal += n
		}
	}

	fmt.Println(horizontal * depth)
}

func secondOne(inarr []string) {
	horizontal := 0
	depth := 0
	aim := 0

	for idx := range inarr {
		splitten := strings.Split(inarr[idx], " ")
		c := splitten[0]
		n, _ := strconv.Atoi(splitten[1])

		switch c {
		case "up":
			aim -= n
		case "down":
			aim += n
		case "forward":
			horizontal += n
			depth += n * aim
		}
	}

	fmt.Println(horizontal * depth)

}

func Run() {
	_, inarr := utils.LoadFile("day2", "\n")
	firstOne(inarr)
	secondOne(inarr)
}
