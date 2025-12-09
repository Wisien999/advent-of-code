package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Parse the input
	var points [][2]int
	var ingredientIDs []int

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	isRangeSection := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Blank line separates the ranges and the available IDs
			isRangeSection = false
			continue
		}

		if isRangeSection {
			// Parse range line (e.g., "3-5")
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				fmt.Println("Invalid range format:", line)
				return
			}
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing range:", line)
				return
			}

			points = append(points, [2]int{start, 1})
			points = append(points, [2]int{end, -1})
		} else {
			// Parse ingredient ID line
			id, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error parsing ingredient ID:", line)
				return
			}
			ingredientIDs = append(ingredientIDs, id)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	var ranges [][2]int
	start := -1
	count := 0
	for _, x := range points {
		p := x[0]
		count += x[1]

		if count == 0 {
			ranges = append(ranges, [2]int{start, p})
		} else if count == 1 && x[1] == 1 {
			start = p
		}
	}

	freshCnt := 0

	for _, id := range ingredientIDs {
		_, found := sort.Find(len(ranges), func(i int) int {
			rang := ranges[i]

			start, end := rang[0], rang[1]

			if id < start {
				return -1
			}

			if end < id {
				return 1
			}

			return 0
		})

		if found {
			freshCnt++
		}

	}

	allFresh := 0
	for _, rang := range ranges {
		leng := rang[1] - rang[0] + 1
		allFresh += leng

	}

	fmt.Printf("available fresh ingradients = %d\n", freshCnt)
	fmt.Printf("all fresh ingradients = %d\n", allFresh)

}
