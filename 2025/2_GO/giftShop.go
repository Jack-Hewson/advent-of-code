package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	invalid_ids := 0

	file, _ := os.ReadFile("input.txt")
	fileContents := strings.TrimSpace((string(file)))
	fileContentIds := strings.SplitSeq(fileContents, ",")

	for val := range fileContentIds {
		idRange := strings.Split(val, `-`)
		min, _ := strconv.Atoi(idRange[0])
		max, _ := strconv.Atoi(idRange[1])

		for i := min; i <= max; i++ {
			iString := strconv.Itoa(i)
			middle := len(iString) / 2
			firstHalf := iString[:middle]
			secondHalf := iString[middle:]

			if firstHalf == secondHalf {
				fmt.Println(i)
				invalid_ids += i
			}
		}
	}

	fmt.Println("Total:", invalid_ids)
}
