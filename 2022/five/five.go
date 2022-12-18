package five

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack struct {
	elements []rune
}

func (s *stack) push(r rune) {
	s.elements = append(s.elements, r)
}

func (s *stack) pop() rune {
	var r rune
	if len(s.elements) == 0 {
		return r
	}
	r = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return r
}

func (s *stack) popN(n int) []rune {
	if len(s.elements) < n {
		panic(fmt.Sprintf("cannot pop %d items from stack len %d", n, len(s.elements)))
	}
	r := s.elements[len(s.elements)-n : len(s.elements)]
	s.elements = s.elements[:len(s.elements)-n]
	return r
}

func (s *stack) pushN(r []rune) {
	s.elements = append(s.elements, r...)
}

func (s *stack) addToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s stack) String() string {
	var str string
	for _, r := range s.elements {
		str += string(r) + " "
	}
	return str
}

func run(filename, header string, stackSize, part int) string {
	// Read input file
	input, _ := os.Open(filename)
	defer input.Close()
	sc := bufio.NewScanner(input)

	// create slice of stacks
	stacks := make([]stack, stackSize)

	// Parsing the input
	sc.Scan()
	for sc.Text() != header {
		for i, r := range sc.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].addToBottom(r)
			}
		}
		sc.Scan()
	}
	// Read empty line
	sc.Scan()

	for sc.Scan() {
		var toMove, from, to int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &toMove, &from, &to)

		// Part 2, move blocks of elements
		if part == 2 {
			stacks[to-1].pushN(stacks[from-1].popN(toMove))
			continue
		}

		for move := 0; move < toMove; move++ {
			// Part 1, move 1 at a time
			stacks[to-1].push(stacks[from-1].pop())
		}
	}

	var out string
	for _, s := range stacks {
		if part == 2 {

			out = out + string(s.popN(1))
			continue
		}
		out = out + string(s.pop())
	}

	return strings.TrimSpace(strings.ReplaceAll(out, " ", ""))
}

// func readInput(filename string) ([]byte, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("Cannot open input file %s: %s", filename, err)
// 	}
// 	defer file.Close()
//
// 	in, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		return nil, fmt.Errorf("Cannot read input file %s: %s", filename, err)
// 	}
//
// 	return in, nil
// }
//
// func reverse(a []string) []string {
// 	reversed := make([]string, len(a), cap(a))
//
// 	for i, s := range a[:len(a)/2] {
// 		reversed[i], reversed[len(a)-1-i] = a[len(a)-1-i], s
// 	}
//
// 	return reversed
// }
