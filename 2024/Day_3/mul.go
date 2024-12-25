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

func parse_nums(matches []string){
	exp := `([0-9]+),([0-9]+)`
	r, err := regexp.Compile(exp)

	if err != nil {
		log.Fatal("Failed to compile regex")
	}

	for _, mul := range matches{
		match := r.FindAllStringSubmatch(mul, -1)
		num1, err := strconv.Atoi(match[1])

	}
}

func main() {
	matches := parse_exp()

}
