package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions = [][]int{
	// i  j
	{-1, -1}, // up left
	// {-1, 0},  // up center
	{-1, 1}, // up right
	// {0, -1},  // center left
	// {0, 1},   // center right
	{1, -1}, // bottom left
	// {1, 0},   // bottom center
	{1, 1},   // bottom right
}

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

	result := countWordInstances(grid, "XMAS")
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
	height := len(grid)
	width := len(grid[0])
	count := 0

	for i := range grid {
		for j := range grid[i] {
			//log.Println("i:", i, " j:", j)
			foundMatch := false
			for k := range directions {
				offsetI := 0
				offsetJ := 0
				foundMatch = true
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
					offsetI += directions[k][0]
					offsetJ += directions[k][1]
				}

				if foundMatch {
					count += 1
				}
			}
		}
	}
	return count
}
