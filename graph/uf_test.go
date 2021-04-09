package graph

import "testing"

func TestUF_initWeight(t *testing.T) {
	g := NewUF()
	g.initWeight(10)
}
