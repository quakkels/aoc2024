package main

import (
	//"log"
	"testing"
)

func TestFindGuard(t *testing.T) {
	inputa := []string {
		"..",
		".>",
	}

	p1a := newPart1(inputa)
	resulta := p1a.findGuard()

	if resulta[0] != 1 || resulta[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resulta)
	}

	inputb := []string {
		"..",
		".v",
	}

	p1b := newPart1(inputb)
	resultb := p1b.findGuard()

	if resultb[0] != 1 || resultb[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resultb)
	}

	inputc := []string {
		"..",
		".<",
	}

	p1c := newPart1(inputc)
	resultc := p1c.findGuard()

	if resultc[0] != 1 || resultc[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resultc)
	}

	inputd := []string {
		"..",
		".^",
	}

	p1d := newPart1(inputd)
	resultd := p1d.findGuard()

	if resultd[0] != 1 || resultd[1] != 1 {
		t.Fatal("Expected result: 1,1. Actual:", resultd)
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

