package main

import (
	"AoC/utils"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func recurse_safe_report(report []int, count int) bool {
	if count > 1 {
		return false
	}
	ret := safe_report_part2(report, count+1)
	return ret
}

func safe_report_part2(report []int, count int) bool {
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
				return recurse_safe_report(utils.RemoveIndex(report, i), count+1) || recurse_safe_report(utils.RemoveIndex(report, i-1), count+1)
			}
			orderset = true
		} else {
			if descending {
				if report[i] > report[i-1] {
					return recurse_safe_report(utils.RemoveIndex(report, i), count+1) || recurse_safe_report(utils.RemoveIndex(report, i-1), count+1)
				}
			} else {
				if report[i] < report[i-1] {
					return recurse_safe_report(utils.RemoveIndex(report, i), count+1) || recurse_safe_report(utils.RemoveIndex(report, i-1), count+1)
				}
			}
		}

		diff := absInt(report[i] - report[i-1])
		if diff < 1 || diff > 3 {
			return recurse_safe_report(utils.RemoveIndex(report, i), count+1) || recurse_safe_report(utils.RemoveIndex(report, i-1), count+1)
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

func main() {
	var safe_count int = 0
	var recurse_count int = 0
	filename := os.Args[1]
	lines := utils.ReadFile(filename)
	reports := parseNums(lines)
	for _, report := range reports {
		issafe := safe_report_part2(report, recurse_count)
		fmt.Println(report)
		fmt.Println(issafe)
		if issafe {
			safe_count += 1
		}
	}
	fmt.Println("Count: ", safe_count)
	// fmt.Println(safe_report_part2([]int{25, 22, 19, 21, 20, 17, 14, 13}, count))

	// fmt.Println(safe_report_part2([]int{30, 27, 29, 26, 28}, count))
	// fmt.Println(safe_report_part2([]int{86, 87, 88, 91, 91}, count))

}
