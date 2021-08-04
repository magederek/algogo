package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SearchSuite struct {
	suite.Suite

	g      Graph
	search Search
	paths  Paths
}

func TestSearchSuite(t *testing.T) {
	suite.Run(t, new(SearchSuite))
}

func (s *SearchSuite) SetupTest() {
	vertices := 13
	edges := [][2]int{
		{0, 5},
		{4, 3},
		{0, 1},
		{9, 12},
		{6, 4},
		{5, 4},
		{0, 2},
		{11, 12},
		{9, 10},
		{0, 6},
		{7, 8},
		{9, 11},
		{5, 3},
	}
	s.g = NewAdjacencyListGraph(vertices, edges)

	dfs := &DepthFirstSearch{}
	s.search = dfs
	s.paths = dfs
}

func (s *SearchSuite) TestMarked() {
	searchSrc := 0
	marked := make([]int, 0)
	for v := 0; v < s.g.V(); v++ {
		if s.search.Marked(s.g, searchSrc, v) {
			marked = append(marked, v)
		}
	}

	s.T().Logf("marked: %v", marked)
}

func (s *SearchSuite) TestCount() {
	searchSrc := 0
	count := s.search.Count(s.g, searchSrc)

	s.T().Logf("count: %d", count)
}

func (s *SearchSuite) TestPathTo() {
	srcVertix := 0
	destVertix := 3

	path := s.paths.PathTo(s.g, srcVertix, destVertix)

	s.T().Logf("path [%d->%d]: %v", srcVertix, destVertix, path)
}
