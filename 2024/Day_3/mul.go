package main

import (
	"AoC/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func parse_exp() []string {
	var matches []string
	exp := `mul\([0-9]+,[0-9]+\)`
	r, err := regexp.Compile(exp)

	if err != nil {
		log.Fatal("Failed to compile regex")
	}

	lines := utils.ReadFile("memory_file.txt")

	for _, line := range lines {
		match := r.FindAllString(line, -1)
		matches = append(matches, match...)
		fmt.Println(matches)
		fmt.Println(len(matches))
	}
	return matches
}

func parse_nums(matches []string) [][]int {
	var nums [][]int
	exp := `([0-9]+),([0-9]+)`
	r, err := regexp.Compile(exp)

	if err != nil {
		log.Fatal("Failed to compile regex")
	}

	for _, mul := range matches {
		match := r.FindStringSubmatch(mul)
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("issue converting string")
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal("issue converting string")
		}
		nums = append(nums, []int{num1, num2})
	}
	return nums
}

func multiply(nums [][]int) int {
	var total int = 0
	for _, num := range nums {
		total += num[0] * num[1]
	}
	return total
}

func main() {
	matches := parse_exp()
	nums := parse_nums(matches)

	fmt.Println(nums)
}
