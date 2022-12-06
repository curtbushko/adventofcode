package aoc

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(0)
}

func parseInput(in []byte) string {
	return string(in)
}

func readInput(filename string) ([]byte, error) {
	in, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Cannot read input file %s: %s", filename, err)
	}
	return in, nil
}
