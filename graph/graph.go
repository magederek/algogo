package graph

import "fmt"

type Graph interface {
	V() int
	E() int
	AddEdge(v, w int)
	Adj(v int) []int
	String() string
}

type AdjacencyListGraph struct {
	v   int
	e   int
	adj [][]int
}

func NewAdjacencyListGraph(vertices int, edges [][2]int) *AdjacencyListGraph {
	g := &AdjacencyListGraph{
		v:   vertices,
		adj: make([][]int, vertices),
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func (g *AdjacencyListGraph) V() int {
	return g.v
}

func (g *AdjacencyListGraph) E() int {
	return g.e
}

func (g *AdjacencyListGraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

func (g *AdjacencyListGraph) Adj(v int) []int {
	return g.adj[v]
}

// String 图的邻接表的字符串表示法
func (g *AdjacencyListGraph) String() string {
	var s string
	s += fmt.Sprintf("%d vertices, %d edges\n", g.v, g.e)

	for v := 0; v < g.v; v++ {
		s += fmt.Sprintf("%d: ", v)
		for _, w := range g.adj[v] {
			s += fmt.Sprintf("%d ", w)
		}
		s += "\n"
	}

	return s
}
