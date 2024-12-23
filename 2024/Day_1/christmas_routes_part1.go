package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile() ([]int, []int) {
	var nums1 []int
	var nums2 []int
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
		fmt.Printf("num1=%d\n", col1_int)

		col2_int, err := strconv.Atoi(num_row[1])

		if err != nil {
			log.Fatal("Error converting string to int")
		}
		fmt.Printf("num2=%d\n", col2_int)

		nums1 = append(nums1, col1_int)
		nums2 = append(nums2, col2_int)
		// fmt.Printf("len(nums1)=%d and len(nums2)=%d", len(nums1), len(nums2))
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := 0; i < 1000; i++ {
		fmt.Printf("nums1[%d]=%d, nums2[%d]=%d\n", i, nums1[i], i, nums2[i])
		_, err := writer.WriteString(strconv.Itoa(nums1[i]) + " " + strconv.Itoa(nums2[i]) + "\n")

		if err != nil {
			log.Fatal("didn't write to numsout.txt")
		}
	}
	return nums1, nums2
}

// Absolute value for int
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var nums1 []int
	var nums2 []int
	var diff_list []int
	var diff int
	var sum int

	fmt.Printf("len(nums1)=%d, cap(nums1)=%d", len(nums1), cap(nums1))
	nums1, nums2 = readFile()
	fmt.Printf("len(nums1)=%d, cap(nums1)=%d", len(nums1), cap(nums1))

	for i := 0; i < len(nums1); i++ {
		diff = absInt(nums1[i] - nums2[i])
		// fmt.Println("diff: ", diff)
		diff_list = append(diff_list, diff)
		sum += diff
	}

	fmt.Println("sum=", sum)
	return
}
