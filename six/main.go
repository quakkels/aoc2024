package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 6")

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

	p1 := newPart1(input)
	result1 := p1.part1()
	fmt.Println("part1>", result1)
}

type part1 struct {
	input []string
	turns map[byte][]int
}

func newPart1(input []string) part1{
	p1 := part1{}
	p1.turns = make(map[byte][]int)
	p1.turns['^'] = []int{0, 1} // turn east
	p1.turns['>'] = []int{1, 0} // turn south
	p1.turns['v'] = []int{0, -1} // turn west
	p1.turns['<'] = []int{-1, 0} // turn north

	p1.input = input

	return p1
}

func (p1 *part1) part1() int {
	result := 0
	return result
}

func (p1 *part1) findGuard() []int {
	for i := range p1.input {
		for j := range p1.input[i] {
			_, exists := p1.turns[p1.input[i][j]]
			if exists {
				return []int {i, j}
			}
		}
	}

	return []int{-1,-1}
}
