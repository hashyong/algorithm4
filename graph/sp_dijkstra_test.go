package graph

import (
	"fmt"
	"testing"
)

func TestDijkstra_DistTo(t *testing.T) {
	g := NewDirectedEWG().In("mst_data.json")
	fmt.Println(g.ToString())

	g1 := NewDijkstra().init(g, 0)
	fmt.Println(g1.hasPathTo(1))
	g1.pathTo(1)
	g1.pathTo(3)
}

func TestDijkstraAllPairsSP_Init(t *testing.T) {
	g := NewDirectedEWG().In("mst_data.json")
	fmt.Println(g.ToString())

	g1 := NewDijkstraAllPairsSP().Init(g)
	g1.path(0, 1)
	g1.path(3, 1)
}
