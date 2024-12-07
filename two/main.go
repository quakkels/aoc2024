package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")

	/*
		test1 := "1 2 7 8 9"
		fmt.Println("test1", test1, isSafeReport(test1))
		//*/

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

	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		report := scanner.Text()
		if isSafeReport(report) {
			fmt.Println("safe", report)
			safeCount += 1
		} else {
			fmt.Println("unsafe", report)
		}
	}

	fmt.Println("Number of safe reports: ", safeCount)
}

func isSafeReport(report string) bool {
	parts := strings.Split(report, " ")
	var numbers []int
	for i := range parts {
		number, _ := strconv.Atoi(parts[i])
		numbers = append(numbers, number)
	}

	isDirectionUp := false
	for i := range numbers {
		if i != 0 {
			prev := numbers[i-1]
			curr := numbers[i]
			diff := 0

			if prev > curr {
				diff = prev - curr
			} else {
				diff = curr - prev
			}

			if diff < 1 || diff > 3 {
				// require change within bounds
				fmt.Println(i, "change not within bounds")
				return false
			}

			if i == 1 {
				fmt.Println("checking direction")
				isDirectionUp = prev < curr
			} else if isDirectionUp != (prev < curr) {
				// require consistent direction of change
				fmt.Println(i, "change direction not consistent", prev, curr)
				return false
			}
		}
	}
	fmt.Println("safe")
	return true
}
