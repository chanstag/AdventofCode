package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string) []string {
	var lines []string
	fmt.Println("Processing...")

	fileIn, err := os.Open(filename)

	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(fileIn)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

/*
* https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
* if we just remove the element from the original slice,  then we are operating on the underlying
* array and we want to preserve the original array
 */
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// func main() {

// 	filename := os.Args[1]
// 	lines := readFile(filename)
// 	for _, line := range lines {
// 		fmt.Println(line)
// 	}
// }
