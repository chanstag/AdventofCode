package main

import (
	"log"
	"strconv"
	"strings"
)

/*
 */

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func absDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}
	return x - y
}

func safe_report(report []int) bool {
	/*
		The levels are either all increasing or all decreasing.
		Any two adjacent levels differ by at least one and at most three.
	*/
	var descending bool
	var orderset bool = false
	for i, _ := range report {
		if i == 0 {
			continue
		}
		if !orderset {
			if report[i] < report[i-1] {
				descending = true
			} else if report[i] > report[i-1] {
				descending = false
			} else {
				return false
			}
			orderset = true
		} else {
			if descending {
				if report[i] > report[i-1] {
					return false
				}
			} else {
				if report[i] < report[i-1] {
					return false
				}
			}
		}

		diff := absInt(report[i] - report[i-1])
		if diff < 1 || diff > 3 {
			return false
		}

	}
	return true
}

func parseNums(lines []string) [][]int {
	reports := make([][]int, 0)
	for _, line := range lines {
		report := make([]int, 0)
		strNumbers := strings.Fields(line)
		for _, num := range strNumbers {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err, " Cannot convert value: ", val)
			}
			report = append(report, val)

		}
		reports = append(reports, report)
	}
	return reports
}

// func main() {
// 	var count int = 0
// 	filename := os.Args[1]
// 	lines := utils.ReadFile(filename)
// 	reports := parseNums(lines)
// 	for _, report := range reports {
// 		fmt.Println(report)
// 		issafe := safe_report(report)
// 		if issafe {
// 			count += 1
// 		}
// 	}
// 	fmt.Println("Count: ", count)
// 	// fmt.Println(safe_report([]int{1, 2, 3, 7}))
// }
