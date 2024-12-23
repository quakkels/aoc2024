package main

import (
	//"log"
	"testing"
)

func TestRightTurn(t *testing.T) {
	input := []string {
		">",
	}

	p1 := newPart1(input)

	if p1.currentDirection != '>' {
		t.Fatal("Expected result: >. Actual:", p1.currentDirection)
	}

	p1.currentDirection = p1.getRightTurn()
	if p1.currentDirection != 'v' {
		t.Fatal("Expected result: v. Actual:", p1.currentDirection)
	}

	p1.currentDirection = p1.getRightTurn()
	if p1.currentDirection != '<' {
		t.Fatal("Expected result: <. Actual:", p1.currentDirection)
	}

	p1.currentDirection = p1.getRightTurn()
	if p1.currentDirection != '^' {
		t.Fatal("Expected result: ^. Actual:", p1.currentDirection)
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

	p1 := newPart1(input)
	result := p1.part1()

	if result != 41 {
		t.Fatal("Expected result: 41, actual:", result)
	}
}

func TestPart1ExitUp(t *testing.T) {
	input:= []string {
		".",
		"^",
	}

	p1 := newPart1(input)
	result := p1.part1()

	if result != 2 {
		t.Fatal("Expected result: 2, actual:", result)
	}
}

func TestPart1ExitRight(t *testing.T) {
	input:= []string {
		">",
	}

	p1 := newPart1(input)
	result := p1.part1()

	if result != 1 {
		t.Fatal("Expected result: 1, actual:", result)
	}
}

func TestPart1ExitLeft(t *testing.T) {
	input:= []string {
		"<",
	}

	p1 := newPart1(input)
	result := p1.part1()

	if result != 1 {
		t.Fatal("Expected result: 1, actual:", result)
	}
}

func TestPart1ExitDown(t *testing.T) {
	input:= []string {
		"v",
		".",
	}

	p1 := newPart1(input)
	result := p1.part1()

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

	p1 := newPart1(input)
	result := p1.part1()

	if result != 5 {
		t.Fatal("Expected result: 5, actual:", result)
	}
}

