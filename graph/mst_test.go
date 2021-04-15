package graph

import (
	"fmt"
	"testing"
)

func TestEdgeWeightedGraph_EdgeWeightedGraphIn(t *testing.T) {
	fmt.Println(NewEWG().In("mst_data.json").ToString())
}
