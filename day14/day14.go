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
				// fmt.Println(insert[0][:1], insert[0][1:])
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
			fmt.Println(idx, template[idx], newTemplate)
		}
		template = append(make([]string, 0, len(newTemplate)), newTemplate...)
		fmt.Println(template)
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
		_, letterOk := lettersCounter[template[idx]]
		_, pairOk := pairsCounter[template[idx]+template[idx+1]]

		if letterOk {
			lettersCounter[template[idx]] += 1
		} else {
			lettersCounter[template[idx]] = 1
		}

		if pairOk {
			pairsCounter[template[idx]+template[idx+1]] += 1
		} else {
			pairsCounter[template[idx]+template[idx+1]] = 1
		}

		if idx == len(template)-2 {
			_, lastLetterOk := lettersCounter[template[idx+1]]
			if lastLetterOk {
				lettersCounter[template[idx+1]] += 1
			} else {
				lettersCounter[template[idx+1]] = 1
			}
		}
	}

	// fmt.Println(lettersCounter)
	fmt.Println(pairsCounter)

	for step := uint64(0); step < iters; step++ {
		tempPairCounts := make(map[string]uint64)
		for _, insert := range insertions {

			oldPair := insert[0]
			_, ok := pairsCounter[oldPair]
			if ok {
				c := pairsCounter[oldPair]
				if c > 0 {
					newLetter := insert[1]
					newPairL := insert[0][:1] + newLetter
					newPairR := newLetter + insert[0][1:]
					//fmt.Println("old: ", oldPair, " insertletter: ", newLetter, " newLeft: ", newPairL, "newRight: ", newPairR)
					incrementOrInitialize(lettersCounter, newLetter, c)
					incrementOrInitialize(tempPairCounts, newPairL, c)
					incrementOrInitialize(tempPairCounts, newPairR, c)
					decrement(tempPairCounts, oldPair, c)

					// lettersCounter = incrementOrInitialize(lettersCounter, newLetter)
					// pairsCounter = incrementOrInitialize(pairsCounter, newPairL)
					// pairsCounter = incrementOrInitialize(pairsCounter, newPairR)
					// pairsCounter = decrementIfExists(pairsCounter, oldPair)
				} else {

				}
			}

		}
		for key, val := range tempPairCounts {
			_, ok := pairsCounter[key]
			if ok {
				pairsCounter[key] += val
			} else {
				// fmt.Println(key, "didnt exist")
				pairsCounter[key] = val
			}
			if pairsCounter[key] == 0 {
				delete(pairsCounter, key)
			}
		}

		fmt.Println(step + 1) //"pairs", pairsCounter, "letters", lettersCounter)

	}

	max := uint64(0)
	min := uint64(18446744073709551615)

	//fmt.Println(pairsCounter)
	//fmt.Println(lettersCounter)

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

	//firstOne(template, insertions, 10)
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

func decrement(counters map[string]uint64, key string, c uint64) uint64 {
	_, ok := counters[key]
	if ok {
		counters[key] -= c
	} else {
		counters[key] = -c
	}
	return c
}
