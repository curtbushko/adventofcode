package main

import (
	"bufio"
	"fmt"
	"os"
)

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
