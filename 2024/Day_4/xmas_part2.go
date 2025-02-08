package main

import (
	"AoC/utils"
	"fmt"
	"strings"
)

func match_xmas(word string) bool {
	word = strings.ToUpper(word)
	return word == "XMAS"
}

/*
top-left-to-bottom-right
bottom-right-to-top-left
bottom-left-to-top-right
top-right-to-bottom-left
*/
func process_window(lines []string) int {
	var total_matches = 0
	var total_top_left_right = 0
	var total_bot_right_left = 0
	var total_top_right_left = 0
	var total_bot_left_right = 0

	//match 4 rows
	current_row := 0
	current_col := 0
	for i, line := range lines {

		current_row = i
		for j, _ := range line {

			current_col = j
			top_left_right := match_top_left_right_diagnol(current_row, current_col, lines)
			// total_top_left_right += top_left_right
			bot_right_left := match_bot_right_left_diagnol(current_row, current_col, lines)
			// total_bot_right_left += bot_right_left
			top_right_left := match_top_right_left_diagnol(current_row, current_col+2, lines)
			// total_top_right_left += top_right_left
			bot_left_right := match_bot_left_right_diagnol(current_row+2, current_col, lines)
			// total_bot_left_right += bot_left_right
			xmatch := (top_left_right || bot_right_left) && (top_right_left || bot_left_right)
			if xmatch {
				total_matches += 1
			}

		}

	}
	fmt.Printf("%d %d %d %d %d %d %d %d", total_top_left_right, total_bot_right_left, total_top_right_left, total_bot_left_right)
	return total_matches
}

func match_top_right_left_diagnol(current_row int, current_col int, lines []string) bool {
	var word = ""
	if len(lines)-current_row < 4 || current_col < 3 {
		return false
	}

	for i := range 4 {
		word += string(lines[current_row+i][current_col-i])
	}
	if match_xmas(string(word)) {
		return true
	}

	return false
}

func match_top_left_right_diagnol(current_row int, current_col int, lines []string) bool {
	var word = ""
	if len(lines)-current_row < 4 || len(lines[0])-current_col < 4 {
		return false
	}

	for i := range 4 {
		word += string(lines[current_row+i][current_col+i])
	}
	if match_xmas(string(word)) {
		return true
	}

	return false

}

func match_bot_right_left_diagnol(current_row int, current_col int, lines []string) bool {
	var word = ""
	if current_row < 3 || current_col < 3 {
		return false
	}

	for i := range 4 {
		word += string(lines[current_row-i][current_col-i])
	}

	if match_xmas(string(word)) {
		return true
	}
	return false
}

func match_bot_left_right_diagnol(current_row int, current_col int, lines []string) bool {
	var word = ""
	if current_row < 3 || len(lines[0])-current_col < 4 {
		return false
	}

	for i := range 4 {
		word += string(lines[current_row-i][current_col+i])
	}
	if match_xmas(string(word)) {
		return true
	}
	return false
}

func main() {
	lines := utils.ReadFile("")
	process_window(lines)
}
