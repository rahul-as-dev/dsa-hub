package main

import (
	"fmt"
	"math"
)

type Node struct {
	Name  string
	score int
}
type Edge struct {
	to     Node
	weight int
}

func createGraph() map[Node][]Edge {
	// adjacency list representing graph
	graph := make(map[Node][]Edge)
	// Create nodes
	nodeA := Node{Name: "A"}
	nodeB := Node{Name: "B"}
	nodeC := Node{Name: "C"}
	nodeD := Node{Name: "D"}
	graph[nodeA] = []Edge{
		{to: nodeB, weight: 1},
		{to: nodeC, weight: 4},
	}
	graph[nodeB] = []Edge{
		{to: nodeC, weight: 2},
		{to: nodeD, weight: 5},
	}
	graph[nodeC] = []Edge{
		{to: nodeD, weight: 1},
	}
	graph[nodeD] = []Edge{}
	return graph
}
func main() {
	graph := createGraph()
	visited := make(map[Node]bool)
	ans := dijkstra(graph, visited, Node{Name: "A"}, Node{Name: "D"})
	fmt.Println(ans)
}

func dijkstra(graph map[Node][]Edge, visited map[Node]bool, current Node, end Node) int {
	var queue []Node
	current.score = 0
	queue = append(queue, current)
	visited[current] = true
	ans := math.MaxInt
	for len(queue) > 0 {
		prev := queue[0]
		queue = queue[1:]
		for _, edge := range graph[current] {
			current = edge.to
			current.score = math.MaxInt
			if edge.weight+prev.score < current.score {
				current.score = edge.weight + prev.score
			}
			visited[current] = true
			queue = append(queue, current)
			if current.Name == end.Name {
				ans = min(ans, current.score)
			}
		}
	}
	return ans
}
