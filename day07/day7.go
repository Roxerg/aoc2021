package day07

import (
	"fmt"
	"math"
	"roxerg_aoc/utils"
	"sort"
	"strconv"
	"strings"
)

func firstOne(inarr []int) {
	sort.Ints(inarr)
	median := int(math.Abs(float64((inarr[((len(inarr))/2)-1] + inarr[((len(inarr))/2)]) / 2)))
	cost := 0
	for i := range inarr {
		cost += int(math.Abs(float64(inarr[i] - median)))
	}
	fmt.Println(cost)
}

func secondOne(inarr []int) {
	cost := math.MaxInt
	for x := 0; x < Max(inarr); x++ {
		newCost := 0
		for i := range inarr {
			moves := int(math.Abs(float64(inarr[i] - x)))
			for fuelCost := 1; fuelCost <= moves; fuelCost++ {
				newCost += fuelCost
			}
		}
		if newCost < cost {
			cost = newCost
		}
	}
	fmt.Println(cost)
}

func Run() {
	instr, _ := utils.LoadFile("day07", "\n")
	inarr := strings.Split(instr, ",")
	inarrInt := []int{}
	for i := range inarr {
		num, _ := strconv.Atoi(strings.TrimSpace(inarr[i]))
		inarrInt = append(inarrInt, num)
	}

	firstOne(inarrInt)
	secondOne(inarrInt)
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
