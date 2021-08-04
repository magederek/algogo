package graph

type Search interface {
	// Marked v和s是连通的吗
	Marked(g Graph, s int, v int) bool
	// Count 与s连通的顶点总数
	Count(g Graph, s int) int
}

type Paths interface {
	// HasPathTo 是否存在从s到v的路径
	HasPathTo(g Graph, s, v int) bool
	// PathTo s到v的路径，如果不存在则返回nil
	PathTo(g Graph, s, v int) []int
}

type DepthFirstSearch struct {
}

func (d *DepthFirstSearch) Dfs(g Graph, s int, visit func(lv, cv int) bool) {
	marked := make([]bool, g.V())
	d.dfsRecursive(g, s, marked, visit)
}

func (d *DepthFirstSearch) dfsRecursive(g Graph, s int, marked []bool, visit func(lv, cv int) bool) (isContinue bool) {
	marked[s] = true

	for _, w := range g.Adj(s) {
		if !marked[w] {
			isContinue = visit(s, w)
			if !isContinue {
				return false
			}

			isContinue = d.dfsRecursive(g, w, marked, visit)
			if !isContinue {
				return false
			}
		}
	}

	return true
}

// Marked v和s是连通的吗
func (d *DepthFirstSearch) Marked(g Graph, s int, v int) bool {
	if s == v {
		return true
	}

	var marked bool

	d.Dfs(g, s, func(lv, cv int) bool {
		if v == cv {
			marked = true
			return false
		}
		return true
	})

	return marked
}

// Count 与s连通的顶点总数
func (d *DepthFirstSearch) Count(g Graph, s int) int {
	count := 1

	d.Dfs(g, s, func(lv, cv int) bool {
		count++
		return true
	})
	return count
}

// HasPathTo 是否存在从s到v的路径
func (d *DepthFirstSearch) HasPathTo(g Graph, s, v int) bool {
	return d.Marked(g, s, v)
}

// PathTo s到v的路径，如果不存在则返回nil
func (d *DepthFirstSearch) PathTo(g Graph, s, v int) []int {
	edgeTo := make([]int, g.V())

	d.Dfs(g, s, func(lv, cv int) bool {
		edgeTo[cv] = lv
		return true
	})

	path := make([]int, 0)
	path = append([]int{v}, path...)
	for lv := edgeTo[v]; lv != s; lv = edgeTo[lv] {
		path = append([]int{lv}, path...)
	}
	path = append([]int{s}, path...)

	return path
}
