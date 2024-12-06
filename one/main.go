package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

 
func main() {
	fmt.Println("Challenge One")

	// parse input into left and right slices
	filename := "input.txt"
	//filename = "test.txt"
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

	slices.Sort(left)
	slices.Sort(right)

	// calculate diffs between left and right
	sum := 0
	for i := 0; i < len(left); i++ {
		if left[i] <= right[i] {
			sum += right[i] - left[i]
		} else {
			sum += left[i] - right[i]
		}
	}
	
	fmt.Println(sum)
}
