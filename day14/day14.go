package day14

import (
	"fmt"
	"roxerg_aoc/utils"
	"strings"
)

func firstOne(template []string, insertions [][]string, iters uint64) {

	for step := uint64(0); step < iters; step++ {
		newTemplate := []string{}
		for idx := range template[:len(template)-1] {
			inserted := false
			for _, insert := range insertions {
				if template[idx] == insert[0][:1] && template[idx+1] == insert[0][1:] {
					inserted = true
					if idx == 0 {
						newTemplate = append(newTemplate, template[idx])
					}
					newTemplate = append(newTemplate, insert[1])
					newTemplate = append(newTemplate, template[idx+1])
				}
			}
			if !inserted {
				newTemplate = append(newTemplate, template[idx+1])
			}
		}
		template = append(make([]string, 0, len(newTemplate)), newTemplate...)
	}

	counts := make(map[string]int)
	for _, e := range template {
		_, ok := counts[e]
		if !ok {
			counts[e] = 1
		} else {
			counts[e] += 1
		}
	}
	max := 0
	min := 99999999

	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	fmt.Println(max - min)
}

func secondOne(template []string, insertions [][]string, iters uint64) {

	lettersCounter := make(map[string]uint64)
	pairsCounter := make(map[string]uint64)

	for idx := range template[:len(template)-1] {
		incrementOrInitialize(lettersCounter, template[idx], 1)
		incrementOrInitialize(pairsCounter, template[idx]+template[idx+1], 1)

		if idx == len(template)-2 {
			incrementOrInitialize(lettersCounter, template[idx+1], 1)
		}
	}

	for step := uint64(0); step < iters; step++ {
		tempPairCounts := make(map[string]uint64)
		for _, insert := range insertions {

			oldPair := insert[0]
			_, ok := pairsCounter[oldPair]
			if ok {
				c := pairsCounter[oldPair]
				if c > 0 {
					newLetter := insert[1]
					incrementOrInitialize(lettersCounter, newLetter, c)
					incrementOrInitialize(tempPairCounts, insert[0][:1]+newLetter, c)
					incrementOrInitialize(tempPairCounts, newLetter+insert[0][1:], c)
					decrementOrInitialize(tempPairCounts, oldPair, c)
				} else {

				}
			}

		}
		for key, val := range tempPairCounts {
			incrementOrInitialize(pairsCounter, key, val)
			if pairsCounter[key] == 0 {
				delete(pairsCounter, key)
			}
		}

	}

	max := uint64(0)
	min := uint64(18446744073709551615)

	for _, count := range lettersCounter {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	fmt.Println(max - min)
}

func Run() {
	_, inarr := utils.LoadFile("day14", "\n")
	template := strings.Split(inarr[0], "")
	insertions := [][]string{}
	for _, line := range inarr[2:] {
		insertions = append(insertions, strings.Split(line, " -> "))
	}

	// naive solution
	firstOne(template, insertions, 10)

	// what i should have done from the start
	secondOne(template, insertions, 40)
}

func incrementOrInitialize(counters map[string]uint64, key string, c uint64) uint64 {
	_, ok := counters[key]
	if ok {
		counters[key] += c
	} else {
		counters[key] = c
	}
	return c
}

func decrementOrInitialize(counters map[string]uint64, key string, c uint64) uint64 {
	return incrementOrInitialize(counters, key, -c)
}
