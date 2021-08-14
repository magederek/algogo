package graph

// MST 最小生成树
type MST interface {
	// Edges 最小生成树的所有边
	Edges() []*Edge
	// Weight 最小生成树的权重
	Weight() float64
}
