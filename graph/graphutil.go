package graph

// Degree 计算v的度数
func Degree(g Graph, v int) int {
	var degree int
	for range g.Adj(v) {
		degree++
	}
	return degree
}

// MaxDegree 计算所有顶点的最大度数
func MaxDegree(g Graph) int {
	var max int
	for v := 0; v < g.V(); v++ {
		if degree := Degree(g, v); degree > max {
			max = degree
		}
	}
	return max
}

// AvgDegree 计算所有顶点的平均度数
func AvgDegree(g Graph) float64 {
	return 2.0 * float64(g.E()) / float64(g.V())
}

// NumberOfSelfLoops 计算自环的个数
func NumberOfSelfLoops(g Graph) int {
	var count int
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if v == w {
				count++
			}
		}
	}
	return count / 2
}
