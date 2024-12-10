package main

import (
	"log"
	"os"
	"testing"
)

func TestIsOutOfBounds(t *testing.T) {
	result := isOutOfBounds(1, 1, 1, 0)
	if !result {
		t.Fatal("bottom bound failure")
	}
	result = isOutOfBounds(1, 1, -1, 0)
	if !result {
		t.Fatal("top bound failure")
	}
	result = isOutOfBounds(1, 1, 0, 1)
	if !result {
		t.Fatal("right bound failure")
	}
	result = isOutOfBounds(1, 1, 0, -1)
	if !result {
		t.Fatal("left bound failure")
	}
	result = isOutOfBounds(1, 1, 0, 0)
	if result {
		t.Fatal("in bounds failure")
	}
}

func TestCounter(t *testing.T) {
	log.SetOutput(os.Stdout)

	var grid []string
	var count int
	grid = append(grid, "m..")
	grid = append(grid, ".a.")
	grid = append(grid, "..s")
	count = countWordInstances(grid, "mas")
	if count != 0 {
		t.Fatal("1> Fail to ignore single word")
	}

	grid = []string{}
	grid = append(grid, "m.m.m...m")
	grid = append(grid, ".a...a.a.")
	grid = append(grid, "s.s...s..")
	count = countWordInstances(grid, "mas")
	if count != 1 {
		t.Fatal("2> Fail to find X word. Expect: 1, Actual:", count)
	}

	grid = []string{}
	grid = append(grid, "m.s...sm..")
	grid = append(grid, ".a...a..a.")
	grid = append(grid, "m.s.m....s")
	count = countWordInstances(grid, "mas")
	if count != 1 {
		t.Fatal("3> Fail to find X word. Expect: 1, Actual:", count)
	}

	grid = []string{}
	grid = append(grid, "s.ms....m")
	grid = append(grid, ".a..a..a.")
	grid = append(grid, "s.m..ms..")
	count = countWordInstances(grid, "mas")
	if count != 1 {
		t.Fatal("4> Fail to find X word. Expect: 1, Actual:", count)
	}

	grid = []string{}
	grid = append(grid, "s.s...s..")
	grid = append(grid, ".a...a.a.")
	grid = append(grid, "m.m.m...m")
	count = countWordInstances(grid, "mas")
	if count != 1 {
		t.Fatal("5> Fail to find X word. Expect: 1, Actual:", count)
	}

	grid = []string{}
	grid = append(grid, ".M.S......")
	grid = append(grid, "..A..MSMS.")
	grid = append(grid, ".M.S.MAA..")
	grid = append(grid, "..A.ASMSM.")
	grid = append(grid, ".M.S.M....")
	grid = append(grid, "..........")
	grid = append(grid, "S.S.S.S.S.")
	grid = append(grid, ".A.A.A.A..")
	grid = append(grid, "M.M.M.M.M.")
	grid = append(grid, "..........")
	count = countWordInstances(grid, "MAS")
	if count != 9 {
		t.Fatal("6> Fail to find 9 X's. Expect: 9, Actual:", count)
	}
}
