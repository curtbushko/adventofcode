package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coords struct {
	x int
	y int
}

func main() {
	fmt.Println("hello world")
}

func run(filename string) int {
	input, _ := os.Open(filename)
	defer input.Close()
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	return 0
}

func createGrid(input string) []Coords {
	length := len(input)
	return
}
