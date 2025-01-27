package main

import (
	"AoC/utils"
)

func match_xmas(word string) bool {
	return word == "xmas"
}

/*

 */
/*
match 4 rows
match 4 columns
top-left-to-bottom-right
bottom-right-to-top-left
bottom-left-to-top-right
top-right-to-bottom-left
*/
func process_window() {
	lines := utils.ReadFile("xmas.txt")
	var total_matches = 0
	//match 4 rows
	current_row := 0
	current_col := 0
	for i, line := range lines {
		if len(lines)-i < 4 {
			break
		}
		current_col = 0
		for j, _ := range line {
			if len(line)-j < 4 {
				break
			}
			total_matches += match_row(current_row, current_col, lines) + match_col(current_row, current_col, lines) + match_top_left_right_diagnol(current_row, current_col, lines)

			current_col = j
		}
		current_row = i
	}
}
func match_top_left_right_diagnol(current_row int, current_col int, lines []string) int {
	var chars []rune
	if len(lines)-1-current_row < 4 || len(lines[0])-1-current_col < 4 {
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

	for i := range 4 {
		chars = append(chars, rune(lines[current_row+4-i][current_col+4-i]))
	}

	if match_xmas(string(chars)) {
		return 1
	}
	return 0
}

func match_bot_left_right_diagnol(current_row int, current_col int, lines []string) int {
	var chars []rune
	if current_row < 4 || len(lines[0])-1-current_col < 4 {
		return 0
	}

	for i := range 4 {
		chars = append(chars, rune(lines[current_row+4-i][current_col+4-i]))
	}
	if match_xmas(string(chars)) {
		return 1
	}
	return 0
}

func match_row(current_row int, current_col int, lines []string) int {
	chars := lines[current_row][current_col : current_col+4]

	if match_xmas(string(chars)) {
		return 1
	}
	return 0
}

func match_col(current_row int, current_col int, lines []string) int {
	chars := lines[current_row : current_row+4][current_col]

	if match_xmas(string(chars)) {
		return 1
	}
	return 0

}

func main() {
	process_window()
	// lines := []string{
	// 	"xmas",
	// 	"masx",
	// 	"asxm",
	// 	"sxma",
	// }
	// // chars := lines[0 : 0+4][1]
	// // fmt.Printf("%s", chars)
	// var chars []rune
	// for i := range 4 {
	// 	chars = append(chars, rune(lines[0+i][0+i]))
	// }
	// fmt.Printf("%s", string(chars))
	return
}
