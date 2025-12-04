package main

import (
	"bufio"
	"fmt"
	"os"
)

func simulate_raotation(position int, direction string, rotation int) (int, int) {
	old_pos := position
	at_zero := rotation / 100
	rotation = rotation % 100

	if direction == "L" {
		rotation = -rotation
	}
	fmt.Printf("pos = %d, at_zero = %d, rotation = %d\n", position, at_zero, rotation)

	position += rotation

	if position >= 100 {
		position -= 100
		at_zero++
		fmt.Printf("corrected position >= 100, pos = %d\n", position)
	} else if position < 0 {
		position += 100
		if old_pos > 0 {
			at_zero++
		}
		fmt.Printf("corrected position < 0, pos = %d\n", position)
	} else if position == 0 {
		at_zero++
	}
	fmt.Printf("at_zero = %d, rotation = %d, pos = %d\n", at_zero, rotation, position)

	return position, at_zero
}

func main() {
	// Open the file
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	position := 50
	answer := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var direction string
		var rotation int
		_, err := fmt.Sscanf(line, "%1s%d", &direction, &rotation)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}

		var at_zero int
		position, at_zero = simulate_raotation(position, direction, rotation)

		answer += at_zero
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Print(answer)
}
