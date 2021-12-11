package day06

import (
	"fmt"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

func firstOne(inarr []int) {

	days := 80
	fishes := inarr

	for i := 0; i < days; i++ {
		lenBefore := len(fishes)
		for n := 0; n < lenBefore; n++ {
			if fishes[n] == 0 {
				fishes[n] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[n] = fishes[n] - 1
			}
		}

	}

	fmt.Println(len(fishes))

}

func secondOne(inarr []int) {
	days := 256
	fishes := inarr
	ageCounters := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for n := 0; n < len(fishes); n++ {
		idx := fishes[n]
		ageCounters[idx] += 1
	}

	for i := 0; i < days; i++ {
		newBoys := ageCounters[0]
		ageCounters = Shift(ageCounters)
		ageCounters[6] += newBoys
	}

	sum := 0
	for i := range ageCounters {
		sum += ageCounters[i]
	}
	fmt.Println(sum)

}

func Run() {
	instr, _ := utils.LoadFile("day06", "\n")
	inarr := strings.Split(instr, ",")
	inarrInt := []int{}
	for i := range inarr {
		num, _ := strconv.Atoi(inarr[i])
		inarrInt = append(inarrInt, num)
	}

	firstOne(inarrInt)
	secondOne(inarrInt)
}

func Shift(s []int) []int {
	last := s[0]
	rest := s[1:]
	rest = append(rest, last)
	return rest
}
