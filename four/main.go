package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 4")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("missing input filepath.")
		return
	}
	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, text)
	}

	result := countWordInstances(grid, "MAS")
	fmt.Println("Result>", result)
}

func isOutOfBounds(height, width, i, j int) bool {
	if i < 0 || i >= height {
		return true
	}
	if j < 0 || j >= width {
		return true
	}
	return false
}

func countWordInstances(grid []string, word string) int {

	log.Println("Counting...")
	count := 0

	for i := range grid {
		for j := range grid[i] {
			foundBackSlash := checkLine(grid, word, i, j, 1, 1) // direction south east
			if foundBackSlash {
				foundForwardSlash := checkLine(grid, word, i, j + 2, 1, -1) // direction south west
				if foundForwardSlash {
					count++
				}
				foundForwardSlash = checkLine(grid, word, i+2, j, -1, 1) // direction north east
				if foundForwardSlash {
					count++
				}
			}

			foundBackSlash = checkLine(grid, word, i, j, -1, -1) // direction north west
			if foundBackSlash {
				foundForwardSlash := checkLine(grid, word, i-2, j, 1, -1) // direction south west
				if foundForwardSlash {
					count++
				}
				foundForwardSlash = checkLine(grid, word, i, j-2, -1, 1) // direction north east
				if foundForwardSlash {
					count++
				}
			}
		}
	}
	return count
}

func checkLine(grid []string, word string, i, j, directionI, directionJ int) bool {
	height := len(grid)
	width := len(grid[0])
	offsetI := 0
	offsetJ := 0
	foundMatch := true
	for l := range word {
		var currentI, currentJ = i + offsetI, j + offsetJ
		if isOutOfBounds(height, width, currentI, currentJ) {
			foundMatch = false
			break
		}
		if grid[currentI][currentJ] != word[l] {
			foundMatch = false
			break
		}
		offsetI += directionI
		offsetJ += directionJ
	}
	return foundMatch
}
