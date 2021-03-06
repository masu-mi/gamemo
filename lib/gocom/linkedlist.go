package gocom

import "bufio"

type basicLinkedList struct {
	size, deg int
	edges     []intSet
}

func newLinkedList(size int) *basicLinkedList {
	ll := &basicLinkedList{size: size, edges: make([]intSet, size)}
	for i := 0; i < size; i++ {
		ll.edges[i] = newIntSet()
	}
	return ll
}

func (ll *basicLinkedList) addEdge(a, b int) {
	ll.addDirectedEdge(a, b)
	ll.addDirectedEdge(b, a)
}

func (ll *basicLinkedList) addDirectedEdge(a, b int) {
	if ll.edges[a].add(b) {
		ll.deg++
	}
}

func (ll *basicLinkedList) exists(a, b int) bool {
	return ll.edges[a].doesContain(b)
}

func nextLinkedList(n, m, offset int, sc *bufio.Scanner) *basicLinkedList {
	ll := newLinkedList(n)
	for i := 0; i < m; i++ {
		x, y := nextInt(sc), nextInt(sc)
		// We use 0-indexed internally
		x -= offset
		y -= offset
		ll.addEdge(x, y)
	}
	return ll
}

func nextDirectedLinkedList(n, m, offset int, sc *bufio.Scanner) *basicLinkedList {
	ll := newLinkedList(n)
	for i := 0; i < m; i++ {
		x, y := nextInt(sc), nextInt(sc)
		// We use 0-indexed internally
		x -= offset
		y -= offset
		ll.addDirectedEdge(x, y)
	}
	return ll
}
