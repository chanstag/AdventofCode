package main

import (
	"AoC/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	// "github.com/dlclark/regexp2"
)

//do\(\)((?:.*?)(mul\([0-9]+,[0-9]+\)))(?=.*don't\(\))

// (do\(\).*?don't\(\))
// func match_start() []string {
// 	var matches []string
// 	exp := `(^.*?don't\(\))`
// 	r, err := regexp.Compile(exp)

// 	if err != nil {
// 		log.Fatal("Failed to compile regex")
// 	}
// 	lines := utils.ReadFile("memory_file.txt")
// 	fmt.Println("Number of lines in memory.txt: ", len(lines))
// 	for _, line := range lines {
// 		match := r.FindAllString(line, -1)
// 		matches = append(matches, match...)

// 	}
// 	return matches
// }

// func parse_exp_2(valid_matches []string) []string {
// 	var matches []string
// 	exp := `mul\([0-9]+,[0-9]+\)`
// 	r, err := regexp.Compile(exp)

// 	if err != nil {
// 		log.Fatal("Failed to compile regex")
// 	}

// 	// lines := utils.ReadFile("memory_file.txt")

// 	for i, line := range valid_matches {
// 		match := r.FindAllString(line, -1)
// 		matches = append(matches, match...)
// 		fmt.Println("Match ", i, ": ", match)
// 		fmt.Println(len(match))
// 	}
// 	return matches
// }

// func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
// 	var matches []string
// 	enabled := true
// 	mult_exp := `mul\([0-9]+,[0-9]+\)`
// 	group_exp := `([0-9]+),([0-9]+)`
// 	mult_r, err := regexp.Compile(mult_exp)
// 	group_r, err := regexp.Compile(group_exp)

// 	if err != nil {
// 		log.Fatal("Failed to compile regex")
// 	}

// 	m, _ := re.FindStringMatch(s)

// 	for m != nil {
// 		if enabled{

// 			mult := mult_r.FindString(m.String())
// 			if mult != ""{
// 				group_match := group_r.FindAllStringSubmatch(mult, -1)
// 				matches = append(matches, group_match.Groups()[1].String())
// 			}
// )
// 		}

// 		m, _ = re.FindNextMatch(m)
// 	}
// 	return matches
// }

// func get_enabled() []string {
// 	var matches []string
// 	exp := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
// 	r := regexp2.MustCompile(exp, 0)

// 	lines := utils.ReadFile("memory_file.txt")

// 	for _, line := range lines {
// 		match := regexp2FindAllString(r, line)
// 		matches = append(matches, match...)
// 	}
// 	fmt.Println("matches: ", matches)

// 	return matches
// }

// func parse_nums_2(matches []string) [][]int {
// 	var nums [][]int
// 	exp := `([0-9]+),([0-9]+)`
// 	r, err := regexp.Compile(exp)

// 	if err != nil {
// 		log.Fatal("Failed to compile regex")
// 	}

// 	for _, mul := range matches {
// 		match := r.FindStringSubmatch(mul)
// 		num1, err := strconv.Atoi(match[1])
// 		if err != nil {
// 			log.Fatal("issue converting string")
// 		}
// 		num2, err := strconv.Atoi(match[2])
// 		if err != nil {
// 			log.Fatal("issue converting string")
// 		}
// 		nums = append(nums, []int{num1, num2})
// 	}
// 	return nums
// }

// func parse() int {
// 	matches := get_enabled()
// 	fmt.Println(matches)
// 	matches = append(matches, match_start()...)
// 	matches = parse_exp_2(matches)
// 	nums := parse_nums_2(matches)
// 	return multiply(nums)
// }

func parse() int {
	exp := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	mul_p := `mul\(\d+,\d+\)`
	enable_p := `do\(\)`
	disable_p := `don't\(\)`

	var total int = 0

	r, err := regexp.Compile(exp)
	r_mul_pattern, _ := regexp.Compile(mul_p)
	r_disabled_pattern, _ := regexp.Compile(disable_p)
	r_enabled_pattern, _ := regexp.Compile(enable_p)
	if err != nil {
		log.Fatal("Can't compile regex")
	}

	lines := utils.ReadFile("memory_file.txt")
	lines = []string{strings.Join(lines[:], "")}
	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		enabled := true
		for _, match := range matches {

			if enabled {
				matched := r_mul_pattern.FindStringIndex(match)
				if matched != nil {
					exp := `([0-9]+),([0-9]+)`
					r_nums, err := regexp.Compile(exp)

					if err != nil {
						log.Fatal("Failed to compile regex")
					}

					match := r_nums.FindStringSubmatch(match)
					num1, err := strconv.Atoi(match[1])
					if err != nil {
						log.Fatal("issue converting string")
					}
					num2, err := strconv.Atoi(match[2])
					if err != nil {
						log.Fatal("issue converting string")
					}

					total += num1 * num2
				}
				matched = r_disabled_pattern.FindStringIndex(match)
				if matched != nil {
					enabled = false
				}
			} else {
				matched := r_enabled_pattern.FindStringIndex(match)
				if matched != nil {
					enabled = true
				}
			}
		}
	}
	return total
}

func multiply(nums [][]int) int {
	var total int = 0
	for _, num := range nums {
		total += num[0] * num[1]
	}
	return total
}

func main() {
	//This gets the correct answer
	total := parse()
	fmt.Println("Total: ", total)
}
