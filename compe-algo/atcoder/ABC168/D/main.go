package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) *graph {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n, m := scanInt(sc), scanInt(sc)
	g := scanGraph(n, m, sc)
	return g
}

type graph struct {
	size  int
	edges []nodeSet
}

func newGraph(n int) *graph {
	g := &graph{
		size:  n,
		edges: make([]nodeSet, n),
	}
	for i := 0; i < n; i++ {
		g.edges[i] = newSet()
	}
	return g
}

func scanGraph(n, m int, sc *bufio.Scanner) *graph {
	g := newGraph(n)
	for i := 0; i < m; i++ {
		x, y := scanInt(sc), scanInt(sc)
		// 0-indexed
		x--
		y--
		g.addEdge(x, y)
	}
	return g
}

func (g *graph) addEdge(x, y int) {
	g.edges[x].add(y)
	g.edges[y].add(x)
}

func (g *graph) addDirectedEdge(x, y int) {
	g.edges[x].add(y)
}

func (g *graph) exists(x, y int) bool {
	return g.edges[x].doesContain(y)
}

type nodeSet map[int]none

func newSet() nodeSet {
	return make(map[int]none)
}

func (s nodeSet) add(item int) {
	s[item] = mark
}

func (s nodeSet) doesContain(item int) bool {
	_, ok := s[item]
	return ok
}

func (s nodeSet) size() int {
	return len(s)
}

func (s nodeSet) members() (l []int) {
	for k := range s {
		l = append(l, k)
	}
	return l
}

var mark none

type none struct{}

func resolve(g *graph) {
	// g is 0-indexed
	parents := make([]int, g.size)
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}
	parents[0] = 0
	q := []int{0}
	// BFS
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for k := range g.edges[cur] {
			if parents[k] != -1 {
				continue
			}
			parents[k] = cur
			q = append(q, k)
		}
	}
	// resolved
	fmt.Println("Yes")
	for i := 1; i < len(parents); i++ { // start from 2
		// recover to 1-indexed
		fmt.Println(parents[i] + 1)
	}
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
