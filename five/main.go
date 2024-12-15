package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("missing input filepath.")
		return
	}
	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	part1, part2 := part1and2(input)
	fmt.Println("five> part1:", part1)
	fmt.Println("five> part2:", part2)
}

type SafetyManualUpdates struct {
	rules [][]int
	updates [][]int
}

func NewSafetyManualUpdates(input []string) *SafetyManualUpdates {
	smu := &SafetyManualUpdates{}

	collectingUpdates := false
	for i := range input {
		if input[i] == "" {
			collectingUpdates = true
			continue
		}
		if collectingUpdates {
			parts := strings.Split(input[i], ",")
			var parsedUpdate []int
			for j := range parts {
				parsed, _ := strconv.Atoi(parts[j])
				parsedUpdate = append(parsedUpdate, parsed)
			}
			smu.updates = append(smu.updates, parsedUpdate)
		} else {
			parts := strings.Split(input[i], "|")
			prevPage, _ := strconv.Atoi(parts[0])
			nextPage, _ := strconv.Atoi(parts[1])
			smu.rules = append(smu.rules, []int{prevPage,nextPage})
		}
	}

	return smu
}

func (smu *SafetyManualUpdates) part1and2() (int, int) {
	validUpdates, invalidUpdates := smu.findValidUpdates()
	validResult := sumMiddlePages(validUpdates)

	smu.sortUpdates(invalidUpdates)
	correctedResults := sumMiddlePages(invalidUpdates)
	
	return validResult, correctedResults
}

func (smu *SafetyManualUpdates) findValidUpdates() ([][]int, [][]int) {
	var validUpdates [][]int
	var invalidUpdates [][]int

	for i := range smu.updates {
		isValidUpdate := true
		previousPages := make(map[int] bool)

		for j := range smu.updates[i] {
			isValidPage := true

			for k := range smu.rules {
				if smu.updates[i][j] == smu.rules[k][1] {
					// if update doesn't have the "previous" page number, then rule doesn't apply
					if !updateHasPage(smu.updates[i], smu.rules[k][0]) {
						continue
					}

					if _, ok := previousPages[smu.rules[k][0]]; !ok {
						isValidPage = false
						break
					}
				}
			}

			if !isValidPage {
				isValidUpdate = false
				break
			} else {
				previousPages[smu.updates[i][j]] = true
			}
		}

		if isValidUpdate {
			validUpdates = append(validUpdates, smu.updates[i])
		} else {
			invalidUpdates = append(invalidUpdates, smu.updates[i])
		}
	}

	return validUpdates, invalidUpdates
}

func (smu *SafetyManualUpdates) sortUpdates(needsSorting [][]int) {
	for i := range needsSorting {
		sort.Slice(needsSorting[i], func (right, left int) bool {
			// if returns false, then right should be before left
			for r := range smu.rules {
				if needsSorting[i][left] == smu.rules[r][0] && needsSorting[i][right] == smu.rules[r][1] {
					return false
				}
				if needsSorting[i][left] == smu.rules[r][1] && needsSorting[i][right] == smu.rules[r][0] {
					return true
				}
			}
			return true
		})
	}
}

func updateHasPage(update []int, page int) bool {
	for i := range update {
		if update[i] == page {
			return true
		}
	}
	return false
}

func sumMiddlePages(updateList [][]int) int {
	sum := 0
	for i := range updateList {
		length := len(updateList[i])
		middleIndex := length/2
		sum += updateList[i][middleIndex]
	}
	return sum
}

func part1and2(input []string) (int,int) {
	smu := NewSafetyManualUpdates(input)
	return smu.part1and2()
}

