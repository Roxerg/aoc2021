package day03

import (
	"fmt"
	"math"
	"roxerg_aoc/utils"
	"strconv"
)

func firstOne(inarr []string) {

	counts := [12]int{}

	for c := range counts {
		counts[c] = 0
	}

	for idx := range inarr {
		for bit_idx := range inarr[idx] {
			if bit_idx > 11 {
				continue
			}
			c, _ := strconv.Atoi(string(inarr[idx][bit_idx]))
			counts[bit_idx] += c
		}
	}

	str_g := ""
	majority := len(inarr) / 2
	for c := range counts {
		if counts[c] > majority {
			str_g += "1"

		} else {
			str_g += "0"
		}
	}

	str_e := ""
	for c := range str_g {
		if string(str_g[c]) == "0" {
			str_e += "1"
		} else {
			str_e += "0"
		}
	}

	gamma_int, _ := strconv.ParseInt(str_g, 2, 64)
	epsilon_int, _ := strconv.ParseInt(str_e, 2, 64)

	fmt.Println(gamma_int * epsilon_int)
}

func secondOne(inarr []string) {

	// one, _ := strconv.ParseInt(DoTheDidgeridoo(inarr, true), 2, 64)
	// two, _ := strconv.ParseInt(DoTheDidgeridoo(inarr, false), 2, 64)

	one := DoTheDidgeridoo(inarr, true)
	two := DoTheDidgeridoo(inarr, false)

	// 111100111111

	fmt.Println(one)
	fmt.Println(two)

}

func Run() {
	_, inarr := utils.LoadFile("day03", "\n")
	// firstOne(inarr)
	secondOne(inarr)
}

func DoTheDidgeridoo(input []string, op bool) string {

	current_idx := 0
	inputs := input

	for len(inputs) != 1 {

		var most_common string
		count := 0

		for idx := range inputs {
			var compare string = "1"
			if !op {
				compare = "0"
			}

			if string(inputs[idx][current_idx]) == compare {
				count += 1
			}
		}

		if op {
			if count >= int(math.Ceil((float64(len(inputs)))/2)) {
				most_common = "1"
			} else {
				most_common = "0"
			}
		} else {
			if count <= int(math.Ceil((float64(len(inputs)))/2)) {
				most_common = "0"
			} else {
				most_common = "1"
			}
		}

		temp := []string{}

		for idx := range inputs {
			if string(inputs[idx][current_idx]) == most_common {
				temp = append(temp, inputs[idx])
			}
		}

		fmt.Println("-----")
		for idx := range inputs {
			fmt.Println(inputs[idx])
		}
		fmt.Println("-----")

		if op {
			fmt.Println(" most common:", most_common, "idx", current_idx, "len", len(inputs), "\n")
			fmt.Println(inputs[0])
		}

		inputs = temp
		current_idx = current_idx + 1
		count = 0

	}

	return inputs[0]

}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
