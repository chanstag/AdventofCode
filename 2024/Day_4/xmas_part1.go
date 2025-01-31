package main

import (
	"AoC/utils"
	"fmt"
)

func match_xmas(word string) bool {
	return word == "XMAS"
}

/*
match 4 rows
match 4 columns
top-left-to-bottom-right
bottom-right-to-top-left
bottom-left-to-top-right
top-right-to-bottom-left
*/
func process_window(lines []string) int {
	var total_matches = 0
	//match 4 rows
	current_row := 0
	current_col := 0
	for i, line := range lines {
		// if len(lines)-i < 4 {
		// 	break
		// }
		current_row = i
		for j, _ := range line {
			// if len(line) - j < 4 {
			// 	break
			// }
			current_col = j
			total_matches += match_row(current_row, current_col, lines) + match_col(current_row, current_col, lines) + match_top_left_right_diagnol(current_row, current_col, lines) + match_bot_right_left_diagnol(current_row, current_col, lines) + match_top_right_left_diagnol(current_row, current_col, lines) +
				match_bot_left_right_diagnol(current_row, current_col, lines)

		}

	}

	return total_matches
}

func match_top_right_left_diagnol(current_row int, current_col int, lines []string) int {
	var chars []rune
	if len(lines)-current_row < 4 || current_col < 3 {
		return 0
	}

	for i := range 4 {
		chars = append(chars, rune(lines[current_row+i][current_col-i]))
	}
	if match_xmas(string(chars)) {
		return 1
	}

	return 0
}

func match_top_left_right_diagnol(current_row int, current_col int, lines []string) int {
	var chars []rune
	if len(lines)-current_row < 4 || len(lines[0])-current_col < 4 {
		return 0
	}

	for i := range 4 {
		chars = append(chars, rune(lines[current_row+i][current_col+i]))
	}
	if match_xmas(string(chars)) {
		return 1
	}

	return 0

}

func match_bot_right_left_diagnol(current_row int, current_col int, lines []string) int {
	var chars []rune
	if current_row < 3 || current_col < 3 {
		return 0
	}

	for i := range 4 {
		chars = append(chars, rune(lines[current_row-i][current_col-i]))
	}

	if match_xmas(string(chars)) {
		return 1
	}
	return 0
}

func match_bot_left_right_diagnol(current_row int, current_col int, lines []string) int {
	var chars []rune
	if current_row < 4 || len(lines[0])-current_col < 4 {
		return 0
	}

	for i := range 4 {
		chars = append(chars, rune(lines[current_row-i][current_col+i]))
	}
	if match_xmas(string(chars)) {
		return 1
	}
	return 0
}

func match_row(current_row int, current_col int, lines []string) int {
	if current_col+4 > len(lines[0]) {
		return 0
	}

	chars := lines[current_row][current_col : current_col+4]

	if match_xmas(string(chars)) {
		return 1
	}
	return 0
}

func match_col(current_row int, current_col int, lines []string) int {
	// Check bounds first
	if current_row+4 > len(lines) {
		return 0
	}

	// Get the character from each row at current_col
	chars := ""
	for i := current_row; i < current_row+4; i++ {
		if current_col >= len(lines[i]) {
			return 0
		}
		chars += string(lines[i][current_col])
	}

	if match_xmas(chars) {
		return 1
	}
	return 0
}

func main() {
	lines := utils.ReadFile("xmas.txt")
	total_matches := process_window(lines)
	fmt.Printf("total matches: %d", total_matches)
}
