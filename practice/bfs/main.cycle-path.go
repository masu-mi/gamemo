package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 1000000
)

var sc *bufio.Scanner

func initScanner(r io.Reader) *bufio.Scanner {
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	return sc
}

func main() {
	sc = initScanner(os.Stdin)
	resolve()
}

func resolve() {
	n := scanInt(sc)
	g := newGraph(n)
	m := scanInt(sc)
	for i := 0; i < m; i++ {
		// premise: this graph includes single ring
		x, y := scanInt(sc), scanInt(sc)
		g.addEdge(x-1, y-1)
	}
	members := getCycleMember(g)
	q := scanInt(sc)
	for i := 0; i < q; i++ {
		// premise: this graph includes single ring
		x, y := scanInt(sc), scanInt(sc)
		if involved(members, x, y) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
	return
}

func involved(ringMembers []bool, x, y int) bool {
	return x != y && ringMembers[x-1] && ringMembers[y-1]
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func getCycleMember(g *graph) []bool {
	degrees := getDegrees(g)
	var q []int
	for i, deg := range degrees {
		if deg == 1 {
			q = append(q, i)
		}
	}

	isCycleMember := make([]bool, g.nodeNum)
	for i := range degrees {
		isCycleMember[i] = true
	}
	for len(q) > 0 {
		h := q[0]
		q = q[1:]
		isCycleMember[h] = false
		for _, next := range g.edges[h] {
			// We use degrees to modify BFS
			degrees[next]--
			if degrees[next] == 1 {
				q = append(q, next)
			}
		}
	}
	return isCycleMember
}

func getDegrees(g *graph) []int {
	degrees := make([]int, g.nodeNum)
	for n, d := range g.edges {
		degrees[n] = len(d)
	}
	return degrees
}

type graph struct {
	nodeNum int
	edges   map[int][]int
}

func newGraph(n int) *graph {
	return &graph{
		nodeNum: n,
		edges:   map[int][]int{},
	}
}

func (g *graph) addEdge(s, t int) {
	g.addDirectedEdge(s, t)
	g.addDirectedEdge(t, t)
}

func (g *graph) addDirectedEdge(s, t int) {
	g.edges[s] = append(g.edges[s], t)
}
