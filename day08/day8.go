package day08

import (
	"fmt"
	"math"
	"roxerg_aoc/utils"
	"sort"
	"strings"
)

func firstOne(inarr []string) {

	inputs := [][]string{}
	outputs := [][]string{}

	uniqueSegLens := []int{2, 3, 4, 7}

	for _, line := range inarr {
		inAndOut := strings.Split(line, "|")
		in := strings.Split(strings.TrimSpace(inAndOut[0]), " ")
		out := strings.Split(strings.TrimSpace(inAndOut[1]), " ")
		inputs = append(inputs, in)
		outputs = append(outputs, out)
	}

	counter := 0

	for _, outputarr := range outputs {
		for _, o := range outputarr {
			for _, s := range uniqueSegLens {
				if len(o) == s {
					counter += 1
					break
				}
			}
		}
	}

	fmt.Println(counter)
}

// func secondOne(inarr []string) {

// 	inputs := [][]string{}
// 	outputs := [][]string{}

// 	uniqueSegLens := []int{2, 3, 4, 5, 6, 7}
// 	sevenLen := 3
// 	oneLen := 2
// 	fourLen := 4
// 	eightLen := 7
// 	sixOrZeroOrNine := 6
// 	twoThreeOrFive := 5

// 	// segments := "abcdefg"

// 	top := "abcdefg"
// 	middle := "abcdefg"
// 	bottom := "abcdefg"
// 	upperLeft := "abcdefg"
// 	upperRight := "abcdefg"
// 	bottomLeft := "abcdefg"
// 	bottomRight := "abcdefg"

// 	for _, line := range inarr {
// 		inAndOut := strings.Split(line, "|")
// 		in := strings.Split(strings.TrimSpace(inAndOut[0]), " ")
// 		out := strings.Split(strings.TrimSpace(inAndOut[1]), " ")
// 		inputs = append(inputs, in)
// 		outputs = append(outputs, out)
// 	}

// 	// counter := 0

// 	for _, inputarr := range inputs {
// 		for _, o := range inputarr {
// 			for _, s := range uniqueSegLens {
// 				if len(o) == s {
// 					if len(o) == sevenLen {
// 						top = Intersection(o, top)
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 					} else if len(o) == oneLen {
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 					} else if len(o) == fourLen {
// 						upperLeft = Intersection(o, upperLeft)
// 						middle = Intersection(o, middle)
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 					} else if len(o) == eightLen {
// 						upperLeft = Intersection(o, upperLeft)
// 						middle = Intersection(o, middle)
// 						top = Intersection(o, top)
// 						bottom = Intersection(o, bottom)
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 						bottomLeft = Intersection(o, bottomLeft)
// 					} else if len(o) == sixOrZeroOrNine {
// 						bottomRight = Intersection(o, bottomRight)
// 						top = Intersection(o, top)
// 						bottom = Intersection(o, bottom)
// 						upperLeft = Intersection(o, upperLeft)
// 					} else if len(o) == twoThreeOrFive {
// 						top = Intersection(o, top)
// 						bottom = Intersection(o, bottom)
// 						middle = Intersection(o, middle)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	for _, inputarr := range outputs {
// 		for _, o := range inputarr {
// 			for _, s := range uniqueSegLens {
// 				if len(o) == s {
// 					if len(o) == sevenLen {
// 						top = Intersection(o, top)
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 					} else if len(o) == oneLen {
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 					} else if len(o) == fourLen {
// 						upperLeft = Intersection(o, upperLeft)
// 						middle = Intersection(o, middle)
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 					} else if len(o) == eightLen {
// 						upperLeft = Intersection(o, upperLeft)
// 						middle = Intersection(o, middle)
// 						top = Intersection(o, top)
// 						bottom = Intersection(o, bottom)
// 						upperRight = Intersection(o, upperRight)
// 						bottomRight = Intersection(o, bottomRight)
// 						bottomLeft = Intersection(o, bottomLeft)
// 					} else if len(o) == sixOrZeroOrNine {
// 						bottomRight = Intersection(o, bottomRight)
// 						top = Intersection(o, top)
// 						bottom = Intersection(o, bottom)
// 						upperLeft = Intersection(o, upperLeft)
// 					} else if len(o) == twoThreeOrFive {
// 						top = Intersection(o, top)
// 						bottom = Intersection(o, bottom)
// 						middle = Intersection(o, middle)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	//                            top, middle, bottom, upleft, upright, downleft, downright
// 	//                            0,    1         2      3        4        5          6
// 	segmentDecode := []string{top, middle, bottom, upperLeft, upperRight, bottomLeft, bottomRight}
// 	for seg := range segmentDecode {
// 		for otherseg := range segmentDecode {
// 			if seg != otherseg {
// 				if len(segmentDecode[seg]) > len(segmentDecode[otherseg]) {
// 					segmentDecode[seg] = LeaveOnlyUnique(segmentDecode[otherseg], segmentDecode[seg])
// 				} else {
// 					segmentDecode[otherseg] = LeaveOnlyUnique(segmentDecode[seg], segmentDecode[otherseg])
// 				}
// 			}
// 		}
// 	}

// 	s := segmentDecode
// 	codes := []string{
// 		s[6] + s[4],                                    // 1
// 		s[0] + s[1] + s[2] + s[4] + s[5],               // 2
// 		s[0] + s[1] + s[2] + s[4] + s[6],               // 3
// 		s[1] + s[3] + s[4] + s[6],                      // 4
// 		s[0] + s[1] + s[2] + s[3] + s[6],               // 5
// 		s[0] + s[1] + s[5] + s[6] + s[3] + s[2],        // 6
// 		s[0] + s[5] + s[6],                             // 7
// 		s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6], // 8
// 		s[0] + s[1] + s[2] + s[3] + s[4] + s[6],        // 9
// 	}

// 	decoder := make(map[string]int)
// 	for i, code := range codes {
// 		codearr := strings.Split(code, "")
// 		sort.Strings(codearr)
// 		decoder[strings.Join(codearr, "")] = i + 1
// 	}

// 	sum := 0
// 	for _, inputarr := range outputs {
// 		num := 0
// 		for idx, o := range inputarr {
// 			oarr := strings.Split(o, "")
// 			sort.Strings(oarr)
// 			num += decoder[strings.Join(oarr, "")] * int(math.Pow10(3-idx))
// 		}
// 		fmt.Println(num)
// 		sum += num
// 	}

// 	fmt.Println(sum)

// }

func secondOne(inarr []string) {

	fmt.Println("----")

	inputs := [][]string{}
	outputs := [][]string{}

	uniqueSegLens := []int{2, 3, 4, 5, 6, 7}
	sevenLen := 3
	oneLen := 2
	fourLen := 4
	eightLen := 7
	sixOrZeroOrNine := 6
	twoThreeOrFive := 5

	// segments := "abcdefg"

	for _, line := range inarr {
		inAndOut := strings.Split(line, "|")
		in := strings.Split(strings.TrimSpace(inAndOut[0]), " ")
		out := strings.Split(strings.TrimSpace(inAndOut[1]), " ")
		inputs = append(inputs, in)
		outputs = append(outputs, out)
	}

	// counter := 0
	sum := 0

	for bigIdx, inputarr := range inputs {

		top := "abcdefg"
		middle := "abcdefg"
		bottom := "abcdefg"
		upperLeft := "abcdefg"
		upperRight := "abcdefg"
		bottomLeft := "abcdefg"
		bottomRight := "abcdefg"

		for _, o := range inputarr {
			for _, s := range uniqueSegLens {
				if len(o) == s {
					if len(o) == sevenLen {
						top = Intersection(o, top)
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
					} else if len(o) == oneLen {
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
					} else if len(o) == fourLen {
						upperLeft = Intersection(o, upperLeft)
						middle = Intersection(o, middle)
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
					} else if len(o) == eightLen {
						upperLeft = Intersection(o, upperLeft)
						middle = Intersection(o, middle)
						top = Intersection(o, top)
						bottom = Intersection(o, bottom)
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
						bottomLeft = Intersection(o, bottomLeft)
					} else if len(o) == sixOrZeroOrNine {
						bottomRight = Intersection(o, bottomRight)
						top = Intersection(o, top)
						bottom = Intersection(o, bottom)
						upperLeft = Intersection(o, upperLeft)
					} else if len(o) == twoThreeOrFive {
						top = Intersection(o, top)
						bottom = Intersection(o, bottom)
						middle = Intersection(o, middle)
					}
				}
			}
		}
		for _, o := range outputs[bigIdx] {
			for _, s := range uniqueSegLens {
				if len(o) == s {
					if len(o) == sevenLen {
						top = Intersection(o, top)
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
					} else if len(o) == oneLen {
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
					} else if len(o) == fourLen {
						upperLeft = Intersection(o, upperLeft)
						middle = Intersection(o, middle)
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
					} else if len(o) == eightLen {
						upperLeft = Intersection(o, upperLeft)
						middle = Intersection(o, middle)
						top = Intersection(o, top)
						bottom = Intersection(o, bottom)
						upperRight = Intersection(o, upperRight)
						bottomRight = Intersection(o, bottomRight)
						bottomLeft = Intersection(o, bottomLeft)
					} else if len(o) == sixOrZeroOrNine {
						bottomRight = Intersection(o, bottomRight)
						top = Intersection(o, top)
						bottom = Intersection(o, bottom)
						upperLeft = Intersection(o, upperLeft)
					} else if len(o) == twoThreeOrFive {
						top = Intersection(o, top)
						bottom = Intersection(o, bottom)
						middle = Intersection(o, middle)
					}
				}
			}
		}
		//                            top, middle, bottom, upleft, upright, downleft, downright
		//                            0,    1         2      3        4        5          6
		segmentDecode := []string{top, middle, bottom, upperLeft, upperRight, bottomLeft, bottomRight}
		for seg := range segmentDecode {
			for otherseg := range segmentDecode {
				if seg != otherseg {
					if len(segmentDecode[seg]) > len(segmentDecode[otherseg]) {
						segmentDecode[seg] = LeaveOnlyUnique(segmentDecode[otherseg], segmentDecode[seg])
					} else {
						segmentDecode[otherseg] = LeaveOnlyUnique(segmentDecode[seg], segmentDecode[otherseg])
					}
				}
			}
		}

		s := segmentDecode
		codes := []string{
			s[0] + s[2] + s[3] + s[4] + s[5] + s[6],        // 0
			s[6] + s[4],                                    // 1
			s[0] + s[1] + s[2] + s[4] + s[5],               // 2
			s[0] + s[1] + s[2] + s[4] + s[6],               // 3
			s[1] + s[3] + s[4] + s[6],                      // 4
			s[0] + s[1] + s[2] + s[3] + s[6],               // 5
			s[0] + s[1] + s[5] + s[6] + s[3] + s[2],        // 6
			s[0] + s[4] + s[6],                             // 7
			s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6], // 8
			s[0] + s[1] + s[2] + s[3] + s[4] + s[6],        // 9
		}

		decoder := make(map[string]int)
		for i, code := range codes {
			codearr := strings.Split(code, "")
			sort.Strings(codearr)
			decoder[strings.Join(codearr, "")] = i
		}

		fmt.Println("outputs", outputs[bigIdx])
		fmt.Println("decoder", decoder)
		num := 0
		for idx, o := range outputs[bigIdx] {
			oarr := strings.Split(o, "")
			sort.Strings(oarr)
			num += decoder[strings.Join(oarr, "")] * int(math.Pow10(3-idx))
		}
		fmt.Println(num)
		sum += num
	}

	fmt.Println(sum)

}

func Run() {
	_, inarr := utils.LoadFile("day08", "\n")
	// inarr := strings.Split(instr, ",")
	// inarrInt := []int{}
	// for i := range inarr {
	// 	num, _ := strconv.Atoi(strings.TrimSpace(inarr[i]))
	// 	inarrInt = append(inarrInt, num)
	// }

	firstOne(inarr)
	secondOne(inarr)
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

func Intersection(a string, b string) string {
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")
	sort.Strings(arrA)
	sort.Strings(arrB)
	output := ""
	for ai := range arrA {
		for ab := range arrB {
			if arrA[ai] == arrB[ab] {
				output += arrA[ai]
			}
		}
	}
	return output
}

func LeaveOnlyUnique(a string, b string) string {
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")
	sort.Strings(arrA)
	sort.Strings(arrB)

	// newA := ""
	newB := ""
	for _, bchar := range arrB {
		shared := false
		for _, achar := range arrA {
			if achar == bchar {
				shared = true
			}
		}
		if !shared {
			newB += bchar
		}
	}
	if newB == "" {
		newB = b
	}

	return newB
}
