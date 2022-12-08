package eight

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Grid struct {
	previous string
	current  string
	next     string
}

func ProcessGrid(input string) (total int, err error) {
	g := &Grid{}

	file, err := os.Open(input)
	if err != nil {
		return 0, fmt.Errorf("Could not open file %s", input)
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	for {

		cursor := 0
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// last line was the previous line
				total = total + g.ProcessLastLine()
				break
			}
		}

		switch cursor {
		case 0:
			g.current = line
			total = total + g.ProcessFirstLine()
		case 1:
			g.previous = g.current
			g.current = line
		default:
			g.previous = g.current
			g.current = g.next
			g.next = line
			total = total + g.ProcessCurrentLine()
		}

		cursor++
	}
	return
}

func (g Grid) ProcessCurrentLine() int {
	var total int

	for i, v := range g.current {
		prev := strconv.Atoi(g.previous[i])
		cur := strconv.Atoi(v)
		next := strconv.Atoi(g.next[i])
		fmt.Println("prev:", int(prev))
		fmt.Println("next:", int(next))
		if i == 0 {
			total++
		}

		if int(cur) > int(prev) || int(cur) > int(next) {
			total++
		}

		if i == len(g.current) {
			total++
		}
	}

	return total
}

func (g Grid) ProcessLastLine() int {
	return 0
}

func (g Grid) ProcessFirstLine() int {
	return len(g.current)
}

func (g Grid) Print() {
	fmt.Println("------")
	fmt.Println("previous:", g.previous)
	fmt.Println("current:", g.current)
	fmt.Println("next:", g.next)
}
