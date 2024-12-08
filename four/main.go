package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 4")

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("missing input filepath.")
		return
	}
	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(">", text)
	}

	fmt.Println("Result>")
}
