package main

import (
	"log"
	"testing"
)

func TestPart1(t *testing.T) {
	input:= []string {
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}

	result, _ := part1and2(input)

	if result != 143 {
		t.Fatal("Expected result: 143, actual:", result)
	}
}

func TestFindValidUpdates(t *testing.T) {
	input:= []string {
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29", // valid
		"97,61,53,29,13", // valid
		"75,29,13", // valid
		"75,97,47,61,53", // invalid
		"61,13,29", // invalid
		"97,13,75,29,47", // invalid
	}

	smu := NewSafetyManualUpdates(input)

	// act
	result, result2 := smu.findValidUpdates()

	if len(result) != 3 {
		t.Fatal("Expected result: 3, actual:", result)
	}

	if len(result2) != 3 {
		t.Fatal("Expected result: 3, actual:", result2)
	}

	// act SORT
	log.Println(result2)
	smu.sortUpdates(result2)
	log.Println(result2)

	smu.updates = result2
	valid, invalid := smu.findValidUpdates()
	if len(valid) != 3 {
		t.Fatal("Expected result: 3, actual:", len(valid))
	}
	if len(invalid) != 0 {
		t.Fatal("Expected result: 0, actual:", len(invalid))
	}
}

func TestSafetyManualUpdates(t *testing.T) {
	input := []string {
		"12|13",
		"8|12",
		"1|2",
		"",
		"12,1,13", // valid... rule 8|12 doesn't apply because 8 isn't in the update
		"1,13,12", // invalid... rule 12|13 applies because 12 IS present in the update
		"1,2,5,8,12,13", // valid... all rules are followed
	}

	smu := NewSafetyManualUpdates(input)

	if len(smu.rules) != 3 {
		t.Fatal("Expected rule count: 3. Actual:", len(smu.rules))
	}

	if len(smu.updates) != 3 {
		t.Fatal("Expected update count: 3. Actual:", len(smu.updates))
	}

	validUpdates, invalidUpdates := smu.findValidUpdates()
	if len(validUpdates) != 2 {
		t.Fatal("Expected valid update count: 2. Actual:", len(validUpdates))
	}

	if len(invalidUpdates) != 1 {
		t.Fatal("Expected invalid update count: 1. Actual:", len(invalidUpdates))
	}
}

func TestSumMiddlePages(t *testing.T) {
	updates := [][]int{
		{1,2,3},
		{1,2,3,4,5}}

	result := sumMiddlePages(updates)

	if result != 5 {
		t.Fatal("sumMiddlePages expected: 5. Actual:", result) 
	}
}
