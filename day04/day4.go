package day04

import (
	"fmt"
	"roxerg_aoc/utils"
	"strconv"
	"strings"
)

func firstOne(inarr []string) {

	calls := strings.Split(inarr[0], ",")
	restofinput := inarr[2:]
	checkbox := "x"

	var bingoboards = [][][]string{}

	idx := 0
	for idx < len(restofinput) {

		if len(restofinput[idx]) < 2 {
			idx += 1
			continue
		}
		var board = [][]string{}
		for i := 0; i < 5; i++ {
			cleaned := strings.TrimSuffix(restofinput[idx], "\r\n")
			cleaned_arr := strings.Split(cleaned, " ")
			toRemoveIdx := []int{}

			for n := range cleaned_arr {
				if cleaned_arr[n] == "" {
					toRemoveIdx = append(toRemoveIdx, n)
				}
				// fmt.Println("before clean", cleaned_arr[n])
				cleaned_arr[n] = strings.TrimSpace(cleaned_arr[n])
				// fmt.Println("after", cleaned_arr[n])
			}

			fmt.Println(cleaned_arr, toRemoveIdx)
			temp_cleaned_arr := []string{}
			for n := range cleaned_arr {
				if cleaned_arr[n] != "" {
					temp_cleaned_arr = append(temp_cleaned_arr, cleaned_arr[n])
				}
			}
			cleaned_arr = temp_cleaned_arr
			fmt.Println(cleaned_arr)

			board = append(board, cleaned_arr)
			idx += 1
		}
		bingoboards = append(bingoboards, board)
		// idx += 1

	}

	for idx := range calls {
		call := calls[idx]
		win := false
		win_board := 0
		for board_idx := range bingoboards {
			for row_idx := range bingoboards[board_idx] {
				// mark
				for val_idx := range bingoboards[board_idx][row_idx] {
					if bingoboards[board_idx][row_idx][val_idx] == call {
						bingoboards[board_idx][row_idx][val_idx] = checkbox
					}
				}
			}
		}

		// check win
		for board_idx := range bingoboards {
			cols_to_check := []int{}

			// check cols

			for row_idx := range bingoboards[board_idx] {
				win_str := ""
				row_win := true
				for val_idx := range bingoboards[board_idx][row_idx] {
					if row_idx == 0 && bingoboards[board_idx][row_idx][val_idx] == checkbox {
						cols_to_check = append(cols_to_check, val_idx)
					}
					if bingoboards[board_idx][row_idx][val_idx] != checkbox {
						row_win = false
						break
					}
					win_str += bingoboards[board_idx][row_idx][val_idx]
				}
				if row_win {
					win_board = board_idx
					win = row_win
					break
				}
			}

			for col_idx := range bingoboards[1][1] {
				win_str := ""
				// fmt.Println(col_idx)
				col_win := true
				for row_idx := range bingoboards[board_idx] {
					if bingoboards[board_idx][row_idx][col_idx] != checkbox {
						col_win = false
						break
					} else {
						win_str += bingoboards[board_idx][row_idx][col_idx]
					}
				}
				if col_win {
					win_board = board_idx
					win = col_win
					break
				}

			}

		}
		// calculate board if win
		if win {
			sum := 0
			for row_idx := range bingoboards[win_board] {
				fmt.Println(bingoboards[win_board][row_idx])
				for val_idx := range bingoboards[win_board][row_idx] {
					if bingoboards[win_board][row_idx][val_idx] != checkbox {
						addition, _ := strconv.Atoi(bingoboards[win_board][row_idx][val_idx])
						fmt.Println(bingoboards[win_board][row_idx][val_idx], "==", addition)
						fmt.Println(bingoboards[win_board][row_idx][val_idx] == "")
						sum += addition
					}
				}
			}
			call_num, _ := strconv.Atoi(call)
			fmt.Println(call_num)
			fmt.Println(call_num * sum)
			return
		}
	}
}

func secondOne(inarr []string) {
	calls := strings.Split(inarr[0], ",")
	restofinput := inarr[2:]
	checkbox := "x"
	winrars := []int{}

	var bingoboards = [][][]string{}

	idx := 0
	for idx < len(restofinput) {

		if len(restofinput[idx]) < 2 {
			idx += 1
			continue
		}
		var board = [][]string{}
		for i := 0; i < 5; i++ {
			cleaned := strings.TrimSuffix(restofinput[idx], "\r\n")
			cleaned_arr := strings.Split(cleaned, " ")
			toRemoveIdx := []int{}

			for n := range cleaned_arr {
				if cleaned_arr[n] == "" {
					toRemoveIdx = append(toRemoveIdx, n)
				}
				// fmt.Println("before clean", cleaned_arr[n])
				cleaned_arr[n] = strings.TrimSpace(cleaned_arr[n])
				// fmt.Println("after", cleaned_arr[n])
			}

			temp_cleaned_arr := []string{}
			for n := range cleaned_arr {
				if cleaned_arr[n] != "" {
					temp_cleaned_arr = append(temp_cleaned_arr, cleaned_arr[n])
				}
			}
			cleaned_arr = temp_cleaned_arr

			board = append(board, cleaned_arr)
			idx += 1
		}
		bingoboards = append(bingoboards, board)
		// idx += 1

	}

	for idx := range calls {
		call := calls[idx]
		win := false
		win_board := 0
		for board_idx := range bingoboards {
			for row_idx := range bingoboards[board_idx] {
				// mark
				for val_idx := range bingoboards[board_idx][row_idx] {
					if bingoboards[board_idx][row_idx][val_idx] == call {
						bingoboards[board_idx][row_idx][val_idx] = checkbox
					}
				}
			}
		}

		// check win
		for board_idx := range bingoboards {
			for past_win_idx := range winrars {
				if winrars[past_win_idx] == board_idx {
					continue
				}
			}

			cols_to_check := []int{}

			// check cols

			for row_idx := range bingoboards[board_idx] {
				win_str := ""
				row_win := true
				for val_idx := range bingoboards[board_idx][row_idx] {
					if row_idx == 0 && bingoboards[board_idx][row_idx][val_idx] == checkbox {
						cols_to_check = append(cols_to_check, val_idx)
					}
					if bingoboards[board_idx][row_idx][val_idx] != checkbox {
						row_win = false
						break
					}
					win_str += bingoboards[board_idx][row_idx][val_idx]
				}
				if row_win {
					win_board = board_idx
					win = row_win
					break
				}
			}

			for col_idx := range bingoboards[1][1] {
				win_str := ""
				// fmt.Println(col_idx)
				col_win := true
				for row_idx := range bingoboards[board_idx] {
					if bingoboards[board_idx][row_idx][col_idx] != checkbox {
						col_win = false
						break
					} else {
						win_str += bingoboards[board_idx][row_idx][col_idx]
					}
				}
				if col_win {
					win_board = board_idx
					win = col_win
					break
				}

			}

			// calculate board if win

			new_win := true

			// fmt.Println(winrars, win_board)
			for past_win_idx := range winrars {
				if winrars[past_win_idx] == win_board {
					new_win = false
				}
			}

			if win && new_win {
				winrars = append(winrars, win_board)
				sum := 0
				for row_idx := range bingoboards[win_board] {
					//fmt.Println(bingoboards[win_board][row_idx])
					for val_idx := range bingoboards[win_board][row_idx] {
						if bingoboards[win_board][row_idx][val_idx] != checkbox {
							addition, _ := strconv.Atoi(bingoboards[win_board][row_idx][val_idx])
							sum += addition
						}
					}
				}
				call_num, _ := strconv.Atoi(call)
				fmt.Println(call_num)
				fmt.Println(call_num * sum)
				fmt.Println(len(winrars), "/", len(bingoboards))
			}

		}

	}

}

func Run() {
	_, inarr := utils.LoadFile("day04", "\n")
	//firstOne(inarr)
	secondOne(inarr)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
