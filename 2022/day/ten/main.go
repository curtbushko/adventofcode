package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")
}

func run(filename string) int {
	input, _ := os.ReadFile(filename)

	operations := makeOperations(string(input))

	total := processOperations(operations)

	return total
}

// Break the cycles down into an operation per cycle
func makeOperations(input string) []int {
	out := make([]int, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if parts[0] == "noop" {
			out = append(out, 0)
		} else {
			v, _ := strconv.Atoi(parts[1])
			out = append(out, 0, v)
		}
	}
	return out
}

func processOperations(ops []int) int {
	total := 0
	x := 1
	cycle := 0

	for op, value := range ops {
		cycle = op + 1
		pos := op % 40
		if pos == 0 { // really 40 but index 0
			fmt.Printf("\n")
		}
		if pos >= x-1 && pos <= x+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		if saveCycle(cycle) {
			total = total + (cycle * x)
		}
		x = x + value
	}

	return total
}

func saveCycle(c int) bool {
	if c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220 {
		return true
	}
	return false
}

// func runtwo(filename string) int {
// 	input, _ := os.Open(filename)
// 	defer input.Close()
// 	sc := bufio.NewScanner(input)
//
// 	var value int
// 	var signal int
// 	total := 0
// 	x := 1
// 	cycle := 1
// 	for sc.Scan() {
// 		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 || cycle == 260 {
// 			signal = x * cycle
// 			fmt.Println("x:", x, "cycle:", cycle, "signal:", signal)
// 			total = total + signal
// 		}
// 		if sc.Text() != "noop" {
// 			parts := strings.Split(sc.Text(), " ")
// 			value, _ = strconv.Atoi(parts[1])
// 			x = x + value
// 			// fmt.Println("x:", x)
// 			cycle = cycle + 2
// 		} else {
// 			cycle++
// 		}
// 	}
// 	fmt.Println("cycle:", cycle)
// 	return total
// }
