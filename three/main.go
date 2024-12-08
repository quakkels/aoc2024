package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 3")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("missing input filepath.")
		return
	}
	filepath := args[0]

	fileData, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	text := string(fileData)
	tokenizer := NewTokenizer(text)
	var muls []MulExpression

	for {
		if tokenizer.ch == 0 { break }
		mul := tokenizer.nextMul()
		muls = append(muls, mul)
	}

	fmt.Println("We done:", EvalMuls(muls))
}

type MulExpression struct {
	LeftNumber int
	RightNumber int
}

type Tokenizer struct {
	input string
	ch byte
	position int // position of current char
	readPosition int // position of read in input (after current char)
}

func NewTokenizer(input string) *Tokenizer {
	tokenizer := &Tokenizer{input:input}
	tokenizer.readChar()
	return tokenizer
}

func (tokenizer *Tokenizer) readChar() {
	if tokenizer.readPosition >= len(tokenizer.input) {
		tokenizer.ch = 0 // end of input
	} else {
		tokenizer.ch = tokenizer.input[tokenizer.readPosition]
	}
	tokenizer.position = tokenizer.readPosition
	tokenizer.readPosition += 1
}

func (tokenizer *Tokenizer) nextMul() MulExpression {
	tokenizer.ignoreExtraChars()

	for tokenizer.ch != 0 {
		ok, mul := tokenizer.readMul();
		if ok {
			return mul
		}
		tokenizer.ignoreExtraChars()
	}

	tokenizer.readChar()
	var mul MulExpression
	return mul
}

func (tokenizer *Tokenizer) readMul() (bool, MulExpression) {
	var mul MulExpression
	if tokenizer.ch != 'm' {
		return false, mul 
	}
	tokenizer.readChar()

	if tokenizer.ch != 'u' {
		return false, mul 
	}
	tokenizer.readChar()

	if tokenizer.ch != 'l' {
		return false, mul 
	}
	tokenizer.readChar()

	if tokenizer.ch != '(' {
		return false, mul 
	}
	tokenizer.readChar()

	// grab left number
	if !isDigit(tokenizer.ch) {
		return false, mul
	}
	leftNumber := tokenizer.readNumber()
	if leftNumber > 999 {
		return false, mul
	}

	if tokenizer.ch != ',' {
		return false, mul 
	}
	tokenizer.readChar()

	// grab right number
	if !isDigit(tokenizer.ch) {
		return false, mul
	}
	rightNumber := tokenizer.readNumber()
	if rightNumber > 999 {
		return false, mul
	}

	if tokenizer.ch != ')' {
		return false, mul 
	}
	tokenizer.readChar()

	mul.LeftNumber = leftNumber
	mul.RightNumber = rightNumber
	return true, mul
}

func (tokenizer *Tokenizer) readNumber() int {
	position := tokenizer.position
	for isDigit(tokenizer.ch) {
		tokenizer.readChar()
	}

	number, err := strconv.Atoi(tokenizer.input[position:tokenizer.position])
	if err != nil {
		panic(err)
	}

	return number
}

func (tokenizer *Tokenizer) ignoreExtraChars() {
	for tokenizer.ch != 'm' {
		tokenizer.readChar()
		if tokenizer.ch == 0 { break }
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func EvalMuls(muls []MulExpression) int {
	result := 0
	for i := range muls {
		prod := muls[i].LeftNumber * muls[i].RightNumber
		result += prod
	}
	return result
}
