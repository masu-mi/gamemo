package gocom

import (
	"testing"
)

func Test(t *testing.T) {
	// TODO not tested
	g := newLinkedList(5)
	g.addEdge(0, 1)
	if !g.exists(0, 1) {
		t.Errorf("edge(0, 1) exists in G")
	}
	if !g.exists(1, 0) {
		t.Errorf("edge(0, 1) exists in G")
	}
	if g.exists(1, 1) {
		t.Errorf("edge(0, 1) shouldn't exist in G")
	}
	if g.exists(1, 2) {
		t.Errorf("edge(0, 1) shouldn't exist in G")
	}
}
