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
