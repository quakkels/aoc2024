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

	p1 := newDay6(input)
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

type day6 struct {
	input []string
	directions [4]byte
	currentDirection byte
	currentPosition [2]int
	vectors map[byte][2]int
	visitedPositions PositionsHashSet
}

func newDay6(input []string) day6{
	d6 := day6{}

	d6.directions = [4]byte {
		'>',
		'v',
		'<',
		'^',
	}

	d6.vectors = make(map[byte][2]int)
	d6.vectors['>'] = [2]int{0, 1} // east
	d6.vectors['v'] = [2]int{1, 0} // south
	d6.vectors['<'] = [2]int{0, -1} // west
	d6.vectors['^'] = [2]int{-1, 0} // north

	d6.input = input

	d6.currentPosition, d6.currentDirection = d6.findStart()

	d6.visitedPositions = newPositionsHashSet()
	d6.visitedPositions.Add(d6.currentPosition)

	return d6
}

func (d6 *day6) part1() int {
	for {
		// move one cell
		nexti := d6.currentPosition[0] + d6.vectors[d6.currentDirection][0]
		nextj := d6.currentPosition[1] + d6.vectors[d6.currentDirection][1]

		isObstacle, isOffMap := d6.isObstacle(nexti, nextj)
		if isObstacle {
			if isOffMap {
				break // off the edge of the world
			}
			// so turn and try again
			// the next cell is a wall
			d6.currentDirection = d6.getRightTurn()
			continue
		}

		// we move
		d6.currentPosition[0] = nexti
		d6.currentPosition[1] = nextj
		d6.visitedPositions.Add(d6.currentPosition)
	}

	return d6.visitedPositions.Size()
}

func (d6 *day6) findStart() ([2]int, byte) {
	for i := range d6.input {
		for j := range d6.input[i] {
			_, exists := d6.vectors[d6.input[i][j]]
			if exists {
				return [2]int {i, j}, d6.input[i][j]
			}
		}
	}

	return [2]int{-1,-1}, 0 
}

func (d6 *day6) isInbounds() bool {
	return len(d6.input) > d6.currentPosition[0] && len(d6.input[0]) > d6.currentPosition[1]
}

func (d6 *day6) getRightTurn() byte {
	// oof, what am I doing here?
	nextIndex := 0
	for i := range d6.directions {
		if d6.directions[i] == d6.currentDirection {
			nextIndex = i + 1
		}
	}

	if nextIndex >= len(d6.directions) {
		nextIndex = 0
	}
	return d6.directions[nextIndex]
}

func (d6 *day6) isObstacle(i, j int) (isObstacle, isOffMap bool) {
	if len(d6.input) <= i ||  i < 0 {
		return true, true
	}
	if len(d6.input[0]) <= j ||  j < 0 {
		return true, true
	}
	return d6.input[i][j] == '#', false
}

