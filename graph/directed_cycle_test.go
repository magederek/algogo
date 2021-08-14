package graph

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DirectedCycleSuite struct {
	suite.Suite

	g  Digraph
	dc DirectedCycle
}

func TestDirectedCycleSuite(t *testing.T) {
	suite.Run(t, new(DirectedCycleSuite))
}

func (s *DirectedCycleSuite) SetupTest() {
	vertices := 13
	edges := [][2]int{
		{0, 5},
		{3, 5},
		{5, 4},
		{4, 3},
	}
	s.g = NewAdjacencyListDigraph(vertices, edges)
	s.dc = NewDfsDirectedCycle(s.g)
}

func (s *DirectedCycleSuite) TestHasCycle() {
	hasCycle := s.dc.HasCycle()
	s.Assert().True(hasCycle)
}

func (s *DirectedCycleSuite) TestCycle() {
	cycle := s.dc.Cycle()
	s.Assert().Equal(cycle, []int{3, 5, 4, 3})
}
