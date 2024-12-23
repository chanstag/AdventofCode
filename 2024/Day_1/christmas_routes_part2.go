package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile() (map[int]int, map[int]int) {
	nums1 := make(map[int]int)
	nums2 := make(map[int]int)
	fmt.Println("Processing...")

	fileIn, err := os.Open("numsin.txt")

	if err != nil {
		log.Fatal()
	}

	fileOut, err := os.Create("numsout.txt")

	if err != nil {
		log.Fatal()
	}
	defer fileOut.Close()

	scanner := bufio.NewScanner(fileIn)
	writer := bufio.NewWriter(fileOut)
	defer writer.Flush()

	for scanner.Scan() {

		line := scanner.Text()
		// fmt.Printf("line=%s\n", line)
		num_row := strings.Fields(line)
		fmt.Printf("len(row)=%d, %s, %s", len(num_row), num_row[0], num_row[1])
		col1_int, err := strconv.Atoi(num_row[0])
		if err != nil {
			log.Fatal("Error converting string to int")
		}
		// fmt.Printf("num1=%d\n", col1_int)

		col2_int, err := strconv.Atoi(num_row[1])

		if err != nil {
			log.Fatal("Error converting string to int")
		}
		// fmt.Printf("num2=%d\n", col2_int)

		val, ok := nums1[col1_int]
		if !ok {
			nums1[col1_int] = 1
		} else {
			nums1[col1_int] = val + 1
		}

		val, ok = nums2[col2_int]

		if !ok {
			nums2[col2_int] = 1
		} else {
			nums2[col2_int] = val + 1
		}
	}

	return nums1, nums2
}

func main() {
	var nums1, nums2 map[int]int
	similarity_sum := 0
	nums1, nums2 = readFile()

	for key, _ := range nums1 {
		nums2_val, ok := nums2[key]
		if ok {
			similarity_sum += nums2_val * key
		}
	}
	fmt.Println("sum=%d", similarity_sum)
}
