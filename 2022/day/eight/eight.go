package eight

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Grid struct {
	previous []int
	current  []int
	next     []int
}

func ProcessGrid(input string) (total int, err error) {
	g := &Grid{
		previous: []int{},
		current:  []int{},
		next:     []int{},
	}

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
			g.current = linetoSlice(line)
			total = total + g.ProcessFirstLine()
		case 1:
			g.previous = g.current
			g.current = linetoSlice(line)
		default:
			g.previous = g.current
			g.current = g.next
			g.next = linetoSlice(line)
			total = total + g.ProcessCurrentLine()
		}

		cursor++
	}
	return
}

func (g Grid) ProcessCurrentLine() int {
	var total int

	for i, _ := range g.current {
		fmt.Println("[", g.previous[i], "] [", g.current[i], "] [", g.next[i], "]")
		if i == 0 || i == len(g.current)-1 {
			total++
			break
		}

		if g.current[i] > g.previous[i] || // top
			g.current[i] > g.next[i] || // below
			g.current[i] > g.current[g.current[i-1]] || // below
			g.current[i] > g.current[i+1] { // below
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

func linetoSlice(str string) []int {
	slice := make([]int, len(str))
	for _, i := range str {
		slice = append(slice, int(i))
	}
	return slice
}
