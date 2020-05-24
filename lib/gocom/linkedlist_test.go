package gocom

import (
	"testing"

	"github.com/k0kubun/pp"
)

func Test(t *testing.T) {
	// TODO not tested
	g := newLinkedList(5)
	g.addEdge(0, 1)
	pp.Println(g)
	g.addEdge(2, 4)
	g.addEdge(3, 4)
	pp.Println(g)
	g.addEdge(2, 4)
	pp.Println(g)
}
