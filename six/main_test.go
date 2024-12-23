package main

import (
	//"log"
	"testing"
)

func TestRightTurn(t *testing.T) {
	input := []string {
		">",
	}

	d6 := newDay6(input)

	if d6.currentDirection != '>' {
		t.Fatal("Expected result: >. Actual:", d6.currentDirection)
	}

	d6.currentDirection = d6.getRightTurn()
	if d6.currentDirection != 'v' {
		t.Fatal("Expected result: v. Actual:", d6.currentDirection)
	}

	d6.currentDirection = d6.getRightTurn()
	if d6.currentDirection != '<' {
		t.Fatal("Expected result: <. Actual:", d6.currentDirection)
	}

	d6.currentDirection = d6.getRightTurn()
	if d6.currentDirection != '^' {
		t.Fatal("Expected result: ^. Actual:", d6.currentDirection)
	}
}

func TestPart1(t *testing.T) {
	input:= []string {
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	d6 := newDay6(input)
	result := d6.part1()

	if result != 41 {
		t.Fatal("Expected result: 41, actual:", result)
	}
}

func TestPart1ExitUp(t *testing.T) {
	input:= []string {
		".",
		"^",
	}

	d6 := newDay6(input)
	result := d6.part1()

	if result != 2 {
		t.Fatal("Expected result: 2, actual:", result)
	}
}

func TestPart1ExitRight(t *testing.T) {
	input:= []string {
		">",
	}

	d6 := newDay6(input)
	result := d6.part1()

	if result != 1 {
		t.Fatal("Expected result: 1, actual:", result)
	}
}

func TestPart1ExitLeft(t *testing.T) {
	input:= []string {
		"<",
	}

	d6 := newDay6(input)
	result := d6.part1()

	if result != 1 {
		t.Fatal("Expected result: 1, actual:", result)
	}
}

func TestPart1ExitDown(t *testing.T) {
	input:= []string {
		"v",
		".",
	}

	d6 := newDay6(input)
	result := d6.part1()

	if result != 2 {
		t.Fatal("Expected result: 2, actual:", result)
	}
}

func TestPart1CrossStart(t *testing.T) {
	input:= []string {
		".#..",
		"...#",
		".^..",
		"..#.",
	}

	d6 := newDay6(input)
	result := d6.part1()

	if result != 5 {
		t.Fatal("Expected result: 5, actual:", result)
	}
}

