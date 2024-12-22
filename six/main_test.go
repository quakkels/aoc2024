package main

import (
	"log"
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

func TestFindGuard(t *testing.T) {
	inputa := []string {
		"..",
		".>",
	}

	p1a := newPart1(inputa)
	resulta, dira := p1a.findStart()

	if resulta[0] != 1 || resulta[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resulta)
	}
	
	if dira != '>' {
		t.Fatal("Expected result: >. Actual:", string(dira))
	}

	inputb := []string {
		"..",
		".v",
	}

	p1b := newPart1(inputb)
	resultb, dirb := p1b.findStart()

	if resultb[0] != 1 || resultb[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resultb)
	}

	if dirb != 'v' {
		t.Fatal("Expected result: v. Actual:", string(dirb))
	}

	inputc := []string {
		"..",
		".<",
	}

	p1c := newPart1(inputc)
	resultc, dirc := p1c.findStart()

	if resultc[0] != 1 || resultc[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resultc)
	}

	if dirc != '<' {
		t.Fatal("Expected result: <. Actual:", string(dirc))
	}

	inputd := []string {
		"..",
		".^",
	}

	p1d := newPart1(inputd)
	resultd, dird := p1d.findStart()

	if resultd[0] != 1 || resultd[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resultd)
	}

	if dird != '^' {
		t.Fatal("Expected result: ^. Actual:", string(dird))
	}
}

func TestIsInbounds(t *testing.T) {
	input := []string {
		"..",
		"..",
	}

	p1 := newPart1(input)
	p1.currentPosition = []int{0,0}

	isInboundsA := p1.isInbounds()

	if !isInboundsA {
		t.Fatal("Expected: true. Actual:", isInboundsA)
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
	log.Println(input)

	p1 := newPart1(input)
	result := p1.part1()

	if result != 41 {
		t.Fatal("Expected result: 41, actual:", result)
	}
}

