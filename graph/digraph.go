package graph

import "fmt"

type Digraph interface {
	V() int
	E() int
	AddEdge(v, w int)
	Adj(v int) []int
	Reverse() Digraph
	String() string
}

type AdjacencyListDigraph struct {
	v   int
	e   int
	adj [][]int
}

func NewAdjacencyListDigraph(vertices int, edges [][2]int) *AdjacencyListDigraph {
	g := &AdjacencyListDigraph{
		v:   vertices,
		adj: make([][]int, vertices),
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	return g
}

func (g *AdjacencyListDigraph) V() int {
	return g.v
}

func (g *AdjacencyListDigraph) E() int {
	return g.e
}

func (g *AdjacencyListDigraph) AddEdge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

func (g *AdjacencyListDigraph) Adj(v int) []int {
	return g.adj[v]
}

func (g *AdjacencyListDigraph) Reverse() Digraph {
	digraph := NewAdjacencyListDigraph(g.V(), nil)
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			digraph.AddEdge(w, v)
		}
	}
	return digraph
}

// String 图的邻接表的字符串表示法
func (g *AdjacencyListDigraph) String() string {
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
