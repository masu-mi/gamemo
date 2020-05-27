package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	parseProblem(os.Stdin)
}

func parseProblem(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	m := scanInt(sc)
	cg := scanGraph(n, m, sc)

	c, p := findLLP(cg)
	fmt.Printf("%d: %v\n", c, p)
}

func findLLP(cg *graph) (cost int, path []int) {
	cs := make([]int, cg.size)
	ps := make([]int, cg.size)
	costs := make([]int, cg.size)
	for i := 0; i < cg.size; i++ {
		costs[i] = math.MinInt32
		cs[i] = -1
	}
	ps[0] = 0
	costs[0] = 0
	_dfs(cg, 0, cg.size-1, costs, ps, cs)
	n := 0
	c := 0
	path = []int{0}
	for n != cg.size-1 && c < cg.size {
		n = cs[n]
		if n == -1 {
			break
		}
		path = append(path, n)
		c++
	}
	return costs[0], path
}

func _dfs(cg *graph, s, t int, costs []int, ps, cs []int) {
	if s == t {
		costs[t] = 0
		cs[s] = t
		return
	}
	for n, c := range cg.edges[s] {
		if n == ps[s] {
			continue
		}
		ps[n] = s
		_dfs(cg, n, t, costs, ps, cs)
		if cost := costs[n] + c; costs[s] < cost {
			costs[s] = cost
			cs[s] = n
		}
	}
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
		c := scanInt(sc)
		g.addEdge(x, y, c)
	}
	return g
}

func (g *graph) addEdge(x, y, c int) {
	g.edges[x].add(y, c)
	g.edges[y].add(x, c)
}

func (g *graph) addDirectedEdge(x, y, c int) {
	g.edges[x].add(y, c)
}

func (g *graph) exists(x, y int) bool {
	return g.edges[x].doesContain(y)
}

type nodeSet map[int]int

func newSet() nodeSet {
	return make(map[int]int)
}

func (s nodeSet) add(item, cost int) {
	s[item] = cost
}

func (s nodeSet) doesContain(item int) bool {
	_, ok := s[item]
	return ok
}
func (s nodeSet) cost(item int) int {
	c, ok := s[item]
	if ok {
		return c
	}
	return math.MinInt32
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

func resolve(n int) int {
	return n
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
