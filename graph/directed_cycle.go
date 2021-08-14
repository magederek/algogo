package graph

// DirectedCycle 寻找有向环
type DirectedCycle interface {
	// HasCycle G是否包含有向环
	HasCycle() bool
	// Cycle 有向环中的所有顶点
	Cycle() []int
}

type DfsDirectedCycle struct {
	digraph Digraph

	marked  []bool
	edgeTo  []int
	cycle   []int
	onStack []bool
}

func NewDfsDirectedCycle(digraph Digraph) *DfsDirectedCycle {
	directedCycle := new(DfsDirectedCycle)
	directedCycle.digraph = digraph
	directedCycle.marked = make([]bool, digraph.V())
	directedCycle.edgeTo = make([]int, digraph.V())
	directedCycle.onStack = make([]bool, digraph.V())

	for v := 0; v < digraph.V(); v++ {
		if !directedCycle.marked[v] {
			directedCycle.dfs(digraph, v)
		}
	}

	return directedCycle
}

func (c *DfsDirectedCycle) dfs(g Digraph, v int) {
	c.onStack[v] = true
	c.marked[v] = true

	for _, w := range g.Adj(v) {
		if c.HasCycle() {
			return
		}

		if !c.marked[w] {
			c.edgeTo[w] = v
			c.dfs(g, w)
		} else if c.onStack[w] {
			for lv := v; lv != w; lv = c.edgeTo[lv] {
				c.cycle = append([]int{lv}, c.cycle...)
			}
			c.cycle = append([]int{w}, c.cycle...)
			c.cycle = append([]int{v}, c.cycle...)
		}

	}

	c.onStack[v] = false
}

// HasCycle G是否包含有向环
func (c *DfsDirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

// Cycle 有向环中的所有顶点
func (c *DfsDirectedCycle) Cycle() []int {
	return c.cycle
}
