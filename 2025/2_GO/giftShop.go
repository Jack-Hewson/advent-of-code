package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID(n int) bool {
	s := strconv.Itoa(n)

	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

func main() {
	// example_ids := []string{
	// 	"11-22",
	// 	"95-115",
	// 	"998-1012",
	// 	"1188511880-1188511890",
	// 	"222220-222224",
	// 	"1698522-1698528",
	// 	"446443-446449",
	// 	"38593856-38593862",
	// 	"565653-565659",
	// 	"824824821-824824827",
	// 	"2121212118-2121212124",
	// }
	file, _ := os.ReadFile("input.txt")

	fileContentIds := strings.Split(strings.TrimSpace(string(file)), ",")

	invalid_ids := 0
	for _, id := range fileContentIds {
		idRange := strings.Split(id, "-")
		min, _ := strconv.Atoi(idRange[0])
		max, _ := strconv.Atoi(idRange[1])

		for i := min; i <= max; i++ {
			if isInvalidID(i) {
				invalid_ids += i
			}
		}
	}

	fmt.Println("Total:", invalid_ids)
}
