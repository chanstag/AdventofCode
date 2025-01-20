package main

import (
	"AoC/utils"
	"fmt"
)

func match_xmas(word string) bool {
	return word == "xmas"
}

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
			match_rows(current_row, current_col, lines)
			current_col = j
		}
		current_row = i
	}
}

func match_rows(current_row int, current_col int, lines []string) {
	// total_matches := 0
	current_window := lines[current_row : current_row+4]
	for i, row := range current_window {
		chars := row[current_col : current_col+4]
		fmt.Printf("index %d: %s ", i, chars)
	}

}

func main() {
	process_window()
}
