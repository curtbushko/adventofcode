package eight

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Grid struct {
	previous []int
	current  []int
	next     []int
}

func ProcessGrid(input string) (total int, err error) {
	g := &Grid{
		previous: []int{},
	}

	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("Could not open file %s", input)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	cursor := 0
	for {

		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// last line was the previous line
				total = total + g.ProcessLastLine()
				break
			}
		}

		// line = next
		switch cursor {
		case 0:
			g.previous = make([]int, len(line))
			g.current = make([]int, len(line))
			g.next = linetoSlice(line)
			total = total + g.ProcessFirstLine()
		case 1:
			copy(g.current, g.next)
			g.next = linetoSlice(line)
		default:
			copy(g.previous, g.current)
			copy(g.current, g.next)
			g.next = linetoSlice(line)
			total = total + g.ProcessCurrentLine()
		}
		cursor++
	}
	return
}

func (g Grid) ProcessCurrentLine() int {
	var total int

	for i := range g.current {
		fmt.Println(g.previous[i], g.current[i], g.next[i])
		if i == 0 {
			total++
			continue
		}

		if i == (len(g.current) - 1) {
			total++
			break
		}
		// top
		for j = i; j > 0; j-- {
		}
		if g.current[i] >= g.previous[i] || // top
			g.current[i] >= g.next[i] || // below
			g.current[i] >= g.current[i-1] || // below
			g.current[i] >= g.current[i+1] { // below
			total++
		}

	}

	return total
}

func (g Grid) isVisible(cursor int) bool {
	// top
	for i := cursor; i > 0; i-- {
		if g.current[cursor] > g.previous[i] {
			return true
		}
	}

	// right
	for i := cursor; i < len(g.current)-1; i++ {
		if g.current[cursor] > g.current[i+1] {
			return true
		}
	}

	// left
	for i := cursor; i > 0; i-- {
		if g.current[cursor] > g.current[i-1] {
			return true
		}
	}

	// down
	for i := cursor; i > 0; i-- {
		if g.current[cursor] > g.current[i-1] {
			return true
		}
	}
}

func (g Grid) ProcessLastLine() int {
	return len(g.current)
}

func (g Grid) ProcessFirstLine() int {
	return len(g.current)
}

func linetoSlice(str string) []int {
	slice := make([]int, len(str))
	for i := range str {
		slice[i], _ = strconv.Atoi(string(str[i]))
	}
	return slice
}
