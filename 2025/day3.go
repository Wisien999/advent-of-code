package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func improveJoltage(state []int, currDig int, start int) {
	i := start
	for i < len(state) && state[i] >= currDig {
		i++
	}

	if i >= len(state) {
		return
	}

	state[i] = currDig
	i++

	for ; i < len(state); i++ {
		state[i] = 0
	}
}

// findLargestJoltage finds the largest possible joltage from a single bank of batteries.
func findLargestJoltage(bank string) int {
	var maxDigits [12]int
	for i := 0; i < len(bank)-11; i++ {
		dig := int(bank[i] - '0')

		improveJoltage(maxDigits[:], dig, 0)
	}

	for i := len(bank) - 11; i < len(bank); i++ {
		dig := int(bank[i] - '0')

		improveJoltage(maxDigits[:], dig, i-len(bank)+12)
	}

	res := 0

	for i := 0; i < len(maxDigits); i++ {
		res = res*10 + maxDigits[i]
	}

	return res
}

func main() {
	// Open the input file
	file, err := os.Open("input3.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := 0
	largestJoltages := []int{}

	// Process each bank of batteries
	for scanner.Scan() {
		bank := strings.TrimSpace(scanner.Text())
		largest := findLargestJoltage(bank)
		largestJoltages = append(largestJoltages, largest)
		totalJoltage += largest
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		return
	}

	// Output the results
	fmt.Printf("Largest Joltages Per Bank: %v\n", largestJoltages)
	fmt.Printf("Total Output Joltage: %d\n", totalJoltage)
}
