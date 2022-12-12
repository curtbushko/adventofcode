package main

import (
	"bufio"
	"fmt"
	"os"
)

var Grid = make([][]int, 0)
var start = make([][]int, 0)
var end = make([][]int, 0)

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

func createGrid(filename string) []Coords {
	length := len(input)


	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("Could not open file %s", filename)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	for {

		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// last line was the previous line
				break
			}
		}

	return
}
