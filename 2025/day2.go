package main

import (
	"fmt"
	"strconv"
	"strings"
)

// isInvalidProductID checks whether an ID is invalid due to being composed of any sequence of digits repeated at least twice.
func isInvalidProductID(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)

	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}

		chunk := idStr[:size]
		repeated := strings.Repeat(chunk, length/size)
		if repeated == idStr {
			return true
		}
	}

	return false
}

// findInvalidIDs iterates through the provided ranges and finds all invalid IDs
func findInvalidIDs(ranges []string) []int {
	invalidIDs := []int{}

	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		if len(bounds) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(bounds[0])
		end, err2 := strconv.Atoi(bounds[1])
		if err1 != nil || err2 != nil {
			continue
		}

		for id := start; id <= end; id++ {
			if isInvalidProductID(id) {
				invalidIDs = append(invalidIDs, id)
			}
		}
	}

	return invalidIDs
}

func main() {
	// Updated input ranges
	input := "2157315-2351307,9277418835-9277548385,4316210399-4316270469,5108-10166,872858020-872881548,537939-575851,712-1001,326613-416466,53866-90153,907856-1011878,145-267,806649-874324,6161532344-6161720341,1-19,543444404-543597493,35316486-35418695,20-38,84775309-84908167,197736-309460,112892-187377,336-552,4789179-4964962,726183-793532,595834-656619,1838-3473,3529-5102,48-84,92914229-92940627,65847714-65945664,64090783-64286175,419838-474093,85-113,34939-52753,14849-30381"
	ranges := strings.Split(input, ",")

	invalidIDs := findInvalidIDs(ranges)
	total := 0
	for _, id := range invalidIDs {
		total += id
	}

	fmt.Printf("Updated Invalid IDs: %v\n", invalidIDs)
	fmt.Printf("Updated Sum of Invalid IDs: %d\n", total)
}
