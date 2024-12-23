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

type PositionsHashSet struct {
	hashSet map[[2]int]struct{}
}

func newPositionsHashSet() PositionsHashSet {
	return PositionsHashSet{
		hashSet: make(map[[2]int]struct{}),
	}
}

func (hs *PositionsHashSet) Add(value [2]int) {
	hs.hashSet[value] = struct{}{}
}

func (hs *PositionsHashSet) Contains(value [2]int) bool {
	_, ok := hs.hashSet[value]
	return ok
}

func (hs *PositionsHashSet) Size() int {
	return len(hs.hashSet)
}

type part1 struct {
	input []string
	directions [4]byte
	currentDirection byte
	currentPosition [2]int
	vectors map[byte][2]int
	visitedPositions PositionsHashSet
}

func newPart1(input []string) part1{
	p1 := part1{}

	p1.directions = [4]byte {
		'>',
		'v',
		'<',
		'^',
	}

	p1.vectors = make(map[byte][2]int)
	p1.vectors['>'] = [2]int{0, 1} // east
	p1.vectors['v'] = [2]int{1, 0} // south
	p1.vectors['<'] = [2]int{0, -1} // west
	p1.vectors['^'] = [2]int{-1, 0} // north

	p1.input = input

	p1.currentPosition, p1.currentDirection = p1.findStart()

	p1.visitedPositions = newPositionsHashSet()
	p1.visitedPositions.Add(p1.currentPosition)

	return p1
}

func (p1 *part1) part1() int {
	for {
		// move one cell
		nexti := p1.currentPosition[0] + p1.vectors[p1.currentDirection][0]
		nextj := p1.currentPosition[1] + p1.vectors[p1.currentDirection][1]

		isObstacle, isOffMap := p1.isObstacle(nexti, nextj)
		if isObstacle {
			if isOffMap {
				break // off the edge of the world
			}
			// so turn and try again
			// the next cell is a wall
			p1.currentDirection = p1.getRightTurn()
			continue
		}

		// we move
		p1.currentPosition[0] = nexti
		p1.currentPosition[1] = nextj
		p1.visitedPositions.Add(p1.currentPosition)
	}

	return p1.visitedPositions.Size()
}

func (p1 *part1) findStart() ([2]int, byte) {
	for i := range p1.input {
		for j := range p1.input[i] {
			_, exists := p1.vectors[p1.input[i][j]]
			if exists {
				return [2]int {i, j}, p1.input[i][j]
			}
		}
	}

	return [2]int{-1,-1}, 0 
}

func (p1 *part1) isInbounds() bool {
	return len(p1.input) > p1.currentPosition[0] && len(p1.input[0]) > p1.currentPosition[1]
}

func (p1 *part1) getRightTurn() byte {
	// oof, what am I doing here?
	nextIndex := 0
	for i := range p1.directions {
		if p1.directions[i] == p1.currentDirection {
			nextIndex = i + 1
		}
	}

	if nextIndex >= len(p1.directions) {
		nextIndex = 0
	}
	return p1.directions[nextIndex]
}

func (p1 *part1) isObstacle(i, j int) (isObstacle, isOffMap bool) {
	if len(p1.input) <= i ||  i < 0 {
		return true, true
	}
	if len(p1.input[0]) <= j ||  j < 0 {
		return true, true
	}
	return p1.input[i][j] == '#', false
}

