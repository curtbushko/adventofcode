package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	north  int = -1
	south  int = 1
	east   int = 1
	west   int = -1
	weight int = 1 // the number of steps allowed (I am sure this will change in part 2)
)

type Node struct {
	Name    string // We will concat the x,y together to make a name. They should be unique.
	Value   int    // used in Dijkstra to calculate weight/shortest path
	Rune    rune
	Through *Node
}

type Edge struct {
	Node   *Node
	Weight int // needed for calculating shortest path with Dijkstra
}

type Graph struct {
	Start, End    *Node
	Width, Height int
	Nodes         []*Node
	Edges         map[string][]*Edge
	mutex         sync.RWMutex
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

func createGraph(filename string) *Graph {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	g := &Graph{
		Width:  0,
		Height: 0,
		Edges:  make(map[string][]*Edge),
	}
	y := 0
	nodes := make(map[string]*Node) // Used to create the edges
	for {
		line, err := rd.ReadString('\n')
		if err != nil && err == io.EOF {
			// last line was the previous line

			break
		}

		for x, r := range strings.TrimSuffix(line, "\n") {
			name := fmt.Sprintf("%d-%d", y, x)
			n := &Node{
				Name:  name,
				Value: math.MaxInt,
				Rune:  r,
			}
			g.AddNode(n)
			nodes[name] = n
			switch n.Rune {
			case 'S':
				n.Rune = 'a' - 1
				g.Start = n
			case 'E':
				n.Rune = 'z' + 1
				g.End = n
			}
		}
		y++
		g.Width = len(line) - 1
	}
	g.Height = y

	// Add all the edges
	for _, node := range g.Nodes {
		coords := strings.Split(node.Name, "-")
		y, _ := strconv.Atoi(coords[0])
		x, _ := strconv.Atoi(coords[1])

		north := fmt.Sprintf("%d-%d", y-1, x)
		south := fmt.Sprintf("%d-%d", y+1, x)
		east := fmt.Sprintf("%d-%d", y, x+1)
		west := fmt.Sprintf("%d-%d", y, x-1)

		// r1 := node.Rune

		if _, exists := nodes[north]; exists {
			if (node.Rune == nodes[north].Rune) || (node.Rune+1 == nodes[north].Rune) || (node.Rune > nodes[north].Rune) {
				fmt.Printf("Node Rune: %c North Rune: %c\n", node.Rune, nodes[north].Rune)
				g.AddEdge(nodes[node.Name], nodes[north], weight)
			}
		}
		if _, exists := nodes[south]; exists {
			if (node.Rune == nodes[south].Rune) || (node.Rune+1 == nodes[south].Rune) || (node.Rune > nodes[south].Rune) {
				fmt.Printf("Node Rune: %c South Rune: %c\n", node.Rune, nodes[south].Rune)
				g.AddEdge(nodes[node.Name], nodes[south], weight)
			}
		}
		if _, exists := nodes[east]; exists {
			if (node.Rune == nodes[east].Rune) || (node.Rune+1 == nodes[east].Rune) || (node.Rune > nodes[east].Rune) {
				fmt.Printf("Node Rune: %c East Rune: %c\n", node.Rune, nodes[east].Rune)
				g.AddEdge(nodes[node.Name], nodes[east], weight)
			}
		}
		if _, exists := nodes[west]; exists {
			if (node.Rune == nodes[west].Rune) || (node.Rune+1 == nodes[west].Rune) || (node.Rune > nodes[west].Rune) {
				fmt.Printf("Node Rune: %c West Rune: %c\n", node.Rune, nodes[west].Rune)
				g.AddEdge(nodes[node.Name], nodes[west], weight)
			}
		}

	}

	dijkstra(g)
	for _, node := range g.Nodes {
		fmt.Printf("Shortest time from %c to %c is %d\n",
			g.Start.Rune, node.Rune, node.Value)
		for n := node; n.Through != nil; n = n.Through {
			fmt.Printf("%c <- ", n.Rune)
		}
		fmt.Printf("%c\n", g.Start.Rune)
		fmt.Println()
	}

	return g
}

func (g *Graph) Print() {
	fmt.Println("-----", g.Height, g.Width)
	for y, n := range g.Nodes {
		if y%g.Width == 0 {
			fmt.Println()
		}
		fmt.Printf("%c", n.Rune)
	}
	fmt.Println()
}

func (g *Graph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.Name == name {
			node = n
		}
	}
	return
}

func (g *Graph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[n1.Name] = append(g.Edges[n1.Name], &Edge{n2, weight})
	// g.Edges[n2.Name] = append(g.Edges[n2.Name], &Edge{n1, weight})
}

// Stole this heap
type Heap struct {
	elements []*Node
	mutex    sync.RWMutex
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].Value < h.elements[parent(i)].Value; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// pop the top of the heap, which is the min Value
func (h *Heap) Pop() (i *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i = h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.rearrange(0)
	return
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].Value < h.elements[smallest].Value {
		smallest = left
	}
	if right < size && h.elements[right].Value < h.elements[smallest].Value {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func dijkstra(graph *Graph) {
	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := graph.GetNode(graph.Start.Name)
	startNode.Value = 0
	heap.Push(startNode)

	fmt.Println("Starting Dijkstra")
	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.Name] = true
		edges := graph.Edges[current.Name]
		for _, edge := range edges {
			if !visited[edge.Node.Name] {
				heap.Push(edge.Node)
				if current.Value+edge.Weight < edge.Node.Value {
					edge.Node.Value = current.Value + edge.Weight
					edge.Node.Through = current
				}
			}
		}
	}

	fmt.Println("Done")
}
