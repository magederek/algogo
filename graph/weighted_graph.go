package graph

import (
	"fmt"
)

type EdgeWeightedGraph struct {
	v   int
	e   int
	adj [][]*Edge
}

func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	g := new(EdgeWeightedGraph)
	g.v = v
	g.e = 0
	g.adj = make([][]*Edge, g.v)
	for vertex := 0; vertex < g.v; vertex++ {
		g.adj[vertex] = make([]*Edge, 0)
	}
	return g
}

func (g *EdgeWeightedGraph) V() int {
	return g.v
}

func (g *EdgeWeightedGraph) E() int {
	return g.e
}

func (g *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.v
	w := e.Other(v)
	g.adj[v] = append(g.adj[v], e)
	g.adj[w] = append(g.adj[w], e)
	g.e++
}

func (g *EdgeWeightedGraph) Adj(v int) []*Edge {
	return g.adj[v]
}

type Edge struct {
	v      int
	w      int
	weight float64
}

func NewEdge(v, w int, weight float64) *Edge {
	edge := new(Edge)
	edge.v = v
	edge.w = w
	edge.weight = weight
	return edge
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) Either() int {
	return e.v
}

func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	if vertex == e.w {
		return e.v
	}
	return -1
}

func (e *Edge) CompareTo(that *Edge) int {
	if e.Weight() < that.Weight() {
		return -1
	} else if e.Weight() > that.Weight() {
		return +1
	}
	return 0
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.weight)
}
