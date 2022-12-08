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
	g := &Grid{}

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

		switch cursor {
		case 0:
			g.current = linetoSlice(line)
			g.previous = linetoSlice(line) //buffer
			g.next = linetoSlice(line)     // buffer
			total = total + g.ProcessFirstLine()
		case 1:
			copy(g.previous, g.current)
			g.current = linetoSlice(line)
		default:
			copy(g.previous, g.current)
			copy(g.current, g.next)
			g.next = linetoSlice(line)
			total = total + g.ProcessCurrentLine()
		}

		fmt.Println(cursor)
		g.Print()
		cursor++
	}
	return
}

func (g Grid) ProcessCurrentLine() int {
	var total int

	for i := range g.current {
		fmt.Println("[", g.previous[i], "] [", g.current[i], "] [", g.next[i], "]")
		if i == 0 {
			total++
			continue
		}

		if i == (len(g.current) - 1) {
			total++
			break
		}
		if g.current[i] > g.previous[i] || // top
			g.current[i] > g.next[i] || // below
			g.current[i] > g.current[i-1] || // below
			g.current[i] > g.current[i+1] { // below
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
	slice := make([]int, 0)
	for _, i := range str {
		slice = append(slice, int(i))
	}
	return slice
}

// func copyIntSlice(dest, src []int) []int {
// 	if len(dest) == 0 {
// 		dest := make([]int, len(src))
// 	}
// 	copy(dest, src)
// 	return dest
// }
