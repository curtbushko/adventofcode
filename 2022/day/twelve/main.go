package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	above int = -1
	below int = 1
	right int = 1
	left  int = -1
	steps int = 1
)

type Cell struct {
	X, Y     int
	Rune     rune
	Walkable bool // I feel that this will be needed for part 2 (ie walls)
}

type Grid struct {
	Data          [][]*Cell
	Width, Height int
	Start, End    Cell
}

// A cell node for building a tree
type Node struct {
	X, Y     int
	Rune     rune
	children *[]Node
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

func createGrid(filename string) *Grid {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	g := &Grid{Width: 0, Height: 0}
	y := 0
	for {
		line, err := rd.ReadString('\n')
		if err != nil && err == io.EOF {
			// last line was the previous line

			break
		}

		g.Data = append(g.Data, []*Cell{}) // Create our row(y) cells
		for x, r := range strings.TrimSuffix(line, "\n") {
			g.Data[y] = append(g.Data[y], &Cell{X: x, Y: y, Walkable: true, Rune: r})
			switch g.Data[y][x].Rune {
			case 'S':
				g.Start.X = x
				g.Start.Y = y
			case 'E':
				g.End.X = x
				g.End.Y = y
			}
		}
		y++
		g.Width = len(line) - 1
	}
	g.Height = y

	return g
}

func createNodes(g Grid) *Node {
	n := &Node{
		X: g.Start.X,
		Y: g.Start.Y,
	}

	end := g.Data[g.End.Y][g.End.X]
	for { // We will stop when we can no longer make nodes
	}

	return n
}

func (g *Grid) Print() {
	fmt.Println("-----", g.Height, g.Width)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			fmt.Printf("%c", g.Data[y][x].Rune)
		}
		fmt.Printf("\n")
	}
}
