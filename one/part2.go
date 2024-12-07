package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

 
func main() {
	fmt.Println("Challenge One, Part Two")

	// parse input into left and right slices
	filename := "input.txt"
	// filename = "test.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var left []int
	var right []int
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")

		leftv, _ := strconv.Atoi(parts[0])
		left = append(left, leftv)

		rightv, _ := strconv.Atoi(parts[1])
		right = append(right, rightv)

	}

	// calculate diffs between left and right
	similarity := 0
	for i := 0; i < len(left); i++ {
		similarity += scoreSimilarity(left[i], right)
	}
	
	fmt.Println(similarity)
}

func scoreSimilarity(target int, list []int) int{
	count := 0
	for i := range list {
		if list[i] == target {
			count += 1
		}
	}

	score := target * count
	return score
}

