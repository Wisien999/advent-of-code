package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input grid from file
	grid, err := readInput("input4.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	ranges, ids, err := parseIngredientDatabase("2025/input4.txt")
	if err != nil {
		fmt.Println("Error parsing ingredient database:", err)
		return
	}

	freshCount := countFreshIngredients(ranges, ids)
	fmt.Println("Fresh ingredient count:", freshCount)
	totalRemoved := calculateTotalRemovableRolls(grid)
	fmt.Println("Total rolls removed:", totalRemoved)
	// Calculate accessible rolls
	accessibleCount := calculateAccessibleRolls(grid)

	// Print the result
	fmt.Println("Number of accessible rolls:", accessibleCount)
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func calculateAccessibleRolls(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	accessibleCount := 0

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '@' {
				adjacentCount := 0
				for _, dir := range directions {
					nr, nc := i+dir[0], j+dir[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
						adjacentCount++
					}
				}
				if adjacentCount < 4 {
					accessibleCount++
				}
			}
		}
	}
func parseIngredientDatabase(filename string) ([][2]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var ranges [][2]int
	var ids []int
	var isRangesSection bool = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isRangesSection = false
			continue
		}

func countFreshIngredients(ranges [][2]int, ids []int) int {
	freshCount := 0

	for _, id := range ids {
		isFresh := false
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshCount++
		}
	}

	return freshCount
}
		if isRangesSection {
			var start, end int
			fmt.Sscanf(line, "%d-%d", &start, &end)
			ranges = append(ranges, [2]int{start, end})
		} else {
			var id int
			fmt.Sscanf(line, "%d", &id)
			ids = append(ids, id)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return ranges, ids, nil
}

	return accessibleCount
}

func calculateTotalRemovableRolls(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalRemoved := 0
	for {
		toRemove := [][2]int{}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] == '@' {
					adjacentCount := 0
					for _, dir := range directions {
						nr, nc := i+dir[0], j+dir[1]
						if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '@' {
							adjacentCount++
						}
					}
					if adjacentCount < 4 {
						toRemove = append(toRemove, [2]int{i, j})
					}
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, pos := range toRemove {
			grid[pos[0]][pos[1]] = '.'
			totalRemoved++
		}
	}

	return totalRemoved
}
