package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GraphSuite struct {
	suite.Suite

	g Graph
}

func TestGraphSuite(t *testing.T) {
	suite.Run(t, new(GraphSuite))
}

func (s *GraphSuite) SetupTest() {
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
}

func (s *GraphSuite) TestString() {
	s.T().Logf("graph: %s", s.g)
}
