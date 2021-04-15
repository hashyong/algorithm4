package graph

import (
	"fmt"
	"testing"
)

func TestLazyPrimMST_LazyMSTInit(t *testing.T) {
	g := NewEWG().In("mst_data.json")
	fmt.Println(g.ToString())

	g1 := NewLazyPrimMSt().LazyMSTInit(g)
	g1.weight()
}
