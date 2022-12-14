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
	Edges         map[Node][]*Edge
	mutex         sync.RWMutex
}

type Vertex struct {
	Node     *Node
	Distance int
}

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
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
		Edges:  make(map[Node][]*Edge),
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

	stuff, more := getShortestPath(g.Start, g.End, g)
	fmt.Println(stuff)
	fmt.Println(more)

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

func getShortestPath(startNode *Node, endNode *Node, g *Graph) ([]int, int) {
	visited := make(map[int]bool)
	dist := make(map[int]int)
	prev := make(map[int]int)
	// pq := make(PriorityQueue, 1)
	// heap.Init(&pq)
	q := NodeQueue{}
	pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)
	// im := 0
	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]

		for _, val := range near {
			if !visited[val.Node.Value] {
				if dist[v.Node.Value]+val.Weight < dist[val.Node.Value] {
					store := Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Value] + val.Weight,
					}
					dist[val.Node.Value] = dist[v.Node.Value] + val.Weight
					// prev[val.Node.Value] = fmt.Sprintf("->%s", v.Node.Value)
					prev[val.Node.Value] = v.Node.Value
					pq.Enqueue(store)
				}
				// visited[val.Node.value] = true
			}
		}
	}
	fmt.Println(dist)
	fmt.Println(prev)
	pathval := prev[endNode.Value]
	var finalArr []int
	finalArr = append(finalArr, endNode.Value)
	for pathval != startNode.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)
	fmt.Println(finalArr)
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Value]
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

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.mutex.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	// g.Edges[*n2] = append(g.Edges[*n2], &ed2)
	g.mutex.Unlock()
}

// Enqueue adds an Node to the end of the queue
func (s *NodeQueue) Enqueue(t Vertex) {
	s.Lock.Lock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		s.Lock.Unlock()
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Vertex{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
	// s.items = append(s.items, t)
	s.Lock.Unlock()
}

// Dequeue removes an Node from the start of the queue
func (s *NodeQueue) Dequeue() *Vertex {
	s.Lock.Lock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Lock.Unlock()
	return &item
}

// NewQ Creates New Queue
func (s *NodeQueue) NewQ() *NodeQueue {
	s.Lock.Lock()
	s.Items = []Vertex{}
	s.Lock.Unlock()
	return s
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}
